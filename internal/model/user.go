package model

import (
	"time"
)

// User model
type User struct {
	ID         uint64     `json:"ID" xorm:"bigint(20) notnull autoincr pk 'ID'"`
	Name       string     `json:"name" xorm:"varchar(18) notnull 'name'"`
	NickName   string     `json:"nick_name" xorm:"varchar(25) notnull 'nick_name'"`
	Password   string     `json:"password" xorm:"varchar(20) notnull 'password'"` // use raw password temporarily
	Status     int8       `json:"status" xorm:"tinyint(4) notnull default(0) 'status'"`
	Email      string     `json:"email" xorm:"varchar(50) notnull 'email'"`
	Role       int8       `json:"role" xorm:"tinyint(4) notnull default(0) 'role'"`
	CreateTime *time.Time `json:"create_time,omitempty" xorm:"datetime created notnull 'create_time'"`
	UpdateTime *time.Time `json:"update_time,omitempty" xorm:"datetime updated notnull 'update_time'"`
	DeleteTime *time.Time `json:"delete_time,omitempty" xorm:"datetime 'delete_time'"`
}

// GetByName returns an user by 'name' if it exist.
// Not including 'ID' field.
func GetUserByName(name string) (bool, User, error) {
	db := orm.NewSession()
	defer db.Close()

	var user User
	has, err := db.Where("name = ?", name).Get(&user)

	return has, user, err
}

// Add adds user to `user` table.
func (u *User) Add() (bool, error) {
	db := orm.NewSession()
	defer db.Close()

	affected, err := db.InsertOne(u)

	return affected > 0, err
}

// Delete deletes an User.
func (u *User) Delete() (bool, error) {
	db := orm.NewSession()
	defer db.Close()

	affected, err := db.Delete(u)

	return affected > 0, err
}
