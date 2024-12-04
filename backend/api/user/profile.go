package profile

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"socialNetwork/pkg/db"
	config "socialNetwork/pkg/db/var"
)

func GetUserData(w http.ResponseWriter, iduserChecking, id string) (bool, config.InformationUser) {
	var structUserData config.InformationUser

	database, err := db.OpenDB("database")
	if err != nil {
		fmt.Println(err)
		return false, structUserData
	}

	var Pseudo sql.NullString

	err = database.QueryRow("SELECT id_user,password, email, first_name, last_name, birthday, pp, nickname, about_me, privacy, follower, follow FROM Users WHERE id_user = ?", id).Scan(&structUserData.Id, &Password, &structUserData.Email, &structUserData.Firstname, &structUserData.Lastname, &structUserData.Birthday, &structUserData.Profil_Pictures, &Pseudo, &structUserData.About, &structUserData.Privacy, &structUserData.Follow, &structUserData.Follower)
	if Pseudo.Valid {
		structUserData.Username = Pseudo.String
	} else {
		structUserData.Username = ""
	}

	structUserData.IsFollowing = isFollowTheUser(database, iduserChecking, id)
	structUserData.IsWaitingFollow = isWaitingFollowTheUser(database, iduserChecking, id)
	if structUserData.Privacy == "Private" {
		structUserData.FollowNumberInbox, err = countNotificationsOfTheProfile(database, structUserData.Id)
	}

	if err != nil {
		fmt.Println(err)
		return false, structUserData
	}
	return true, structUserData
}

func countNotificationsOfTheProfile(db *sql.DB, idWaiting string) (int, error) {

	var count int

	queryCount := "select count(id_who_waiting) from waiting_follow where id_who_accept = ?"

	stmt, err := db.Prepare(queryCount)
	if err != nil {
		fmt.Println("count error : ", err)
		return 0, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(idWaiting).Scan(&count)
	if err != nil {
		fmt.Println("query row:", err)
		return 0, err
	}

	return count, nil
}

func isFollowTheUser(db *sql.DB, idUser, id string) bool {
	query := "SELECT EXISTS (SELECT 1 FROM follow WHERE id_user = ? AND follow = ?)"

	var exists bool
	err := db.QueryRow(query, idUser, id).Scan(&exists)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return false
	}

	return exists
}

func isWaitingFollowTheUser(db *sql.DB, idUser, id string) bool {
	query := "SELECT EXISTS (SELECT 1 FROM waiting_follow WHERE id_who_waiting = ? AND id_who_accept = ?)"

	var exists bool
	err := db.QueryRow(query, idUser, id).Scan(&exists)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return false
	}

	return exists
}

func FetchFollowersFromProfile(iduser, userstateID, typeOfFollow string) []config.UserFollowStruct {

	var err error
	var stmt *sql.Stmt

	var usersWhoWaiting []config.UserFollowStruct

	database, err := db.OpenDB("database")
	if err != nil {
		fmt.Println(err)
		return []config.UserFollowStruct{}
	}

	if typeOfFollow == "followers" {
		stmt, err = database.Prepare(`SELECT u.id_user, u.first_name, u.pp, u.privacy  FROM users u INNER JOIN follow f ON f.id_user = u.id_user AND f.follow = ?`)
	} else if typeOfFollow == "following" {
		stmt, err = database.Prepare(`SELECT u.id_user, u.first_name, u.pp, u.privacy FROM users u INNER JOIN follow wf ON wf.follow = u.id_user AND wf.id_user = ?`)
	}

	if err != nil {
		fmt.Println("Erreur lors de la préparation de la requête :", err)
		return []config.UserFollowStruct{}
	}
	defer stmt.Close()

	rows, err := stmt.Query(iduser)
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête :", err)
		return []config.UserFollowStruct{}
	}
	defer rows.Close()

	for rows.Next() {
		var user config.UserFollowStruct
		if err := rows.Scan(&user.Id, &user.Username, &user.Pictures, &user.Privacy); err != nil {
			log.Fatal("Erreur lors de la lecture des résultats :", err)
		}
		user.HeIsFollowing = isFollowTheUser(database, userstateID, user.Id)
		user.HeIsWaiting = isWaitingFollowTheUser(database, userstateID, user.Id)
		usersWhoWaiting = append(usersWhoWaiting, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal("Erreur lors de l'itération :", err)
	}

	return usersWhoWaiting
}
