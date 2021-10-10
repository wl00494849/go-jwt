package server

import (
	"database/sql"
	"go-jwt/model"

	"golang.org/x/crypto/bcrypt"
)

func UserRegister(user *model.User) {
	stmt, err := Db.Prepare("Insert users set  UserName = ?,Email= ? ,Password = ? ")
	Err.CheckError(err)
	stmt.Exec(user.UserName, user.Email, user.Password)
}

func UserLogin(data map[string]string) {

	pwd, err := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)
	Err.CheckError(err)

	row := Db.QueryRow("Select users form jwt where UserName = ? , Password = ?", data["UserName"], pwd)
	loginCheck(row)
}

func loginCheck(row *sql.Row) {
	var date map[string]string
	row.Scan(date)
	println("%+v", date)
}
