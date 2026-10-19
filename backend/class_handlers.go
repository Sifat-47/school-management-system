package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getClasses(c *gin.Context) {
	c.JSON(http.StatusOK, classes)
}

func getClass(c *gin.Context) {
	id := c.Param("id")
	for _, cl := range classes {
		if cl.ID == id {
			c.JSON(http.StatusOK, cl)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
}

func createClass(c *gin.Context) {
	var newClass Class
	if err := c.ShouldBindJSON(&newClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newClass.ID = generateID("CLS", len(classes))
	newClass.StudentIDs = []string{}
	classes = append(classes, newClass)
	saveClasses()
	c.JSON(http.StatusCreated, newClass)
}

func updateClass(c *gin.Context) {
	id := c.Param("id")
	var updated Class
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, cl := range classes {
		if cl.ID == id {
			updated.ID = id
			classes[i] = updated
			saveClasses()
			c.JSON(http.StatusOK, updated)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
}

func deleteClass(c *gin.Context) {
	id := c.Param("id")
	for i, cl := range classes {
		if cl.ID == id {
			classes = append(classes[:i], classes[i+1:]...)
			saveClasses()
			c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
}
