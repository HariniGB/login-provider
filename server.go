package main

// IMPORTANT LINK FOR REFERENCE:  https://stevenwhite.com/tag/golang/

import (
  // Standard library packages
  "net/http"

  // Third party packages
  "github.com/julienschmidt/httprouter"
  "github.com/HariniGB/go-rest-api/controllers"
  "gopkg.in/mgo.v2"
)

func main() {
  // Instantiate a new router
  r := httprouter.New()

  // Get a UserController instance
  uc := controllers.NewUserController(getSession())

  // Get all users resources
  r.GET("/users", uc.GetUsers)

  // Get a user resource
  r.GET("/user/:id", uc.GetUser)

  // Create a new user
  r.POST("/user", uc.CreateUser)

  //  Update a user
  r.PUT("/user/:id", uc.UpdateUser)

  // Remove an existing user
  r.DELETE("/user/:id", uc.RemoveUser)

  // Fire up the server
  http.ListenAndServe("localhost:3000", r)
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
  // Connect to our local mongo
  s, err := mgo.Dial("mongodb://localhost")

  // Check if connection error, is mongo running?
  if err != nil {
    panic(err)
  }

  // Deliver session
  return s
}
