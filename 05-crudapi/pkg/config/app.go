package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// import (
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
//   )
//   func main() {
// 	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
// 	dsn := "user:pass@tcp(localhost:3306)/simpleapi?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//   }

// declear global variable for databse connection
var (
	db *gorm.DB
)

//  connecting with database using gorm.. gorm can make table by self if there is no table availabe in the database
func Connect() {
	d, err := gorm.Open("mysql", "root:javed@123@/crudapi_db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("connection failed...%v", err)
	}
	db = d
}

//  send database conections to the requester...
func GetDB() *gorm.DB {
	return db
}
