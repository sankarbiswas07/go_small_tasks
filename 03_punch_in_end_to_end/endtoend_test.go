package main

import (
	"testing"
	"time"
	"fmt"
)

// --------------------------------------
// End to End TEST
// --------------------------------------

// userId Unix time
type TableTest struct {
	userId      int
	loggedInFor int
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
		fmt.Println(tt)
		time.Sleep(time.Duration(tt.loggedInFor) * time.Second)
		u.punchIn(s)
	}
}
