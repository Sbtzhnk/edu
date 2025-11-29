package singl

import "fmt"

var instance *dbConnection

type dbConnection struct {
	host    string
	port    int
	dbName  string
	dbLogin string
	dbPass  string
}

func init() {
	GetInstance()
	fmt.Print("Сработал init \n")
}

func GetInstance() *dbConnection {
	if instance == nil {
		instance = &dbConnection{"localhost", 5436, "to-do", "to-do", "123"}
	}
	return instance
}
