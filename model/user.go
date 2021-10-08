package model

type User struct {
	Id       int
	UserName string
	Email    string
	Password []byte
}
