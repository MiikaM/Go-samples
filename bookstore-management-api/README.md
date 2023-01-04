# Description

This project is a **VERY** simple of CRUD api for Mysql with 5 endpoints: 
- "/book/"
    - POST, create a book
    - GET, Get all books
- "/book/{id}
    - POST, update a book
    - GET, get a book
    - DELETE, delete a book

# Try it out

***NOTE!*** This assumes that you have Golang installed on your device and a mysql server running.

You can try the project by cloning the project and moving to "bookstore-management-api" directory and navigating to "**cmd/main**" and and running *"go run ."* command in the directory.

Environment variables for connecting to your database:
- MYSQL_USER
    - username for the mysql server
    - e.g. "root
- MYSQL_USERPASSWORD
    - user password for the mysql server
    - e.g. "password" (hope not)
- MYSQL_CONNECTION
    - Connection string (address)
    - e.g. "(0.0.0.0.1)/9000"
- MYSQL_ENDPOINT
    - database to connect to
    - e.g. "books


