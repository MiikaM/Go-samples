# Description

This project is a **VERY** simple of a CRUD api with 5 endpoints: 
- "/movies"
    - POST, create a movie
    - GET, Get all movies
- "/movies/{id}
    - POST, update a movie
    - GET, get a movie
    - DELETE, delete a movie

The project has 2 movies mocked to the server and is not connected to any database.

***NOTE!*** Update a movie requires the full information of the movie to update, else the api removes all fields not included in the request. 

# Try it out

***NOTE!*** This assumes that you have Golang installed on your device.

You can try the project by cloning the project and moving to "webserver" directory and running *"go run ."* command in the directory.


