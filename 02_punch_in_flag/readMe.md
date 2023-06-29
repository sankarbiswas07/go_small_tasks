  

# Table of Contents

  

- [Problem Statement](#problem-statement)
- [Run The Code Base](#run-the-code-base)
- [Types and Function](#type-and-function)
- [Test Case](#test-case)

  
  

## Problem Statement

  

Write a codebase where one user can login by punch-in. But on second time punch-in from same user will logged-out and show the logged-in time.

No database/ file writing should not be used.

  

## Run The Code Base

  

- Go installed in the machine. My local version is `go1.20.5 linux/amd64`

- Go to the project folder by issuing the command `cd 01_punch_in`

- Issue the command to run `go run main.go`

- For running test case, issue the command `go test`

  

## Types and Functions

***Types***

- `type session map[int]int64` => For storing userId as a key and UnixTime as value.

- `type User struct` => The struct has id for storing userID as integer
<br />

***Functions***

- `func main()` => Main Function, for running and Prompting user to punch-in.

<br />

- `func (s session) prompt()` => Receiver function of type session for prompting and asking for input from user.

- `func (s session) list()` => Receiver function to list the type session.

<br />

- `func (u User) punchIn(s session)` => Receiver function of type User for punch-in which will conditionally (Create if userID not exists in the session map else delete the user form session) call another receiver function from the following two:

- `func (u User) create(s session)` => Create a user in session (Entry of new userID in the map)

- `func (u User) delete(s session)` => Delete a user from session (Delete key from the map)

## Test Case

- `func TestUserPunchIn(t *testing.T)` :
	-  Case 1: User is created and logged in or not
	- Case 2: User is Removed or not while punching on 2nd time

- `func TestUserDelete(t *testing.T)` :
	- Case 1: UserID delete form season map or not 