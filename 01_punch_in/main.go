package main

import (
	"fmt"
	"os"
	"time"
)

// userId Unix time
type Session map[int]int64

type User struct {
	id int
}

func (u User) create(s Session) {
	s[u.id] = time.Now().Unix()
}

func (u User) delete(s Session) {
	difference := time.Unix((time.Now().Unix() - s[u.id]), 0)
	fmt.Printf("UserID => %d Punched Out after %d seconds\n", u.id, difference.Second())
	delete(s, u.id)
}

func (u User) punchIn(s Session) error {

	if u == (User{}) {
		return fmt.Errorf("User can not be an empty object")
	}
	if s == nil {
		return fmt.Errorf("Session can not be nil")
	}

	// Logout if a key exists
	if _, ok := s[u.id]; ok {
		u.delete(s)
		return nil
	}
	u.create(s)
	fmt.Println()
	fmt.Printf("UserId => %d logged-in\n", u.id)
	fmt.Println()
	return nil
}

func (s Session) list() error {
	if s == nil {
		return fmt.Errorf("Session can not be nil")
	}
	for key, value := range s {
		fmt.Printf("UserId: %d, Login: %d\n", key, value)
	}
	return nil
}

func (s Session) prompt() error {

	var u User

	fmt.Println("\nWho is logging in( > 1) ?")
	fmt.Println("Press -1 : Logged in userId list")
	fmt.Println("Press 0 : Terminate program")

	_, err := fmt.Scanf("%d", &u.id)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return nil
	}

	switch u.id {
	case -1:
		s.list()
	case 0:
		os.Exit(0)
	default:
		if err := u.punchIn(s); err != nil {
			return fmt.Errorf("u.punchIs(S): %w", err)
		}
	}
	// Repeat Prompt
	s.prompt()
	return nil
}

func main() {
	s := make(Session)
	if err := s.prompt(); err != nil {
		fmt.Printf("s.prompt(): %v", err)
	}
}
