package cassandralib

import "github.com/gocql/gocql"

type (
	CassandraLibParam struct {
		ClusterIP string
		KeySpace  string
	}

	CassandraLib struct {
		session *gocql.Session
	}
)
