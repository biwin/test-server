package users

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"goram/cassandra"
	"net/http"
)

// Post -- handles POST request to /users/new to create new user
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params
func Post(w http.ResponseWriter, r *http.Request) {
	var errs []string
	var gocqlUUID gocql.UUID

	user, errs := FormToUser(r)

	var created bool = false
	if len(errs) == 0 {
		fmt.Println("creating a new user")
		gocqlUUID = gocql.TimeUUID()
		if err := cassandra.Session.Query(`
		INSERT INTO users (id, firstname, lastname, email, city, age) VALUES (?, ?, ?, ?, ?, ?)`,
			gocqlUUID, user.FirstName, user.LastName, user.Email, user.City, user.Age).Exec(); err != nil {
			errs = append(errs, err.Error())
		} else {
			created = true
		}
	}

	if created {
		fmt.Println("user_id", gocqlUUID)
		json.NewEncoder(w).Encode(NewUserResponse{ID: gocqlUUID})
	} else {
		fmt.Println("errors", errs)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
