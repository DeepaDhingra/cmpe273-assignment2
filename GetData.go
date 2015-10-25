package main
import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
        )
func (uc UserController) GetUser(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {  
    // Grab id
    id := p.ByName("location_id")

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        rw.WriteHeader(404)
        return
    }

    // Grab id
    oid := bson.ObjectIdHex(id)

    // Stub user
	u := UserLocation{
        }
    // Fetch user
    if err := uc.session.DB("mongodatabase").C("CMPE273").FindId(oid).One(&u); err != nil {
        rw.WriteHeader(404)
        return
    }

    // Marshal provided interface into JSON structure
    uj, _ := json.Marshal(u)

    // Write content-type, statuscode, payload
    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(200)
    fmt.Fprintf(rw, "%s", uj)
}
