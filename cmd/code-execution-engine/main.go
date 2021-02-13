package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/beruangcoklat/code-execution-engine/pkg/lib/configlib"
	"github.com/beruangcoklat/code-execution-engine/server"
	"github.com/beruangcoklat/code-execution-engine/server/nsqserver"
)

type flagParam struct {
	appType    string
	configPath string
}

func getFlagParam() flagParam {
	param := flagParam{}

	flag.StringVar(&param.appType, "type", "nsq", "app type")
	flag.StringVar(&param.configPath, "config_path", "configs/config.yaml", "config path")
	flag.Parse()

	return param
}

func main() {
	var err error

	flagParam := getFlagParam()

	err = configlib.Init(flagParam.configPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = server.Init()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	switch flagParam.appType {
	case "nsq":
		err = nsqserver.InitNSQ()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		wait()
	default:
		fmt.Println("app type not found")
		os.Exit(1)
	}
}

func wait() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
