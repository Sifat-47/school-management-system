package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /api/students → Get all students
func getStudents(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

// GET /api/students/:id → Get one student
func getStudent(c *gin.Context) {
	id := c.Param("id")

	for _, s := range students {
		if s.ID == id {
			c.JSON(http.StatusOK, s)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

// POST /api/students → Add new student
func createStudent(c *gin.Context) {
	var newStudent Student

	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newStudent.ID = generateID("STU", len(students))
	students = append(students, newStudent)
	saveStudents()

	c.JSON(http.StatusCreated, newStudent)
}

// PUT /api/students/:id → Update student
func updateStudent(c *gin.Context) {
	id := c.Param("id")
	var updatedStudent Student

	if err := c.ShouldBindJSON(&updatedStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, s := range students {
		if s.ID == id {
			updatedStudent.ID = id
			students[i] = updatedStudent
			saveStudents()
			c.JSON(http.StatusOK, updatedStudent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

// DELETE /api/students/:id → Delete student
func deleteStudent(c *gin.Context) {
	id := c.Param("id")

	for i, s := range students {
		if s.ID == id {
			students = append(students[:i], students[i+1:]...)
			saveStudents()
			c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}
