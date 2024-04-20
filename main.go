package main

import (
	"database/sql"
	"fmt"
	"mini-project-sanbercode/controllers"
	"mini-project-sanbercode/databases"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// Load .env file
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("Error: Could not connect to the database")
		panic(err)
	} else {
		fmt.Println("Connected to the database")
	}

	databases.DbMigrate(DB)
	defer DB.Close()

	//Router GIN
	router := gin.Default()
	router.GET("/person", controllers.GetAllPerson)
	router.POST("/person", controllers.InsertPerson)
	router.PUT("/person/:id", controllers.UpdatePerson)
	router.DELETE("/person/:id", controllers.DeletePerson)

	router.Run(":8080")
}
