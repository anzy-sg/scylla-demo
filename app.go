package main

import (
	"fmt"

	"github.com/gocql/gocql"
)

func main() {
	// Create gocql cluster.
	cluster := gocql.NewCluster("mdw1_1", "mdw1_2")
	cluster.Keyspace = "mailsettings"
	cluster.Consistency = gocql.LocalQuorum

	session, err := gocql.NewSession(*cluster)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	query := `SELECT user_id, require_tls, require_valid_cert, version 
              FROM user_tls_policy
              WHERE user_id = ?`

	q := session.Query(query, 2)

	var userId int
	var requireTls bool
	var requireValidCert bool
	var version float32

	it := q.Iter()
	for it.Scan(&userId, &requireTls, &requireValidCert, &version) {
		fmt.Println(userId, requireTls, requireValidCert, version)
	}
}
