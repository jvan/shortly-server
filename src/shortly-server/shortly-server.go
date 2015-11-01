package main

import (
  "net/http"
  "github.com/labstack/echo"
  mw "github.com/labstack/echo/middleware"

  "database/sql"
  _ "github.com/mattn/go-sqlite3"

  "github.com/rs/cors"

  "strconv"

  "shortly"
)

var db *sql.DB  // database handle


//----- Data Structures -----------------------------------

type User struct {
  Id int        `json:"id"`
  Name string   `json:"name"`
}

type UserArray []User

//----- Routes --------------------------------------------

func ping(c *echo.Context) error {
  return c.String(http.StatusOK, "Ok")
}

func encode_value(c *echo.Context) error {
  // Convert the `val` parameter to an integer.
  value := c.Param("val")
  i, _ := strconv.Atoi(value)

  // Return the base-62 encoded value.
  return c.String(http.StatusOK, shortly.Encode(i))
}

func decode_value(c *echo.Context) error {
  // Decode the base-62 `val` parameter.
  value := c.Param("val")
  i := shortly.Decode(value)

  // Convert the base-10 integer to a string and return.
  s := strconv.Itoa(i)
  return c.String(http.StatusOK, s)
}

func get_all_users(c *echo.Context) error {
  // Query the database for all user records.
  rows, _ := db.Query("SELECT id, username FROM users")

  users := UserArray{}

  // For each row in the result create a new user object and add it to the
  // users array.
  for rows.Next() {
    var user_id int
    var username string
    rows.Scan(&user_id, &username)

    new_user := User{Id: user_id, Name: username}
    users = append(users, new_user)
  }

  // For ember-data compatibility we need to construct a response that 
  // contains a `users` object.
  type UserResponse struct {
    Users UserArray `json:"users"`
  }

  response := UserResponse{Users: users}

  return c.JSON(http.StatusOK, response)
}

func main() {
  e := echo.New()

  e.Use(mw.Logger())
  e.Use(mw.Recover())

  // Enable CORS.
  e.Use(cors.Default().Handler)

  // Initialize the database connection.
  db, _ = sql.Open("sqlite3", "test.db")

  e.Get("/ping", ping)

  e.Get("/encode/:val", encode_value)
  e.Get("/decode/:val", decode_value)

  e.Get("/users", get_all_users)

  e.Run(":1323")
}
