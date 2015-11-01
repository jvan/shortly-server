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

# Create a table for users and add some test data
sqlite3 $DATABASE 'CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, username VARCHAR(64) NOT NULL)'

sqlite3 $DATABASE 'INSERT INTO users (username) VALUES ("alice@gmail.com")'
sqlite3 $DATABASE 'INSERT INTO users (username) VALUES ("bob@gmail.com")'
sqlite3 $DATABASE 'INSERT INTO users (username) VALUES ("carol@gmail.com")'
