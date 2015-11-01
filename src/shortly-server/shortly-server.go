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
	Links []int   `json:"links"`
}

type Link struct {
  Id int        `json:"id"`
  Url string    `json:"url"`
  Key string    `json:"key"`
}


type UserArray []User


//----- Routes --------------------------------------------

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

func get_user(c *echo.Context) error {
	// Query the database for the specified user.
	id := c.Param("user_id")
	row, _ := db.Query("SELECT id, username FROM users WHERE id=$1 LIMIT 1", id)

	// TODO: Return NOT FOUND if no rows are returned.

	// Get first (and only) record.
	row.Next()

	var user_id int
	var username string

	row.Scan(&user_id, &username)

	// Query the database for all links associated with the user.
	rows, _ := db.Query("SELECT id FROM links WHERE user_id=$1", user_id)

	links := []int{}

	for rows.Next() {
		var link_id int
		rows.Scan(&link_id)
		links = append(links, link_id)
	}

	// Initialize a user object.
	user := User{Id: user_id, Name: username, Links: links}

	// For ember-data compatibility we need to construct a response that
	// contains a `user` object.
	type UserResponse struct {
		User User `json:"user"`
	}

	response := UserResponse{User: user}
	return c.JSON(http.StatusOK, response)
}

func get_link(c *echo.Context) error {
	// Query the database for the specified user.
	id := c.Param("link_id")
	row, _ := db.Query("SELECT id, url FROM links WHERE id=$1 LIMIT 1", id)

	// TODO: Return NOT FOUND if no rows are returned.

	// Get the first (and only) record
	row.Next()

	var uid int
	var url string

	row.Scan(&uid, &url)

	// Initialize a link object. Use the `Encode` function to base-62
	// encode the link id.
	link := Link{Id: uid, Url: url, Key: shortly.Encode(uid)}

	// For ember-data compatibility we need to construct a response that
	// contains a `user` object.
	type LinkResponse struct {
		Link Link `json:"link"`
	}

	defer row.Close()

	response := LinkResponse{Link: link}
	return c.JSON(http.StatusOK, response)
}


func create_link(c *echo.Context) error {

  type LinkParams struct {
    Url string    `json:"url"`
    UserId string `json:"user_id"`
  }

  params := &LinkParams{}

  if err := c.Bind(params); err != nil {
    return err
  }

  result, _ := db.Exec("INSERT INTO links (url, user_id) VALUES ($1, $2)", params.Url, params.UserId)

  id, _ := result.LastInsertId()

  rows, _ := db.Query("SELECT id, url FROM links WHERE id=$1 LIMIT 1", id)

  rows.Next()

  var uid int
  var url string
  rows.Scan(&uid, &url)

  link := Link{Id: uid, Url: url, Key: shortly.Encode(uid)}

  type LinkResponse struct {
    Link Link `json:"link"`
  }

  defer rows.Close()

  response := LinkResponse{Link: link}
  return c.JSON(http.StatusCreated, response)
}


func validate_user(c *echo.Context) error {

  type UserParams struct {
    Username string `json:"username"`
    Password string `json:"password"`
  }

  params := &UserParams{}

  if err := c.Bind(params); err != nil {
    return err
  }

	// Query the database for the specified username.
	row, _ := db.Query("SELECT id, username FROM users WHERE username=$1 LIMIT 1", params.Username)

	// TODO: Return NOT FOUND if no rows are returned.

	// Get first (and only) record.
	row.Next()

	var user_id int
	var username string

	row.Scan(&user_id, &username)

	// Query the database for all links associated with the user.
	rows, _ := db.Query("SELECT id FROM links WHERE user_id=$1", user_id)

	links := []int{}

	for rows.Next() {
		var link_id int
		rows.Scan(&link_id)
		links = append(links, link_id)
	}

	// Initialize a user object.
	user := User{Id: user_id, Name: username, Links: links}

	// For ember-data compatibility we need to construct a response that
	// contains a `user` object.
	type UserResponse struct {
		User User `json:"user"`
	}

	response := UserResponse{User: user}
	return c.JSON(http.StatusOK, response)

}

func redirect(c *echo.Context) error {
	// Get the base-62 key and decode it to get the record id.
	key := c.Param("key")
	id := shortly.Decode(key)

	row, _ := db.Query("SELECT url FROM links WHERE id=$1 LIMIT 1", strconv.Itoa(id))

	row.Next()

	var url string
	row.Scan(&url)

	c.Redirect(301, url)

	return nil
}


func main() {
  e := echo.New()

  e.Use(mw.Logger())
  e.Use(mw.Recover())

  // Enable CORS.
  e.Use(cors.Default().Handler)

  // Initialize the database connection.
  db, _ = sql.Open("sqlite3", "test.db")

  e.Get("/users", get_all_users)
	e.Get("/users/:user_id", get_user)

	e.Get("/links/:link_id", get_link)
	e.Post("/links", create_link)

	e.Get("/go/:key", redirect)

	e.Post("/login", validate_user)

  e.Run(":1323")
}
