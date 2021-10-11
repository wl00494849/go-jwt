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

func UserRegister(data map[string]string) {
	pwd, _ := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)

	user := &model.User{
		UserName: data["UserName"],
		Email:    data["Email"],
		Password: pwd,
	}

	stmt, err := Db.Prepare("Insert users set  UserName = ?,Email= ? ,Password = ? ")
	Err.CheckError(err)
	stmt.Exec(user.UserName, user.Email, user.Password)
}

func GetToken(data map[string]string) (string, error) {
	user, ok := pwdCheck(data)

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

func GetUserInfo(cookie string) *model.User {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil
	}

	claims := token.Claims.(*jwt.StandardClaims)
	user := &model.User{}

	row := Db.QueryRow("Select * from users where Id = ?", claims.Issuer)
	row.Scan(&user.Id, &user.UserName, &user.Email, &user.Password)

	if user.Id == 0 {
		return nil
	}

	return user
}

func pwdCheck(data map[string]string) (*model.User, bool) {
	user := getUserLoginData(data)

	if user == nil {
		return nil, false
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["Password"])); err != nil {
		return nil, false
	}

	return user, true
}

func getUserLoginData(data map[string]string) *model.User {
	var user = &model.User{}

	row := Db.QueryRow("Select* from users where UserName = ? ", data["UserName"])
	row.Scan(&user.Id, &user.UserName, &user.Email, &user.Password)

	if user.Id == 0 {
		return nil
	}

	return user
}
