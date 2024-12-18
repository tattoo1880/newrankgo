package model

import (
	"gorm.io/gorm"
	. "gotest/database"
	"log"
)

// Xhs struct
type Xhs struct {
	gorm.Model
	UserId   string `gorm:"primaryKey" json:"userId"`
	UserName string `json:"userName"`
	//todo 可以为空的title
	Title        string `json:"title"`
	CollectCount int    `json:"collectCount"`
	CommentCount int    `json:"commentCount"`
	ShareCount   int    `json:"shareCount"`
	Fans         int    `json:"fans"`
	LikeCount    int    `json:"likeCount"`
	Desc         string `json:"desc"`
	Avatar       string `json:"userHeadUrl"`
	PublicTime   string `json:"publicTime"`
}

// todo crud
// todo 1. create
func (xhs *Xhs) Create() error {
	return MyDB.Create(&xhs).Error
}

// todo 2. readAll
func (xhs *Xhs) ReadAll() ([]Xhs, error) {
	var xhsList []Xhs
	err := MyDB.Find(&xhsList).Error
	return xhsList, err
}

// deleteAll
func (xhs *Xhs) DeleteAll() error {
	var xhsList []Xhs
	if err := MyDB.Find(&xhsList).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := MyDB.Where("1=1").Delete(&xhsList).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
