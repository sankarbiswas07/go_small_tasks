package main

import (
	"fmt"
	"os"
	"time"
)

// userId Unix time
type session map[int]int64

type User struct {
	id int
}

func (u User) create(s session) {
	s[u.id] = time.Now().Unix()
}

func (u User) delete(s session) {
	difference := time.Unix((time.Now().Unix() - s[u.id]), 0)
	fmt.Printf("UserID => %d Punched Out after %d seconds\n", u.id, difference.Second())
	delete(s, u.id)
}

func (u User) punchIn(s session) {
	// Logout if a key exists
	if _, ok := s[u.id]; ok {
		u.delete(s)
	} else {
		u.create(s)
		fmt.Println()
		fmt.Printf("UserId => %d logged-in\n", u.id)
		fmt.Println()
	}
}

func (s session) list() {
	for key, value := range s {
		fmt.Printf("UserId: %d, Login: %d\n", key, value)
	}
}

func (s session) prompt() {

	var u User

	fmt.Println("\nWho is logging in( > 1) ?")
	fmt.Println("Press -1 : Logged in userId list")
	fmt.Println("Press 0 : Terminate program")

	_, err := fmt.Scanf("%d", &u.id)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	switch u.id {
	case -1:
		s.list()
	case 0:
		os.Exit(0)
	default:
		u.punchIn(s)
	}
	// Repeat Prompt
	s.prompt()
}

func main() {
	s := make(session)
	s.prompt()
}
