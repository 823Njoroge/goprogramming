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

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}
  
// Fake database- uses a slice from Course
var courses []Course

// Checks if the course is empty
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

// Main function to start the server
func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Initialize router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createCourse).Methods("POST")

	// Start server
	fmt.Println("Starting server on port 4000...")
	http.ListenAndServe(":4000", r)
}

// Serve the home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Hi! My name is Ian and I'm learning about APIs</h1>"))
}

// Get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Get a single course by ID
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get ID from the request URL


	// Loop through courses to find a matching ID
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given ID")
}

// Create a new course
func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
if r.Body==nil{
	json.NewEncoder(w).Encode("Please send some")
}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) // Decode request body into course object

	// Validate the request
	if course.IsEmpty() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please provide complete course details")
		return
	}

	// Assign a random ID and add the course to the database
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	// Return the newly created course
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)

    // Loop through courses to find the course with the given ID, remove it, and add the updated one
    for index, course := range courses {
        if course.CourseId == params["id"] {
            // Remove the course from the slice
            courses = append(courses[:index], courses[index+1:]...)
            
            // Decode the new course data from the request body
            var updatedCourse Course
            err := json.NewDecoder(r.Body).Decode(&updatedCourse)
            if err != nil {
                http.Error(w, "Invalid request payload", http.StatusBadRequest)
                return
            }
            
            updatedCourse.CourseId = params["id"]
            courses = append(courses, updatedCourse)
            
            json.NewEncoder(w).Encode(updatedCourse)
            return
        }
    }
    
    // Return a 404 error if the course was not found
    http.Error(w, "Course not found", http.StatusNotFound)
}

