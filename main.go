package main

import (
	"courses-info/client"
	"encoding/json"
	"fmt"
)

const (
	baseAPIURL = "https://developers.teachable.com/v1"
)

type Course struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Heading   string `json:"heading"`
	Published bool   `json:"is_published"`
}

type Courses struct {
	Courses []Course `json:"courses"`
}

type UserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserID struct {
	ID int `json:"user_id"`
	//Name  string
	//Email string
}

type Users struct {
	Users []UserID `json:"enrollments"`
}

type PublishedCourse struct {
	Name    string     `json:"course_name"`
	Heading string     `json:"course_heading"`
	Users   []UserData `json:"users"`
}

func main() {
	var courses []PublishedCourse

	allCourses := getCourses()
	publishedCourses := getPublishedCourses(allCourses)

	for _, course := range publishedCourses {
		enrolledUsers := getEnrolledUsersIDByCourse(course.ID)
		usersByCourse := getUsersByCourse(enrolledUsers)

		course := PublishedCourse{Name: course.Name, Heading: course.Heading, Users: usersByCourse}
		courses = append(courses, course)
	}

	coursesJSON, _ := json.Marshal(courses)
	fmt.Println(string(coursesJSON))
}

func getUsersByCourse(enrolledUsers Users) []UserData {
	var users []UserData
	for _, user := range enrolledUsers.Users {
		users = append(users, getUserData(user.ID))
	}
	return users
}

func getCourses() Courses {
	url := fmt.Sprintf("%s/courses", baseAPIURL)

	body, _ := client.DoRequest(url)

	var courses Courses
	json.Unmarshal(body, &courses)

	return courses
}

func getPublishedCourses(allCourses Courses) []Course {
	var publishedCourses []Course
	for _, course := range allCourses.Courses {
		if course.Published {
			publishedCourses = append(publishedCourses, course)
		}
	}
	return publishedCourses
}

func getEnrolledUsersIDByCourse(courseID int) Users {
	url := fmt.Sprintf("%s/courses/%d/enrollments", baseAPIURL, courseID)

	body, _ := client.DoRequest(url)
	//tratar erro

	var users Users
	json.Unmarshal(body, &users)

	return users
}

func getUserData(userID int) UserData {
	url := fmt.Sprintf("%s/users/%d", baseAPIURL, userID)

	body, _ := client.DoRequest(url)

	var userData UserData
	json.Unmarshal(body, &userData)

	return userData
}
