package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "root:root1234@tcp(localhost:3306)/epc"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error DB : %v\n", err)
	}

	fmt.Printf("DB : %v\n", DB)
	// TODO stop the application when db connection refused
}
