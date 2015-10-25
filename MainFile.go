package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "fmt"
        )




type GoogleMapAPI struct {
	Results []struct {
			FormattedAddress string `json:"formatted_address"`
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			
		} `json:"geometry"`
		
	} `json:"results"`
	Status string `json:"status"`
}

    
    type (  
    UserLocation struct {
        Id     bson.ObjectId `json:"id" bson:"_id"`
        Name   string        `json:"name" bson:"name"`
        Address string       `json:"address" bson:"address"`
        City    string       `json:"city" bson:"city"`
        State    string      `json:"state" bson:"state"`
        Zip string		 `json:"zip" bson:"zip"`
        Coordinate struct {   
        	Latitude float64  `json:"lat" bson:"lat"`
        	Longitude float64 `json:"lng" bson:"lng"`
        }

    }    
)


type UserController struct {  
    session *mgo.Session
}

var users []UserLocation



//mongo db session getsession details  ..local mongo 
/*func getSession() *mgo.Session {  
    // Connect to our local mongo
    s, err := mgo.Dial("localhost:27017")

    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
    }
    return s
}*/

func NewUserController(s *mgo.Session) *UserController {  
    return &UserController{s}
}

//Below function --connection with Remote mongolab
func RemoteMongoSession() *mgo.Session {
	mongolab_uri := "mongodb://mongoDeepa:welcome1@ds048878.mongolab.com:48878/mongodatabase"
	session, err := mgo.Dial(mongolab_uri)
	//Check if connection error
  	if err != nil {
    	fmt.Printf("Connection problem with mongo, go error %v\n", err)
  	}	
	return session
}


func getContent(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	// Read the content into a byte array
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// At this point we're done - simply return the bytes
	return body, nil
}


func GetContentDetails(ip string) /*{//*/(*GoogleMapAPI, error) {

	content, err := getContent(ip)
	if err != nil {
		// An error occurred while fetching the JSON
		//return nil, err
	}
	// Fill the record with the data from the JSON
	var record GoogleMapAPI
	err = json.Unmarshal(content, &record)
	if err != nil {
		// An error occurred while converting our JSON to an object
	//	return nil, err
	}
	
	return &record, err
	}



func main() {
    router := httprouter.New()
// Get a UserController instance
uc := NewUserController(RemoteMongoSession())    
 
    router.POST("/locations", uc.create) 
        router.GET("/locations/:location_id", uc.GetUser)
            router.PUT("/locations/:location_id", uc.PutUser)
    
              router.DELETE("/locations/:location_id", uc.DeleteUser)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: router,
    
}
    server.ListenAndServe()
}




