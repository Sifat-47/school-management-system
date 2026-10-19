package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getStats(c *gin.Context) {
	teacherCount := 0
	staffCount := 0

	for _, e := range employees {
		if e.Role == "Teacher" || e.Role == "teacher" {
			teacherCount++
		} else {
			staffCount++
		}
	}

	// Students per class
	classStats := []gin.H{}
	for _, cl := range classes {
		classStats = append(classStats, gin.H{
			"id":            cl.ID,
			"name":          cl.Name,
			"student_count": len(cl.StudentIDs),
			"teacher_id":    cl.TeacherID,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"total_students":  len(students),
		"total_employees": len(employees),
		"total_classes":   len(classes),
		"total_courses":   len(courses),
		"teachers":        teacherCount,
		"staff":           staffCount,
		"classes":         classStats,
	})
}
