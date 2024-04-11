package main

import (
	"log"
	"main/employees"
	"main/model"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := model.Database()
	if err != nil {
		panic(err)
	}
	db.DB()
	r := gin.Default()
	r.GET("/getEmployees", employees.GetEmployees)
	r.GET("/getEmployee/:id", employees.GetEmployeesByID)
	r.POST("/insertEmployee", employees.CreateEmployee)
	r.PUT("/updateEmployee/:id", employees.UpdateEmployee)
	r.DELETE("/deleteEmployees/:id", employees.DeleteEmployee)

	log.Fatal(r.Run(":8989"))
}
