package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func ConnectDatabase() (db *gorm.DB) {
	dsn := "host=localhost user=postgres password=Namle311 dbname=book port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})
	return db
}

func Create(c *gin.Context) {
	var newProduct Product
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}
	db := ConnectDatabase()
	db.Create(&newProduct)
}

func ReadAll(c *gin.Context) {
	var products []Product
	db := ConnectDatabase()
	db.Find(&products)
	c.IndentedJSON(http.StatusOK, products)
}

func ReadOne(c *gin.Context) {
	var newProduct Product
	id := c.Param("id")
	db := ConnectDatabase()
	db.First(&newProduct, "code = ?", id)
	c.IndentedJSON(http.StatusOK, newProduct)
}

func main() {
	router := gin.Default()
	router.POST("/create", Create)
	router.GET("/read", ReadAll)
	router.GET("/read/:id", ReadOne)
	router.Run(":8000")

}
