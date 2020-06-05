package database

import (
	"database/sql"
	"fmt"
)


func InitDB() *sql.DB {
	connectionString := "homestead:secret@tcp(192.168.10.10:3306)/northwind"
	databaseConnection, err := sql.Open("mysql",connectionString)

	if err != nil {
		fmt.Println("Error de conexión con la BD!!")
		panic(err.Error()) // Error handling
	}
	return databaseConnection
}