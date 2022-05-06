package main

import (
	"fmt"

	"github.com/gocql/gocql"
)

func main() {
	// Create gocql cluster.
	cluster := gocql.NewCluster("mdw1_1")
	cluster.Keyspace = "moogle"
	cluster.Consistency = gocql.LocalQuorum
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy("mdw1")
	session, err := gocql.NewSession(*cluster)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	query(session)
}
func insert(session *gocql.Session) {
	insertStr := `INSERT INTO user (id, name, age, state) VALUES (?, ?, ?, ?)`

	err := session.Query(insertStr, 12, "Bob", 30, "AZ").Exec()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Inserted")
}

func query(session *gocql.Session) {
	queryStr := `SELECT id, name, age, state
              FROM user
              WHERE id = ?`

	q := session.Query(queryStr, 2)

	var id int
	var name string
	var age int
	var state string

	it := q.Iter()
	for it.Scan(&id, &name, &age, &state) {
		fmt.Println(id, name, age, state)
	}
}
