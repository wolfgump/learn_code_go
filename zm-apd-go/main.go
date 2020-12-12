package main

import (
	"log"
	"os"
	"zm-apd-go/application/service"
	"zm-apd-go/infrasturcture/persiterence"
	"zm-apd-go/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}
func main() {
	host := os.Getenv("DB_HOST")
	driver := os.Getenv("DB_DRIVER")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dbOperation, dbError := persiterence.NewDbOperation(driver, user, password, port, host, name)
	dbOperation.AutoMigrate()
	if dbError != nil {
		panic(dbError)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	repo := persiterence.NewProductMapper(dbOperation.Db)
	app := service.NewProductApp(repo)
	products := interfaces.NewProducts(*app)

	r.POST("/admin/project/create", products.Create)
	r.POST("/admin/project/update", products.Update)
	r.GET("/admin/project/detail", products.Query)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
