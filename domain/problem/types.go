package problem

import "github.com/beruangcoklat/code-execution-engine/pkg/lib/cassandralib"

type (
	Problem struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}

	ResourceItf interface {
		GetProblemsCass(limit, offset int) ([]Problem, error)
	}

	DomainItf interface {
		GetProblems(limit, offset int) ([]Problem, error)
	}

	Resource struct {
		Cassandra *cassandralib.CassandraLib
	}

	Domain struct {
		Resource ResourceItf
	}
)
