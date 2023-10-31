package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// DB
var courses []Course

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

// helper
func (c *Course) IsEmpty() bool {
	return c.CourseId == ""
}

func main() {

	courses = append(courses, Course{"1", "Math", 100, &Author{"Author1", "web1"}})
	courses = append(courses, Course{"2", "Physics", 100, &Author{"Author2", "web2"}})
	r := mux.NewRouter()
	r.HandleFunc("/course", listAllCourses).Methods("GET")
	r.HandleFunc("/course", addCourse).Methods("POST")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	r.HandleFunc("/posts", getPosts).Methods("GET")
	r.HandleFunc("/posts", createPost).Methods("POST")
	r.HandleFunc("/posts", createPost).Methods("POST")
	log.Fatal(http.ListenAndServe(":4000", r))

}

func listAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("listAllCourses")
	w.Write([]byte("All courses\n"))
	_ = json.NewEncoder(w).Encode(courses)
}

func addCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("AddCourse")

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	course.CourseId = strconv.Itoa(rand.Intn(1000000))
	courses = append(courses, course)
	_ = json.NewEncoder(w).Encode(&courses)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, post)
	json.NewEncoder(w).Encode(&post)
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Delete Course")

	for idx, course := range courses {
		if course.CourseId == mux.Vars(r)["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
		}
	}
	_ = json.NewEncoder(w).Encode(courses)
}
