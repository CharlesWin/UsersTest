package database

import (
	. "UsersTest/internal/config"
	"UsersTest/internal/parser"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

func GetFromId(id int) (*parser.User, error) {
	db, err := sql.Open("sqlite3", GetInstance().DataBase.Path)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM " + GetInstance().DataBase.TableName + " WHERE id=" + strconv.Itoa(id))
	user := parser.User{}
	err = row.Scan(&user.Id, &user.UserName, &user.FullName, &user.City, &user.BirthDate.Time, &user.Department, &user.Gender, &user.ExperienceYears)
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	return &user, err
}

func GetAllUsers() []parser.ShortInfo {
	db, err := sql.Open("sqlite3", GetInstance().DataBase.Path)
	if err != nil {
		Log.Error(err)
		return nil
	}
	defer db.Close()
	rows, err := db.Query("SELECT Id, UserName FROM " + GetInstance().DataBase.TableName)
	if err != nil {
		Log.Error(err)
		return nil
	}

	defer rows.Close()

	var users []parser.ShortInfo
	for rows.Next() {
		user := parser.ShortInfo{}
		err := rows.Scan(&user.Id, &user.UserName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, user)
	}
	return users
}
