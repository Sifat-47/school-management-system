package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	RegistrationID string `json:"registration_id"`
	Password       string `json:"password"`
}

// POST /api/student/login
func studentLogin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	for _, s := range students {
		if s.ID == req.RegistrationID && s.Password == req.Password {
			// Login successful – return student data (without password)
			c.JSON(http.StatusOK, gin.H{
				"message": "Login successful",
				"student": gin.H{
					"id":         s.ID,
					"first_name": s.FirstName,
					"last_name":  s.LastName,
					"age":        s.Age,
					"class_id":   s.ClassID,
				},
			})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Registration ID or Password"})
}
