package loglib

import (
	"os"

	"github.com/beruangcoklat/code-execution-engine/pkg/lib/errorlib"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func Init(filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	logrus.SetOutput(f)
	log.SetFormatter(&log.JSONFormatter{})
	return nil
}

func ErrorTrace(err error, msg string) {
	cerr, ok := err.(*errorlib.CustomError)
	fields := log.Fields{}
	if ok {
		fields = log.Fields{
			"traces": cerr.Traces,
			"err":    cerr.Error(),
		}
	} else {
		fields = log.Fields{
			"err": err.Error(),
		}
	}
	log.WithFields(fields).Error(msg)
}
