package main

import (
	"09_http/auth"
	"09_http/class"
	"09_http/student"
	"09_http/teacher"
	"09_http/token"
	"09_http/user"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/dashboard", requireAuth(handleDashboard))
	http.HandleFunc("/class/", requireAuth(handlerClass))
	http.HandleFunc("/student/", requireAuth(handlerStudent))

	log.Println("Server is running. Logging page http://localhost:8020/login")
	log.Fatal(http.ListenAndServe(":8020", nil))
}

// renderTemplate виконує шаблон із зазначеними даними
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderTemplate(w, "login", nil)
	} else if r.Method == "POST" {
		r.ParseForm()

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		password := r.FormValue("password")
		if user.CheckAccess(id, password) {
			auth.SetLoggedIn(w, token.GenerateToken(id))
			http.Redirect(w, r, "/dashboard", http.StatusFound)
		} else {
			renderTemplate(w, "login", map[string]string{"ErrorMessage": "Invalid username or password"})
		}
	}
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	id, _ := auth.CheckLoggedIn(r)
	currentTeacher, err := teacher.Find(id)
	if err != nil {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}

	data := struct {
		Name    string
		Classes []*class.Class
	}{
		Name:    currentTeacher.Name,
		Classes: class.FindClassesByTeacher(id),
	}
	renderTemplate(w, "dashboard", data)
}

func handlerClass(w http.ResponseWriter, r *http.Request) {
	classIdString := strings.TrimPrefix(r.URL.Path, "/class/")
	classId, err := strconv.Atoi(classIdString)
	if err != nil {
		http.Error(w, "Invalid class ID", http.StatusBadRequest)
		return
	}
	classInfo := class.Find(classId)
	if classInfo == nil {
		http.Error(w, "Class not found", http.StatusNotFound)
		return
	}

	data := struct {
		ClassName   string
		TeacherName string
		Students    []*student.Student
	}{
		ClassName:   classInfo.Name,
		TeacherName: classInfo.Teacher.Name,
		Students:    student.FindStudentsByClass(classId),
	}
	renderTemplate(w, "class", data)
}

func handlerStudent(w http.ResponseWriter, r *http.Request) {
	studentIdString := strings.TrimPrefix(r.URL.Path, "/student/")
	studentId, err := strconv.Atoi(studentIdString)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}
	studentInfo := student.Find(studentId)
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
		ClassName:   class.Find(studentInfo.ClassId).Name,
		AvgScores:   studentInfo.AvgScores,
	}
	renderTemplate(w, "student", data)
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
