package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	Email				string 			`gorm:"size:32" json:"email"`
	Name				string 			`gorm:"size:32" json:"name"`
	Password			string 			`gorm:"size:255" json:"-"`
	AvatarURL			string			`gorm:"size:255" json:"avatarURL"`
}

const (
	UserRoleNoLogin	=	iota
	UserRoleAdmin
	UserRoleSignIn
)

func (u *User) Verify(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func GetUser(username string, password string) (*User, error) {
	user := &User{}
	// TODO
	err := GetDB().Where(&User{Name: username}).First(&user).Error
	return user, err
}
func (u *User) BeforeSave() (err error) {
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	u.Password = string(hash)
	return
}