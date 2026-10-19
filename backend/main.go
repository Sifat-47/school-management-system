package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	loadData()

	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "School Management System API is running!",
		})
	})

	// ========== STUDENT ROUTES ==========
	r.GET("/api/students", getStudents)
	r.GET("/api/students/:id", getStudent)
	r.POST("/api/students", createStudent)
	r.PUT("/api/students/:id", updateStudent)
	r.DELETE("/api/students/:id", deleteStudent)

	// ========== EMPLOYEE ROUTES ==========
	r.GET("/api/employees", getEmployees)
	r.GET("/api/employees/:id", getEmployee)
	r.POST("/api/employees", createEmployee)
	r.PUT("/api/employees/:id", updateEmployee)
	r.DELETE("/api/employees/:id", deleteEmployee)

	// ========== CLASS ROUTES ==========
	r.GET("/api/classes", getClasses)
	r.GET("/api/classes/:id", getClass)
	r.POST("/api/classes", createClass)
	r.PUT("/api/classes/:id", updateClass)
	r.DELETE("/api/classes/:id", deleteClass)

	// ========== COURSE ROUTES ==========
	r.GET("/api/courses", getCourses)
	r.GET("/api/courses/:id", getCourse)
	r.POST("/api/courses", createCourse)
	r.PUT("/api/courses/:id", updateCourse)
	r.DELETE("/api/courses/:id", deleteCourse)

	// ========== STATS ==========
	r.GET("/api/stats", getStats)

	r.Run(":8080")
}
