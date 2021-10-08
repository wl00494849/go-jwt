package server

import "go-jwt/model"

func UserRegister(user *model.User) {
	stmt, _ := Db.Prepare("Insert users set Id=? , UserName=? , Email = ? Passward = ?")
	stmt.Exec(user.Id, user.UserName, user.Email, user.Password)
}
