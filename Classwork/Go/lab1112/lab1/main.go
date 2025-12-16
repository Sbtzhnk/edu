package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"size:64"`
	LastName  string `gorm:"size:64"`
	Email     string `gorm:"uniqueIndex;size:128"`
	Age       int
}

func main() {
	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=mysecretpassword dbname=todo_db port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                                                     // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	db.AutoMigrate(&User{})
	user := User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", Age: 30}
	result := db.Create(&user)
	if result.Error != nil {
		panic("failed to create user")
	}
}
