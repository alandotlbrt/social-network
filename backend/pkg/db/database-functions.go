package db

import (
	"database/sql"
	"fmt"
)

func UpdateInDatabase(db *sql.DB, tableName, columnName, rowName, rowValue, value string) (int64, error) {

	var err error
	if db == nil {
		db, err = OpenDB("database")
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	}

	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s = '%s'", tableName, columnName, rowName, rowValue)

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(value)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func CheckInDatabase(db *sql.DB, tableName string, columnName string, value interface{}) (bool, error) {

	var err error

	if db == nil {
		db, err = OpenDB("database")
		if err != nil {
			fmt.Println("in function ", err)
			return false, err
		}
	}

	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s = ?)", tableName, columnName)

	var exists bool
	err = db.QueryRow(query, value).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("error checking value existence: %v", err)
	}

	return exists, nil
}

func CheckInDatabaseAndGet(db *sql.DB, tableName string, columnName string, value string, columnNameReturn string) (string, error) {
	var err error

	if db == nil {
		db, err = OpenDB("database")
		if err != nil {
			fmt.Println("in function ", err)
			return "", err
		}
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", columnNameReturn, tableName, columnName)

	var exists string
	err = db.QueryRow(query, value).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", fmt.Errorf("error checking value existence: %v", err)
	}

	return exists, nil
}

func GetInDatabase(db *sql.DB, tableName string, columnName string, value interface{}, columnNameReturn string) (string, error) {
	var err error
	if db == nil {
		db, err = OpenDB("database")
		if err != nil {
			fmt.Println("in function ", err)
			return "", err
		}
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", columnNameReturn, tableName, columnName)

	var result string
	err = db.QueryRow(query, value).Scan(&result)
	if err != nil {
		fmt.Println("in function ", err)
		return "", err
	}

	return result, nil
}
