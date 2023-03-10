package config

import (
	"github.com/jinzhu/gorm"
)

// file นี้จะเอาไว้ return variable ที่ใช้เรียก database
// ส่วนการ interact กับ database จะใช้ไฟล์อื่น
var (
	db *gorm.DB
)

// สร้าง Connect Funcion
func Connect() {
	d, err := gorm.Open("mysql", "root:456154/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
