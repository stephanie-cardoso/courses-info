package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	var users []UserData
	var courses []PublishedCourse

	allCourses := getCourses()
	publishedCourses := getPublishedCourses(allCourses)

	for _, course := range publishedCourses {
		enrolledUsers := getEnrolledUsersIDByCourse(course.ID)
		for _, user := range enrolledUsers.Users {
			users = append(users, getUserData(user.ID))
		}
		course := PublishedCourse{Name: course.Name, Heading: course.Heading, Users: users}
		courses = append(courses, course)
	}

	coursesJSON, _ := json.Marshal(courses)
	fmt.Println(string(coursesJSON))
}

func getCourses() Courses {
	url := "https://developers.teachable.com/v1/courses"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("apiKey", "7JbSA3ep6XOMV3t8t7QXuXq9HS79Dwnr")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

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
	url := fmt.Sprintf("https://developers.teachable.com/v1/courses/%d/enrollments", courseID)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("apiKey", "7JbSA3ep6XOMV3t8t7QXuXq9HS79Dwnr")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var users Users
	json.Unmarshal(body, &users)

	return users
}

func getUserData(userID int) UserData {
	url := fmt.Sprintf("https://developers.teachable.com/v1/users/%d", userID)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("apiKey", "7JbSA3ep6XOMV3t8t7QXuXq9HS79Dwnr")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var userData UserData
	json.Unmarshal(body, &userData)

	return userData
}
