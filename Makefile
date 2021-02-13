gobuild:
	@go build -v -o code-execution-engine cmd/code-execution-engine/main.go

gorun_nsq:
	make gobuild
	@./code-execution-engine -type=nsq -config_path=configs/config.yaml

kompos:
	@docker-compose -f docker/nsq.yml up -d
