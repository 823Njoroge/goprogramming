package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

// FAKE DATABASE
var courses []Course

// func malware
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}
func main() {

}

//controllers- they go to a different file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Hi! My name is Ian and I'm learning about APIs"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	//grab id from request	
	params:=mux.Vars(r)
	//loop through courses, find matching id and retur the response
	for _, course:=range courses{
		if course.CourseId==params["id"]{
json.NewEncoder(w).Encode(course)
return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}