package Database

import (
	"gopkg.in/mgo.v2"
	"log"
)

const (
	DBNAME     = "your_teams"
	MONGO_IP   = "localhost"
	MONGO_PORT = ":27017"
	MONGO_HOST = MONGO_IP + MONGO_PORT
)

var database *mgo.Database

//-------------------------------- DB --------------------------------//
func StartConectionDatabase() (err error) {
	session, err := mgo.Dial(MONGO_HOST)
	if err != nil {
		log.Print("Failed to establish connection to MongoDB Server: ", err)
		return
	}
	database = session.DB(DBNAME)
	return
}

//-------------------------------- DB --------------------------------//
func OpenSession(docName string) *mgo.Collection {
	return database.C(docName)
}
