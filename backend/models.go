package main

type Student struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	ClassID   string `json:"class_id"`
}

type Employee struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	Subject string `json:"subject"`
}

type Class struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	TeacherID  string   `json:"teacher_id"`
	StudentIDs []string `json:"student_ids"`
}

type Course struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TeacherID   string `json:"teacher_id"`
}
