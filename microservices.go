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
	dsn := "host=localhost user=postgres password=2306 dbname=employee port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Postgres Database Connected Successfully")
	return db
}

type Employee struct {
	Id       uint   `gorm:"primary_key;auto_increment"json:"-"`
	Name     string `gorm:"type:varchar(255)"json:"name"`
	Salary   string `gorm:"type:varchar(255) "json:"salary"`
	Location string `gorm:"type:varchar(255)"json:"location"`
}

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"Project Name":  "Employee Details",
		"Database Name": "PostgreSQL",
	})
}

func CreateData(c *gin.Context) {
	// Initlize Database
	db := Conn()
	var data Employee
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
		"id":      data.Id,
		"message": "data created",
		"data":    &data,
	})
	return
}

func UpdateData(c *gin.Context) {
	// Initlize Database
	db := Conn()
	id := c.Param("id")
	log.Println("Id is :: ", id)
	var data Employee
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
	result := db.Model(&data).Where("id = ?", id).Updates(Employee{Name: data.Name, Salary: data.Salary, Location: data.Location})
	rows := result.RowsAffected
	if rows != 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data doesn't exits"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		" id":     id,
		"message": "data updated",
		"data":    &data,
	})
	return
}

func DeleteData(c *gin.Context) {
	// Initlize Database
	db := Conn()
	id := c.Param("id")
	log.Println("Id is :: ", id)
	result := db.Delete(&Employee{}, id)
	rows := result.RowsAffected
	log.Println("Rows :: ", rows)
	if rows != 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data doesn't exits"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "data deleted",
	})
	return
}

func SingleEmployee(c *gin.Context) {
	// Initlize Database
	db := Conn()
	id := c.Param("id")
	log.Println("Id is :: ", id)
	var data Employee
	result := db.Raw("select id,name,salary,location FROM employees where id = ?", id).Scan(&data)
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

func AllEmployee(c *gin.Context) {
	// Initlize Database
	db := Conn()
	var data Employee
	var employeedata []Employee
	result, err := db.Model(&Employee{}).Select("id, name, salary, location").Rows()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No Data Found"})
		return
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&data.Id, &data.Name, &data.Salary, &data.Location)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Server error"})
			return
		}
		employeedata = append(employeedata, data)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": employeedata,
	})
	return
}

func RequestHandler() {
	r := gin.Default()
	r.GET("/", HomePage)
	r.GET("/employee/:id", SingleEmployee)
	r.GET("/all", AllEmployee)
	r.POST("/create", CreateData)
	r.PUT("/update/:id", UpdateData)
	r.DELETE("/delete/:id", DeleteData)
	log.Fatal(http.ListenAndServe(":8000", r))
	r.Run()
}

func main() {
	Conn()
	RequestHandler()
}
