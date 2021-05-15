package model

import (
	"errors"
	"time"

	"github.com/Pengxn/go-xn/src/util/log"
)

// User model
type User struct {
	ID         uint64     `json:"ID" xorm:"bigint(20) notnull autoincr pk 'ID'"`
	Name       string     `json:"name" xorm:"varchar(18) notnull 'name'"`
	NickName   string     `json:"nick_name" xorm:"varchar(25) notnull 'nick_name'"`
	Password   string     `json:"password" xorm:"varchar(20) notnull 'password'"`
	Status     int8       `json:"status" xorm:"tinyint(4) notnull default(0) 'status'"`
	Email      string     `json:"email" xorm:"varchar(50) notnull 'email'"`
	Role       int8       `json:"role" xorm:"tinyint(4) notnull default(0) 'role'"`
	CreateTime *time.Time `json:"create_time,omitempty" xorm:"datetime created notnull default('0000-00-00 00:00:00') 'create_time'"`
	UpdateTime *time.Time `json:"update_time,omitempty" xorm:"datetime updated notnull default('0000-00-00 00:00:00') 'update_time'"`
	DeleteTime *time.Time `json:"delete_time,omitempty" xorm:"datetime  default('0000-00-00 00:00:00') 'delete_time'"`
}

// GetByName returns an user by 'name' if it exist.
// Not including 'ID' field.
func (u *User) GetByName(name string) (User, error) {
	db := orm.NewSession()
	defer db.Close()

	has, err := db.Omit("ID").
		Where("name = ?", name).
		Get(u)
	if err != nil || !has {
		log.Errorf("User GetByName throw error: %s", err)
		return *u, errors.New("Get user data error")
	}

	return *u, nil
}

// Add adds user to `user` table.
func (u *User) Add() bool {
	db := orm.NewSession()
	defer db.Close()

	affected, err := db.InsertOne(u)
	if err != nil {
		log.Errorf("User database add throw error: %s, param: %+v", err, u)
	}

	return affected > 0
}

// Delete deletes an User.
func (u *User) Delete() bool {
	db := orm.NewSession()
	defer db.Close()

	affected, err := db.Delete(u)
	if err != nil {
		log.Errorf("User database delete throw error: %s, param: %+v", err, u)
	}

	return affected > 0
}
