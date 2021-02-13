package cassandralib

import (
	"strings"

	"github.com/gocql/gocql"
)

func New(param CassandraLibParam) (*CassandraLib, error) {
	clusterIPs := strings.Split(param.ClusterIP, ",")
	cluster := gocql.NewCluster(clusterIPs...)
	cluster.Keyspace = param.KeySpace

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return &CassandraLib{
		session: session,
	}, nil
}

func (c *CassandraLib) QueryExec(query string, args ...interface{}) error {
	return c.session.Query(query, args).Exec()
}

func (c *CassandraLib) QueryIter(query string, args ...interface{}) *gocql.Iter {
	return c.session.Query(query, args...).Iter()
}

func (c *CassandraLib) Close() {
	c.session.Close()
}
