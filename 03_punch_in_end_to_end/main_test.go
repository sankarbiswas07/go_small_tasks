package main

import (
	"testing"
	"time"
)

// --------------------------------------
// UNIT TEST
// --------------------------------------
func TestUserPunchIn(t *testing.T) {
	s := make(Session)
	u := User{id: 1}

	// Verify that user is created and logged in
	u.punchIn(s)
	// Wait for 3 sec
	time.Sleep(3 * time.Second)
	if _, ok := s[u.id]; !ok {
		t.Errorf("UserID => %d was not created and logged in", u.id)
	}

	// Verify that user is logged out when punching in again
	u.punchIn(s)

	if _, ok := s[u.id]; ok {
		t.Errorf("UserId => %d was not logged out when punching in again", u.id)
	}
}

func TestUserDelete(t *testing.T) {
	s := make(Session)
	u := User{id: 1}

	// Log in the user
	u.punchIn(s)
	// Wait for 2 sec
	time.Sleep(2 * time.Second)

	// Verify that user is logged in
	if _, ok := s[u.id]; !ok {
		t.Errorf("UserID => %d is not there in map", u.id)
	}

	// Log out the user
	u.delete(s)

	// Verify that user is logged out
	if _, ok := s[u.id]; ok {
		t.Errorf("UserID => %d is nor deleted", u.id)
	}
}
