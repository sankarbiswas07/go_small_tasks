package main

import (
	"testing"
	"time"
)

// --------------------------------------
// End to End TEST
// --------------------------------------
type TableTest struct {
	userId      int
	loggedInFor int
}

func _secondPunch(s Session, u User, loggedInFor int) {
	// Wait for 3 sec
	time.Sleep(time.Duration(loggedInFor) * time.Second)
	u.punchIn(s)
}

func TestEndToEnd(t *testing.T) {
	tableTest := []TableTest{
		{
			userId:      1,
			loggedInFor: 5,
		},
		{
			userId:      2,
			loggedInFor: 4,
		},
		{
			userId:      3,
			loggedInFor: 3,
		},
		{
			userId:      4,
			loggedInFor: 2,
		},
		{
			userId:      5,
			loggedInFor: 1,
		},
	}
	s := make(Session)

	for _, tt := range tableTest {
		u := User{id: tt.userId}
		// Initiate first punch
		u.punchIn(s)
		_secondPunch(s, u, tt.loggedInFor)
	}

}
