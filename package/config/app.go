package config

import (
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// file นี้จะเอาไว้ return variable ที่ใช้เรียก database
// ส่วนการ interact กับ database จะใช้ไฟล์อื่น
var (
	db *gorm.DB
)

// สร้าง Connect Funcion
// https://github.com/go-sql-driver/mysql#dsn-data-source-name
//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
func Connect() {
	// d, err := gorm.Open("mysql", "root:456154/bookstore?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	panic(err)
	// }
	// db = d
	dsn := "root:456154@tcp(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
