package database

import (
	"fmt"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	fmt.Println("Connecting to DB..")
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "connectwise"
	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		fmt.Printf("Error in connecting to db: %v\n", err)
	} else {
		fmt.Println("Connection succeded !")
	}
}
