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
	UserRoleSignedIn
	UserRoleAdmin
	UserRoleSuperAdmin
)

func (u *User) Verify(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func GetUser(username string, password string) (*User, error) {
	user := &User{}
	err := GetDB().Where(&User{Name: username}).First(&user).Error
	return user, err
}

func GetUserRoleByUsername (username string) (*Correlation, error) {
	user := &User{}
	correlation := &Correlation{}
	err := GetDB().Where(&User{Name: username}).First(&user).Error
	
	if (err != nil) {
		return correlation, err;
	}
	error := GetDB().Where(&Correlation{ID1: user.ID, Type: CorrelationUserRole}).First(&correlation).Error
	return correlation, error

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