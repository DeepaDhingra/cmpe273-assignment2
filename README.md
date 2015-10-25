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

Run the server through command prompt: go run MainFile.go DeleteData.go GetData.go PostData.go PutData.go

Configuration steps in POSTER:

URL: http://localhost:8080/locations
In Content to send tab: Content Type : application/json
In Headers tab: Name: Content-type Value: application/json
click on Add/Change button to add the above said properties.
Following is the example of one of the test cases:


1) Create New Location - POST 


         URL: http://localhost:8080/locations
         
Request:

{
   "name" : "John Smith",
   "address" : "123 Main St",
   "city" : "San Francisco",
   "state" : "CA",
   "zip" : "94113"
}


Response: 
{"id":"562c68d5a02cc51e1457cb66","name":"John Smith","address":"123 Main St","city":"San Francisco","state":"CA","zip":"94113","Coordinate":{"lat":37.7917618,"lng":-122.3943405}}

Please observe the Response Code as 201

HTTP Response Code: 201

2) Get a Location - GET/locations/{location_id}
Request: URL

                  URL: http://localhost:8080/locations/562c68d5a02cc51e1457cb66 (Id created during POST operation)


Response: 
{"id":"562c68d5a02cc51e1457cb66","name":"John Smith","address":"123 Main St","city":"San Francisco","state":"CA","zip":"94113","Coordinate":{"lat":37.7917618,"lng":-122.3943405}}
Please observe the Response Code as 200
HTTP Response Code: 200

3) Update a Location - PUT /locations/{location_id}

                  URL: http://localhost:8080/locations/562c68d5a02cc51e1457cb66 (Id created during POST operation)

Request: {
   "address" : "1600 Amphitheatre Parkway",
   "city" : "Mountain View",
   "state" : "CA",
   "zip" : "94043"
}

Response: 
{"id":"562c68d5a02cc51e1457cb66","name":"John Smith","address":"1600 Amphitheatre Parkway","city":"Mountain View","state":"CA","zip":"94043","Coordinate":{"lat":37.4220352,"lng":-122.0841244}}

Please observe the Response Code as 201
HTTP Response Code: 201

4) Delete a Location - DELETE /locations/{location_id}
Request: URL

                  URL: http://localhost:8080/locations/562c68d5a02cc51e1457cb66 (Id created during POST operation)
Response: HTTP Response Code: 200
