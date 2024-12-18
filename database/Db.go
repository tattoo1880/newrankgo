package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MyDB *gorm.DB

func init() {

	dsn := "root:qwerty7788421@tcp(localhost:3306)/newrank?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	MyDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
		return
	}

}
