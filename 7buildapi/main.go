package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice string `json:"price"`
	Author *Author `json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// fake DB
var courses []Course

// middleware
func (c *Course) IsEmpty() bool{
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}


func main(){
	// getOneCourse()
}

// controllers
func serveHome(w http.ResponseWriter , r *http.Request){
	w.Write([]byte("<h1>Welcome to the Landing Page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){

	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	
	// loop through courses and find matching id and return it
	for _,course := range courses {
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request){

	fmt.Println("Creating a new course : ")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course 
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique id and then append to courses
	var rend = rand.New(rand.NewSource(time.Now().UnixNano()))
	
	course.CourseId = strconv.Itoa(rend.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return



}