package server

import (
	"errors"
	"go-jwt/model"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "test"

func UserRegister(user *model.User) {
	stmt, err := Db.Prepare("Insert users set  UserName = ?,Email= ? ,Password = ? ")
	Err.CheckError(err)
	stmt.Exec(user.UserName, user.Email, user.Password)
}

func UserLogin(data map[string]string) (string, error) {
	user, ok := loginCheck(data)

	if !ok {
		err := errors.New("user not found")
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.Id),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))
	return token, err
}

func loginCheck(data map[string]string) (*model.User, bool) {
	var user = &model.User{}

	row := Db.QueryRow("Select* from users where UserName = ? ", data["UserName"])
	row.Scan(&user.Id, &user.UserName, &user.Email, &user.Password)

	if user.Id == 0 {
		return nil, false
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["Password"])); err != nil {
		return nil, false
	}

	return user, true
}
