package main

import (
	"fmt"
	m "go-android-postgresql/customer-go/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var e error

func main() {
	db, e = gorm.Open("postgres", "user=postgres dbname=postgres password=pratama sslmode=disable")
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Connection Established!")
	}

	defer db.Close()

	db.SingularTable(true)
	db.AutoMigrate(m.Customer{})

	r := gin.Default()
	// Get customers
	r.GET("/customers", getCustomers)
	// Get customer by name
	// r.GET("/customers/test/:phone_number/customer", getCustomerByPhoneNumber)
	r.GET("/customers/:name", getCustomersByName)
	// Insert customer
	r.POST("/customers/", insertCustomer)
	r.Run(":8090")
}

// Get customers
func getCustomers(c *gin.Context) {
	var customers []m.Customer
	if e := db.Find(&customers).Error; e != nil {
		c.AbortWithStatus(404)
		fmt.Println(e)
	} else {
		c.JSON(200, customers)
	}
}

// Get customer by id
func getCustomerById(c *gin.Context) {
	var customer m.Customer
	id := c.Params.ByName("id")
	if e := db.Where("id = ?", id).First(&customer).Error; e != nil {
		c.AbortWithStatus(404)
		fmt.Println(e)
	} else {
		c.JSON(200, customer)
	}
}

// Get customers by name
func getCustomersByName(c *gin.Context) {
	var customers []m.Customer
	name := c.Params.ByName("name")
	if e := db.Where("name = ?", name).Find(&customers).Error; e != nil {
		c.AbortWithStatus(404)
		fmt.Println(e)
	} else {
		c.JSON(200, customers)
	}
}

func getCustomerByPhoneNumber(c *gin.Context) {
	var customer m.Customer
	phoneNumber := c.Params.ByName("phone_number")
	if e := db.Where("phone_number = ?", phoneNumber).First(&customer).Error; e != nil {
		c.AbortWithStatus(404)
		fmt.Println(e)
	} else {
		c.JSON(200, customer)
	}
}

// Insert customer
func insertCustomer(c *gin.Context) {
	var customer m.Customer
	c.BindJSON(&customer)
	db.Create(&customer)
	c.JSON(200, customer)
}
