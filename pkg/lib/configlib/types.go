package configlib

type (
	Config struct {
		NSQ       NsqConfig       `yaml:"nsq"`
		Cassandra CassandraConfig `yaml:"cassandra"`
		Executor  ExecutorConfig  `yaml:"executor"`
		Error     ErrorConfig     `yaml:"error"`
		Log       LogConfig       `yaml:"log"`
	}

	NsqConfig struct {
		NsqLookupd string `yaml:"nsqlookupd"`
		Nsqd       string `yaml:"nsqd"`
	}

	CassandraConfig struct {
		ClusterIP string `yaml:"cluster_ip"`
		KeySpace  string `yaml:"keyspace"`
	}

	ExecutorConfig struct {
		WorkDir string `yaml:"workdir"`
	}

	ErrorConfig struct {
		Prefix string `yaml:"prefix"`
	}

	LogConfig struct {
		Logfile string `yaml:"logfile"`
	}
)
