package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCourses(c *gin.Context) {
	c.JSON(http.StatusOK, courses)
}

func getCourse(c *gin.Context) {
	id := c.Param("id")
	for _, co := range courses {
		if co.ID == id {
			c.JSON(http.StatusOK, co)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
}

func createCourse(c *gin.Context) {
	var newCourse Course
	if err := c.ShouldBindJSON(&newCourse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCourse.ID = generateID("CRS", len(courses))
	courses = append(courses, newCourse)
	saveCourses()
	c.JSON(http.StatusCreated, newCourse)
}

func updateCourse(c *gin.Context) {
	id := c.Param("id")
	var updated Course
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, co := range courses {
		if co.ID == id {
			updated.ID = id
			courses[i] = updated
			saveCourses()
			c.JSON(http.StatusOK, updated)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
}

func deleteCourse(c *gin.Context) {
	id := c.Param("id")
	for i, co := range courses {
		if co.ID == id {
			courses = append(courses[:i], courses[i+1:]...)
			saveCourses()
			c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
}
