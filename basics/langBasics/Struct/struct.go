package main

import "fmt"

type user struct {
	name   string
	rollNo string
	subject
}

type subject struct {
	maths   string
	eng     string
	physics string
}

func structy() {
	user1 := user{
		name:   "sid",
		rollNo: "101",
		subject: subject{
			maths:   "maths101",
			eng:     "eng101",
			physics: "physics101",
		},
	}

	fmt.Printf("Name: %s, Roll: %s, Maths: %s, Eng: %s, Physics: %s\n",
		user1.name, user1.rollNo, user1.maths, user1.subject.maths, user1.physics)

}

func main() {
	structy()
}
