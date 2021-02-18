package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Conn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "2306"
	dbName := "fruit"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Panicln("Error in DB Connection ", err.Error())
	}
	return db
}

type Fruit struct {
	Fid          int64  `json:"fid"`
	Fruitname    string `json:"fruitname"`
	Rate         string  `json:"rateperkg"`
	//Rank         string `json:"rank"`
	Benefits     string `json:"benefits"`
	
}

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"Project Name":  "Fruits Database",
		"Done By": "GAGAN",
	})
}

func CreateFruit(c *gin.Context) {
	// Initlize Database
	db := Conn()
	defer db.Close()
	var data Fruit
	//Reading Request Body
	RequestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid Parameter"})
		return
	}
	//Convert Request Body into Json Formate
	json.Unmarshal(RequestBody, &data)
	log.Println("Data: ", &data)
	result, err := db.ExecContext(c, "insert into fruitset (fruitname,rateperkg,benefits) values (?,?,?)", data.Fruitname, data.Rate, data.Benefits)
	result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid Parameter"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{

		"message": "Done",
	})
	return
}

func UpdateFruit(c *gin.Context) {
	// Initlize Database
	db := Conn()
	defer db.Close()
	fid := c.Param("fid")
	var data Fruit
	//Reading Request Body
	RequestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid Parameter"})
		return
	}
	//Convert Request Body into Json Formate
	json.Unmarshal(RequestBody, &data)
	log.Println("Update Fruit : ", &data)
	result, err := db.ExecContext(c, "update fruitset set fruitname = ? , rateperkg = ? , benefits = ? where fid = ? ", data.Fruitname, data.Rate, data.Benefits, fid)
	rows, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid Parameter"})
		return
	}
	if rows != 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{

		"message": "updated fruitset",
	})
	return
}

func DeleteFruit(c *gin.Context) {
	// Initlize Database
	db := Conn()
	defer db.Close()
	// Store Id
	fid := c.Param("fid")
	log.Println("FId is : ", fid)
	result, err := db.ExecContext(c, "delete from fruitset where fid = ?", fid)
	rows, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid Parameter"})
		return
	}
	if rows != 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{

		"message": "fruitset deleted",
	})
	return
}

func AllFruits(c *gin.Context) {
	var fruit Fruit
	var fruitset []Fruit
	// Initlize Database
	db := Conn()
	defer db.Close()
	result, err := db.Query("select *from fruitset")
	defer result.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server error"})
		return
	}
	for result.Next() {
		err := result.Scan(&fruit.Fid, &fruit.Fruitname, &fruit.Rate, &fruit.Benefits)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Server error"})
			return
		}
		fruitset = append(fruitset, fruit)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": fruitset,
	})
	return
}

func SingleFruit(c *gin.Context) {
	// Initlize Database
	db := Conn()
	defer db.Close()
	// Store Id
	fid := c.Param("fid")
	log.Println("Fid is :", fid)
	result, err := db.Query("select *from fruitset where fid = ?", fid)
	defer result.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server error"})
		return
	}
	for result.Next() {
		var fruit Fruit
		err := result.Scan(&fruit.Fid, &fruit.Fruitname, &fruit.Rate, &fruit.Benefits)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Server error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"FruitsData": fruit,
		})
	}
	return
}

func RequestHandler() {
	r := gin.Default()
	r.GET("/", HomePage)
	r.GET("/allfruits", AllFruits)
	r.GET("/singlefruit/:fid", SingleFruit)
	r.POST("/create", CreateFruit)
	r.PUT("/update/:fid", UpdateFruit)
	r.DELETE("/delete/:fid", DeleteFruit)
	r.Run(":4000")
}

func main() {
	Conn()
	RequestHandler()
}