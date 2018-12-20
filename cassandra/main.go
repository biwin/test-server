package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

var Session *gocql.Session

func init() {
	var err error

	cluster := gocql.NewCluster("167.99.230.134")
	cluster.ConnectTimeout = time.Second * 10
	cluster.DisableInitialHostLookup = true

	cluster.Keyspace = "biwins"
	Session, err = cluster.CreateSession()
	if err != nil {
		//panic(err)
	}
	fmt.Println("cassandra init done")
}
