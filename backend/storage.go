package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var students []Student
var employees []Employee
var classes []Class
var courses []Course

func loadData() {
	loadStudents()
	loadEmployees()
	loadClasses()
	loadCourses()
}

func saveData() {
	saveStudents()
	saveEmployees()
	saveClasses()
	saveCourses()
}

// ---------- Students ----------
func loadStudents() {
	file, err := os.ReadFile("../data/students.json")
	if err != nil {
		students = []Student{}
		return
	}
	json.Unmarshal(file, &students)
}

func saveStudents() {
	data, _ := json.MarshalIndent(students, "", "  ")
	os.WriteFile("../data/students.json", data, 0644)
}

// ---------- Employees ----------
func loadEmployees() {
	file, err := os.ReadFile("../data/employees.json")
	if err != nil {
		employees = []Employee{}
		return
	}
	json.Unmarshal(file, &employees)
}

func saveEmployees() {
	data, _ := json.MarshalIndent(employees, "", "  ")
	os.WriteFile("../data/employees.json", data, 0644)
}

// ---------- Classes ----------
func loadClasses() {
	file, err := os.ReadFile("../data/classes.json")
	if err != nil {
		classes = []Class{}
		return
	}
	json.Unmarshal(file, &classes)
}

func saveClasses() {
	data, _ := json.MarshalIndent(classes, "", "  ")
	os.WriteFile("../data/classes.json", data, 0644)
}

// ---------- Courses ----------
func loadCourses() {
	file, err := os.ReadFile("../data/courses.json")
	if err != nil {
		courses = []Course{}
		return
	}
	json.Unmarshal(file, &courses)
}

func saveCourses() {
	data, _ := json.MarshalIndent(courses, "", "  ")
	os.WriteFile("../data/courses.json", data, 0644)
}

func generateID(prefix string, count int) string {
	return fmt.Sprintf("%s%03d", prefix, count+1)
}
