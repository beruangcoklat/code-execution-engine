package server

import (
	"os"

	problemdomain "github.com/beruangcoklat/code-execution-engine/domain/problem"
	"github.com/beruangcoklat/code-execution-engine/pkg/lib/cassandralib"
	"github.com/beruangcoklat/code-execution-engine/pkg/lib/configlib"
	"github.com/beruangcoklat/code-execution-engine/pkg/lib/errorlib"
	"github.com/beruangcoklat/code-execution-engine/pkg/lib/loglib"
	executecodeuc "github.com/beruangcoklat/code-execution-engine/usecase/executecode"
)

var (
	problemDomain problemdomain.DomainItf

	ExecutecodeUc *executecodeuc.Usecase
)

func initDomains(cass *cassandralib.CassandraLib) error {
	problemDomain = problemdomain.New(&problemdomain.Resource{
		Cassandra: cass,
	})
	return nil
}

func initUsecases() error {
	ExecutecodeUc = executecodeuc.New()
	return nil
}

func Init() error {
	var err error

	cfg := configlib.Get()

	errorlib.Init(cfg.Error.Prefix)
	loglib.Init(cfg.Log.Logfile)
	os.MkdirAll(cfg.Executor.WorkDir, 0755)

	// cass, err := cassandralib.New(cassandralib.CassandraLibParam{
	// 	ClusterIP: cfg.Cassandra.ClusterIP,
	// 	KeySpace:  cfg.Cassandra.KeySpace,
	// })

	// if err != nil {
	// 	return err
	// }

	// err = initDomains(cass)
	// if err != nil {
	// 	return err
	// }

	err = initUsecases()
	if err != nil {
		return err
	}

	return nil
}
