package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conn() *gorm.DB {
	dsn := "host=localhost user=postgres password=2306 dbname=PostgreSql port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Postgres Database Connected Successfully")
	return db
}

type Article struct {
	Id          uint   `gorm:"primary_key;auto_increment"json:"-"`
	Title       string `gorm:"type:varchar(255)"json:"title"`
	Description string `gorm:"type:varchar(255)"json:"description"`
	Author      string `gorm:"type:varchar(255)"json:"author"`
}

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"Project Name":  "Article Management System",
		"Database Name": "PostgreSQL",
	})
}

func CreateArticle(c *gin.Context) {
	// Initlize Database
	db := Conn()
	var data Article
	//Reading Request Body
	RequestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Bad Gatway"})
		return
	}
	//Convert Request Body into Json Formate
	json.Unmarshal(RequestBody, &data)
	log.Println("Data : ", &data)
	result := db.Create(&data)
	rows := result.RowsAffected
	if rows != 1 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid Parameters"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"article id": data.Id,
		"message":    "article saved",
		"data":       &data,
	})
	return
}

func UpdateArticle(c *gin.Context) {
	// Initlize Database
	db := Conn()
	id := c.Param("id")
	log.Println("Id is :: ", id)
	var data Article
	//Reading Request Body
	RequestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Bad Gatway"})
		return
	}
	//Convert Request Body into Json Formate
	json.Unmarshal(RequestBody, &data)
	log.Println("Data : ", &data)
	result := db.Model(&data).Where("id = ?", id).Updates(Article{Title: data.Title, Description: data.Description, Author: data.Author})
	rows := result.RowsAffected
	if rows != 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data doesn't exits"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"article id": id,
		"message":    "article updated",
		"data":       &data,
	})
	return
}

func DeleteArticle(c *gin.Context) {
	// Initlize Database
	db := Conn()
	id := c.Param("id")
	log.Println("Id is :: ", id)
	result := db.Delete(&Article{}, id)
	rows := result.RowsAffected
	log.Println("Rows :: ", rows)
	if rows != 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data doesn't exits"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"article id": id,
		"message":    "article deleted",
	})
	return
}

func SingleArticle(c *gin.Context) {
	// Initlize Database
	db := Conn()
	id := c.Param("id")
	log.Println("Id is :: ", id)
	var data Article
	result := db.Raw("select id,title,description,author FROM articles where id = ?", id).Scan(&data)
	rows := result.RowsAffected
	if rows != 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data doesn't exits"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": &data,
	})
	return
}

func AllArticle(c *gin.Context) {
	// Initlize Database
	db := Conn()
	var data Article
	var articledata []Article
	result, err := db.Model(&Article{}).Select("id, title, description, author").Rows()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No Data Found"})
		return
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&data.Id, &data.Title, &data.Description, &data.Author)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Server error"})
			return
		}
		articledata = append(articledata, data)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": articledata,
	})
	return
}

func RequestHandler() {
	r := gin.Default()
	r.GET("/", HomePage)
	r.GET("/article/:id", SingleArticle)
	r.GET("/all", AllArticle)
	r.POST("/create", CreateArticle)
	r.PUT("/update/:id", UpdateArticle)
	r.DELETE("/delete/:id", DeleteArticle)
	r.Run()
}

func main() {
	Conn()
	RequestHandler()
}
