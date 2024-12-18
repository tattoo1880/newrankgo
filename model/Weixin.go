package model

import (
	"gorm.io/gorm"
	"gotest/database"
	"log"
)

type Weixin struct {
	//gorm.Model
	gorm.Model
	Uid          int    `gorm:"primaryKey" json:"uid"`
	UserName     string `json:"userName"`
	Title        string `json:"title"`
	ClickCount   int    `json:"clickCount"`
	ShareCount   int    `json:"shareCount"`
	WatchCount   int    `json:"watchCount"`
	Avatar       string `json:"userHeadUrl"`
	LikeCount    int    `json:"likeCount"`
	CommentCount int    `json:"commentCount"`
}

func (weixin *Weixin) Create() error {
	return database.MyDB.Create(&weixin).Error
}

func (weixin *Weixin) ReadAll() ([]Weixin, error) {
	var weixinList []Weixin
	err := database.MyDB.Find(&weixinList).Error
	return weixinList, err
}

func (weixin *Weixin) DeleteAll() error {
	var weixinList []Weixin
	if err := database.MyDB.Find(&weixinList).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := database.MyDB.Where("1=1").Delete(&weixinList).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
