package main

import (
	"fmt"

	"github.com/gocql/gocql"
)

func main() {
	// Create gocql cluster.
	cluster := gocql.NewCluster("mdw1_1")
	cluster.Keyspace = "mailsettings"
	cluster.Consistency = gocql.LocalQuorum
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.DCAwareRoundRobinPolicy("las1"))
	session, err := gocql.NewSession(*cluster)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	insert(session)
}
func insert(session *gocql.Session) {
	insertStr := `INSERT INTO user_tls_policy (user_id, require_tls, require_valid_cert, version) VALUES (?, ?, ?, ?)`

	err := session.Query(insertStr, 5, true, true, 1.5).Exec()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Inserted")
}

func query(session *gocql.Session) {
	queryStr := `SELECT user_id, require_tls, require_valid_cert, version 
              FROM user_tls_policy
              WHERE user_id = ?`

	q := session.Query(queryStr, 2)

	var userId int
	var requireTls bool
	var requireValidCert bool
	var version float32

	it := q.Iter()
	for it.Scan(&userId, &requireTls, &requireValidCert, &version) {
		fmt.Println(userId, requireTls, requireValidCert, version)
	}
}
