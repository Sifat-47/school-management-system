package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getEmployees(c *gin.Context) {
	c.JSON(http.StatusOK, employees)
}

func getEmployee(c *gin.Context) {
	id := c.Param("id")
	for _, e := range employees {
		if e.ID == id {
			c.JSON(http.StatusOK, e)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

func createEmployee(c *gin.Context) {
	var newEmployee Employee
	if err := c.ShouldBindJSON(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newEmployee.ID = generateID("EMP", len(employees))
	employees = append(employees, newEmployee)
	saveEmployees()
	c.JSON(http.StatusCreated, newEmployee)
}

func updateEmployee(c *gin.Context) {
	id := c.Param("id")
	var updated Employee
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, e := range employees {
		if e.ID == id {
			updated.ID = id
			employees[i] = updated
			saveEmployees()
			c.JSON(http.StatusOK, updated)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

func deleteEmployee(c *gin.Context) {
	id := c.Param("id")
	for i, e := range employees {
		if e.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			saveEmployees()
			c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}
