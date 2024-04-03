package api

import (
	"09_http/auth"
	"09_http/entites"
	"09_http/template"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Init() {
	fmt.Println("API package initialized")

	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/dashboard", requireAuth(handleDashboard))
	http.HandleFunc("/class/", requireAuth(handlerClass))
	http.HandleFunc("/student/", requireAuth(handlerStudent))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		template.RenderTemplate(w, "login", nil)
	} else if r.Method == "POST" {
		r.ParseForm()

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		password := r.FormValue("password")
		if auth.CheckAccess(id, password) {
			auth.SetLoggedIn(w, auth.GenerateToken(id))
			http.Redirect(w, r, "/dashboard", http.StatusFound)
		} else {
			template.RenderTemplate(w, "login", map[string]string{"ErrorMessage": "Invalid username or password"})
		}
	}
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	id, _ := auth.CheckLoggedIn(r)
	teacher, err := entites.FindTeacher(id)
	if err != nil {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}

	data := struct {
		Name    string
		Classes []*entites.Class
	}{
		Name:    teacher.Name,
		Classes: entites.FindClassesByTeacher(id),
	}
	template.RenderTemplate(w, "dashboard", data)
}

func handlerClass(w http.ResponseWriter, r *http.Request) {
	classIdString := strings.TrimPrefix(r.URL.Path, "/class/")
	classId, err := strconv.Atoi(classIdString)
	if err != nil {
		http.Error(w, "Invalid class ID", http.StatusBadRequest)
		return
	}
	classInfo := entites.FindClass(classId)
	if classInfo == nil {
		http.Error(w, "Class not found", http.StatusNotFound)
		return
	}

	data := struct {
		ClassName   string
		TeacherName string
		Students    []*entites.Student
	}{
		ClassName:   classInfo.Name,
		TeacherName: classInfo.Teacher.Name,
		Students:    entites.FindStudentsByClass(classId),
	}
	template.RenderTemplate(w, "class", data)
}

func handlerStudent(w http.ResponseWriter, r *http.Request) {
	studentIdString := strings.TrimPrefix(r.URL.Path, "/student/")
	studentId, err := strconv.Atoi(studentIdString)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}
	studentInfo := entites.FindStudent(studentId)
	if studentInfo == nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	data := struct {
		StudentName string
		ClassName   string
		AvgScores   map[string]float64
	}{
		StudentName: studentInfo.Name,
		ClassName:   entites.FindClass(studentInfo.ClassId).Name,
		AvgScores:   studentInfo.AvgScores,
	}
	template.RenderTemplate(w, "student", data)
}

func requireAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, loggedIn := auth.CheckLoggedIn(r)
		if !loggedIn {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		handler(w, r)
	}
}
