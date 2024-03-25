#!/bin/bash

echo "Create and Set up database"

./sqlite/sqlite3.exe ./paints.db <<-EOS

# drop users table if exiting
drop table if exists users;

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    role TEXT CHECK (role IN ('admin', 'manager', 'painter')) NOT NULL,
    permission TEXT NOT NULL,
    password TEXT NOT NULL
);

# create the admin user
insert into users values (1, 'admin', 'admin', 'view and edit', 'admin');

# drop paints table if exiting
drop table if exists paints;

CREATE TABLE IF NOT EXISTS paints (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  color TEXT CHECK (color IN ('blue', 'grey', 'black', 'white', 'purple')) NOT NULL UNIQUE,
  quantity INTEGER NOT NULL
);

# populate the paints table with the colors
insert into paints values (1, 'blue', 10);
insert into paints values (2, 'grey', 100);
insert into paints values (3, 'black', 25);
insert into paints values (4, 'white', 0);
insert into paints values (5, 'purple', 75);

EOS