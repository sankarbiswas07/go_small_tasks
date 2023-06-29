package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var isForTest bool

func init() {
	flag.BoolVar(&isForTest, "isForTest", false, "isForTest flag(optional) for testing End to End TestCase")
}

// userId Unix time
type Session map[int]int64

type User struct {
	id int
}

func (u User) create(s Session) error {
	s[u.id] = time.Now().Unix()
	_, exists := s[u.id]
	if !exists {
		return fmt.Errorf("time.Now().Unix(): return error")
	}
	return nil
}

func (u User) delete(s Session) error {
	difference := time.Unix((time.Now().Unix() - s[u.id]), 0)
	delete(s, u.id)
	// Check if the key exists
	if _, exists := s[u.id]; exists {
		return fmt.Errorf("delete(s, u.id): is not working, User is still exists in the session")
	}
	fmt.Printf("UserID => %d Punched Out after %d seconds\n", u.id, difference.Second())
	return nil
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
		// Handle error in delete func if occur
		if err := u.delete(s); err != nil {
			return fmt.Errorf("u.delete(s): %w", err)
		}
		return nil
	}

	if err := u.create(s); err != nil {
		return fmt.Errorf("u.create(s): %s", err)
	}

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
	flag.Parse()
	// Check for errors during flag parsing
	// if err := flag.ErrHelp; err != nil {
	// 	fmt.Print(err)
	// 	fmt.Fprintf(os.Stderr, "flag.Parse(): Error parsing flags: %v\n", err)
	// 	os.Exit(1)
	// }

	fmt.Println("flag", isForTest)

	s := make(Session)
	if err := s.prompt(); err != nil {
		fmt.Printf("s.prompt(): %v", err)
	}
}
