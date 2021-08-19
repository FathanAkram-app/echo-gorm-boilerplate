package models

import (
	"depmod/db"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       uint `gorm:"primarykey,AUTO_INCREMENT"`
	Username string
	Email    string
	Password string
}

// func FetchAllUser() Response {

// 	var users []Users
// 	var res Response

// 	con := db.CreateCon()

// 	con.Find(&users)

// 	res.Status = http.StatusOK
// 	res.Message = "success"
// 	res.Data = users

// 	return res

// }

func RegisterUser(username string, email string, password string) (Response, error) {
	var res Response
	hashed, _ := HashPassword(password)

	user := Users{
		Username: username,
		Email:    email,
		Password: hashed}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = user

	err := validation.Validate(user)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "failed"
		return res, err
	}
	con := db.CreateCon()

	con.Create(&user)

	return res, nil

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
