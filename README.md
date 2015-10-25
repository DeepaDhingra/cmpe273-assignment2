# cmpe273-assignment2
Description : This is a location finder service in Go.
Technical Description: REST service, DataBase: Mongodb (CRUD Operations)
CRUD Location Service [Requirement]

The location service has following REST endpoints to store and retrieve locations. All the data must be persisted into MongoDB. 
To lookup coordinates of a location,i have used  use Google Map Api.
This project use database on free Mongolab. ALL CRUD operations are defined using Mongolab.
For go and Mongodb integration. I have used mgo.

Testing and Execution Steps:
For testing this project: use any REST client. I have used POSTER from  Mozilla Firefox.
Below are the steps for performing CRUD operations:

1) Create New Location - POST 
URL: http://localhost:8080/locations
In Content to send tab: Content Type : application/json
In Headers tab: Name: Content-type Value: application/json
click on Add/Change button to add the above said properties.
