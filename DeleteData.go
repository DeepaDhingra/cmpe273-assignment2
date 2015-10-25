package main
import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "fmt"
    "gopkg.in/mgo.v2/bson"
)
func (uc UserController) DeleteUser(rw http.ResponseWriter, req *http.Request, p httprouter.Params) { 
fmt.Println("indise delete") 
    // Grab id
    id := p.ByName("location_id")
	fmt.Println("id is",id)

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        rw.WriteHeader(404)
        return
    }

    // Grab id
    oid := bson.ObjectIdHex(id)
	fmt.Println("oid is",oid)
    // Delete user
    uc.session.DB("mongodatabase").C("CMPE273").RemoveId(oid)//; err != nil{
	fmt.Println("no error")
    // Write status
    rw.WriteHeader(200)
}

