package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"labix.org/v2/mgo"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		session, collection, err := mgoInit()
		if err != nil {
			return
		}
		defer session.Close()
		err = collection.Insert(time.Now())
		if err != nil {
			fmt.Fprintf(w, "Write returned error")
		}
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}
func mgoInit() (session *mgo.Session, collection *mgo.Collection, err error) {
	session, collection = nil, nil
	server := os.Getenv("DB_PORT")
	if server == "" {
		server = "localhost"
	}
	session, err = mgo.Dial(server)
	if err != nil {
		log.Fatalf("Error connecting:  %v", err)
		return
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "default"
	}
	c := os.Getenv("DB_COLLECTION")
	if c == "" {
		c = "hellworld"
	}
	collection = session.DB(dbName).C(c)
	return
}
