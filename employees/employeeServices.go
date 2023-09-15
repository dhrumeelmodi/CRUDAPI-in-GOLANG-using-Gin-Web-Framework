package employees

import (
	"log"
	"net/http"

	"main/model"

	"github.com/gin-gonic/gin"
)

type NewEmployee struct {
	Name   string `json:"name" binding:"required"`
	Salary string `json:"salary" binding:"required"`
}

type EmployeeUpdate struct {
	Name   string `json:"name"`
	Salary string `json:"salary"`
}

func GetEmployees(c *gin.Context) {
	var emp []model.Employees
	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}
	if err := db.Find(&emp).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, emp)
}

func GetEmployeesByID(c *gin.Context) {
	var emp model.Employees
	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}
	if err := db.Where("id= ?", c.Param("id")).First(&emp).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, emp)
}

func CreateEmployee(c *gin.Context) {
	var emp NewEmployee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newEmp := model.Employees{Name: emp.Name, Salary: emp.Salary}
	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}
	if err := db.Create(&newEmp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newEmp)
}

func UpdateEmployee(c *gin.Context) {
	var emp model.Employees
	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}
	if err := db.Where("id = ?", c.Param("id")).First(&emp).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found!"})
		return
	}
	var updateEmp EmployeeUpdate
	if err := c.ShouldBindJSON(&updateEmp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Model(&emp).Updates(model.Employees{
		Name:   updateEmp.Name,
		Salary: updateEmp.Salary,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, emp)
}

func DeleteEmployee(c *gin.Context) {
	var emp model.Employees
	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}
	if err := db.Where("id = ?", c.Param("id")).First(&emp).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found!"})
		return
	}
	if err := db.Delete(&emp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}
