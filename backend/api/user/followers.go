package profile

import (
	"fmt"
	"log"
	"socialNetwork/pkg/db"
	config "socialNetwork/pkg/db/var"
)

func UserFollowModels(iduser, idtofollow string) bool {

	database, err := db.OpenDB("database")
	if err != nil {
		fmt.Println(err)
		return false
	}

	query := "INSERT INTO follow ('id_user', 'follow') VALUES (?, ?)"

	stmt, err := database.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return false
	}

	_, err = stmt.Exec(iduser, idtofollow)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func UserUnFollowModels(iduser, idunfollow string) bool {

	database, err := db.OpenDB("database")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données :", err)
		return false
	}
	defer database.Close()

	query := "DELETE FROM follow WHERE id_user = ? AND follow = ?"

	stmt, err := database.Prepare(query)
	if err != nil {
		fmt.Println("Erreur lors de la préparation de la requête :", err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(iduser, idunfollow)
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête :", err)
		return false
	}

	return true
}

func WaitingFollowModel(idWaiting, idRequest string) bool {

	database, err := db.OpenDB("database")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données :", err)
		return false
	}
	defer database.Close()

	isPrivate, err := db.CheckInDatabase(database, "users", "privacy", "Private")
	if err != nil || !isPrivate {
		return false
	}

	query := "INSERT INTO waiting_follow ('id_who_waiting', 'id_who_accept') VALUES (?,?)"

	stmt, err := database.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return false
	}

	_, err = stmt.Exec(idWaiting, idRequest)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func UnRequestFollowModels(iduser, idunfollow string) bool {
	database, err := db.OpenDB("database")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données :", err)
		return false
	}
	defer database.Close()

	query := "DELETE FROM waiting_follow WHERE id_who_waiting = ? AND id_who_accept = ?"

	stmt, err := database.Prepare(query)
	if err != nil {
		fmt.Println("Erreur lors de la préparation de la requête :", err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(iduser, idunfollow)
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête :", err)
		return false
	}

	return true
}

func FetchWaitingFollowersModels(iduserWhoAccept string) []config.UserFollowStruct {

	var usersWhoWaiting []config.UserFollowStruct

	database, err := db.OpenDB("database")
	if err != nil {
		fmt.Println(err)
		return []config.UserFollowStruct{}
	}
	stmt, err := database.Prepare(`SELECT u.id_user, u.first_name, u.pp FROM users u INNER JOIN waiting_follow wf ON wf.id_who_waiting = u.id_user AND wf.id_who_accept = ?`)
	if err != nil {
		fmt.Println("Erreur lors de la préparation de la requête :", err)
		return []config.UserFollowStruct{}
	}
	defer stmt.Close()

	rows, err := stmt.Query(iduserWhoAccept)
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête :", err)
		return []config.UserFollowStruct{}
	}
	defer rows.Close()

	for rows.Next() {
		var user config.UserFollowStruct
		if err := rows.Scan(&user.Id, &user.Username, &user.Pictures); err != nil {
			log.Fatal("Erreur lors de la lecture des résultats :", err)
		}
		usersWhoWaiting = append(usersWhoWaiting, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal("Erreur lors de l'itération :", err)
	}

	return usersWhoWaiting
}

func AcceptFollower(id_who_waiting, id_who_accept string) bool {

	database, err := db.OpenDB("database")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données :", err)
		return false
	}
	defer database.Close()

	if UserFollowModels(id_who_waiting, id_who_accept) {
		if !UnRequestFollowModels(id_who_waiting, id_who_accept) {
			return false
		}

	} else {
		return false
	}

	return true
}
