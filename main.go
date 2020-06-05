package main

import (
	"fmt"
	"golangRestApi/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	databaseConnection := database.InitDB()





	
	defer databaseConnection.Close() // Con el defer la conexión se cerrará cuando finalice la función contenedora
	
	fmt.Println(databaseConnection)


}
