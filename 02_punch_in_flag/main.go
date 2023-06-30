package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var isForTest bool

// userId Unix time
type Session map[int]int64

type User struct {
	id int
}

type TestSuit struct {
	testSessionMaxCreate int
	cycleNo              float32
	user                 *User
}

func printInstruction() {
	fmt.Println("------------------------------------")
	fmt.Println("Who is logging in( > 1) ?")
	fmt.Println("Press -1 : Logged in userId list")
	fmt.Println("Press 0 : Terminate program")
	fmt.Println("------------------------------------")
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

func (tc *TestSuit) runEndToEnd() {

	time.Sleep(4 * time.Duration(time.Second))
	user := User{id: tc.testSessionMaxCreate / 2}
	tc.user = &user

	if (tc.testSessionMaxCreate - 2) == 0 {
		tc.testSessionMaxCreate = 6
		tc.cycleNo -= .5
	} else {
		tc.testSessionMaxCreate -= 2
	}
}

func (s Session) prompt(tc *TestSuit) error {

	var u User

	if isForTest && tc.cycleNo != 0 {
		// Run End to End test case
		tc.runEndToEnd()
		u = *tc.user // assign the point of the same user Struct
	} else if isForTest && tc.cycleNo == 0 {
		os.Exit(1) //c losing end to end test case
	} else {
		printInstruction()
		// read from user input
		_, err := fmt.Scanf("%d", &u.id)

		if err != nil {
			fmt.Println("Error reading input:", err)
			return nil
		}
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
	s.prompt(tc)

	return nil
}

func init() {
	flag.BoolVar(&isForTest, "isForTest", false, "isForTest[Bool] flag(optional) for testing End to End TestCase")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
		flag.PrintDefaults()
	}

	// fmt.Println("reflect.TypeOf(isForTest) - ", reflect.TypeOf(isForTest))
	fmt.Println("Value of isForTest -", isForTest)
}

func main() {
	flag.Parse()
	var tc TestSuit

	// Check for errors during flag parsing
	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Error: Unknown arguments\n")
		flag.Usage()
		os.Exit(1)
	}

	// If flag is enable make a test suit
	if isForTest {
		tc = TestSuit{
			testSessionMaxCreate: 6,
			cycleNo:              1,
		}
	}

	s := make(Session)
	if err := s.prompt(&tc); err != nil {
		fmt.Printf("s.prompt(): %v", err)
	}
}
