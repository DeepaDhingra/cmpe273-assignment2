package main

import (
 "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "strings"
		)
func (uc UserController) create(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {

 u := UserLocation{
        }
        
       // Populate the user data
    json.NewDecoder(req.Body).Decode(&u)

    // Add an Id
    u.Id = bson.NewObjectId()

	addr := u.Address
 	stateVar:= u.State
 	cityVar:= u.City
   
   addr = strings.Replace(addr, " ", "+", 6)
   stateVar = strings.Replace(stateVar, " ", "+", 6)
   cityVar = strings.Replace(cityVar, " ", "+", 6)
   fmt.Println("Vakue in addres is",u.Address)
   
    s := []string{}
	s = append(s,"http://maps.google.com/maps/api/geocode/json?address=")
	s = append(s,addr)
	s = append(s,",+")	
	s = append(s,cityVar)	
	s = append(s,",+")
	s = append(s,stateVar)
	s = append(s,"&sensor=false")
	var url string = s[0] + s[1] +s[2] +s[3]+s[4]+s[5]+s[6]
	
	fmt.Println("url is",url)
    record,_:= GetContentDetails(url)
	
	fmt.Println("record is", record)
	fmt.Println("Latitude is", record.Results[0].Geometry.Location.Lat)
	fmt.Println("Longitude is", record.Results[0].Geometry.Location.Lng)
		
	u.Coordinate.Latitude=record.Results[0].Geometry.Location.Lat
    u.Coordinate.Longitude=record.Results[0].Geometry.Location.Lng
        
            //Write the user to mongolab
    uc.session.DB("mongodatabase").C("CMPE273").Insert(u)
          
    // Marshal provided interface into JSON structure
    uj, _ := json.Marshal(u)
        
    fmt.Println("Users are ",uj)
   	rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(201)
    fmt.Fprintf(rw, "%s", uj)
      }

