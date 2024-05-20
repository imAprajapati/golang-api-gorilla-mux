package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for Course - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	WebSite  string `json:"website"`
}

// fake db

var courses []Course

// Middelware helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Hello, World!")
	courses = append(courses, Course{CourseId: "1", CourseName: "Java", CoursePrice: 100, Author: &Author{FullName: "John Doe", WebSite: "www.johndoe.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Python", CoursePrice: 200, Author: &Author{FullName: "Jane Doe", WebSite: "www.janedoe.com"}})
	// create a new router using gorialla/mux
	r := mux.NewRouter()
	r.HandleFunc("/", HomePage).Methods("GET")
	r.HandleFunc("/courses", GetCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")
	r.HandleFunc("/course", CreateCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteCourse).Methods("DELETE")

	// start the server on port 8000
	http.ListenAndServe(":8000", r)
}

// Controller - file

// Server Home Page
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the HomePage</h1>"))
}

// Get all courses
func GetCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Get One Course
func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range courses {
		if item.CourseId == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Course{})
}

// Create a new Course
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateCourse")
	w.Header().Set("Content-Type", "application/json")

	// What if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send a request body")
		return
	}

	// what about {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send a request body")
		return
	}

	// check course name duplication
	for _, item := range courses {
		if item.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course name already exists")
			return
		}
	}

	// create Unique ID
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

// Update a Course
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// if not found
	json.NewEncoder(w).Encode("Course not found")
}

// Delete a Course
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted")
			break
		}
	}
	// if not found
	json.NewEncoder(w).Encode("Course not found")
}
