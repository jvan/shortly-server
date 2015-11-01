#!/bin/bash
#
# Create and initialize a sqlite database for short.ly service.
#

SCRIPT_DIR=$(dirname $(readlink -f $0))

# Install the database in the same directory as the server executable.
DATABASE="${SCRIPT_DIR}/../go/bin/test.db"

if [[ -e "${DATABASE}" ]]; then
  echo -n "Database file already exists. Delete (Y/n)? "
  read response
  if [[ "${response}" == "n" ]] || [[ "${response}" == "N" ]]; then
    echo "Exiting."
  else
    echo "Deleting database."
    rm $DATABASE
  fi
fi

# Create a table for users and add some test data.
sqlite3 $DATABASE 'CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, username VARCHAR(64) NOT NULL)'

# Create a user for unregistered links.
sqlite3 $DATABASE 'INSERT INTO users (id, username) VALUES (0, "unregistered")'

sqlite3 $DATABASE 'INSERT INTO users (username) VALUES ("alice@gmail.com")'
sqlite3 $DATABASE 'INSERT INTO users (username) VALUES ("bob@gmail.com")'
sqlite3 $DATABASE 'INSERT INTO users (username) VALUES ("carol@gmail.com")'


# Create a table for links and add some test data.
sqlite3 $DATABASE 'CREATE TABLE links (id INTEGER PRIMARY KEY AUTOINCREMENT, url STRING, last_click DATETIME, user_id INT, FOREIGN KEY (user_id) REFERENCES users(id))'

# Unregistered links
sqlite3 $DATABASE 'INSERT INTO links (url, user_id) VALUES ("http://foobar.com", 0)'
sqlite3 $DATABASE 'INSERT INTO links (url, user_id) VALUES ("http://example.com", 0)'


# Alice likes search engines
sqlite3 $DATABASE 'INSERT INTO links (url, user_id) VALUES ("http://duckduckgo.com", 1)'
sqlite3 $DATABASE 'INSERT INTO links (url, user_id) VALUES ("http://google.com", 1)'
sqlite3 $DATABASE 'INSERT INTO links (url, user_id) VALUES ("http://yahoo.com", 1)'

# Bob works on the backend
sqlite3 $DATABASE 'INSERT INTO links (url, user_id) VALUES ("http://golang.org", 2)'
sqlite3 $DATABASE 'INSERT INTO links (url, user_id) VALUES ("http://labstack.com/echo/", 2)'

# Carol works on the frontend
sqlite3 $DATABASE 'INSERT INTO links (url, user_id) VALUES ("http://emberjs.com", 3)'
sqlite3 $DATABASE 'INSERT INTO links (url, user_id) VALUES ("http://ember-cli.com", 3)'
sqlite3 $DATABASE 'INSERT INTO links (url, user_id) VALUES ("http://getbootstrap.com", 3)'
