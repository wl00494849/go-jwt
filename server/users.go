package server

import "go-jwt/model"

func UserRegister(user *model.User) {
	stmt, err := Db.Prepare("Insert users set  UserName=?,Email = ?,Password = ?")
	Err.CheckError(err)
	stmt.Exec(user.UserName, user.Email, user.Password)
}
