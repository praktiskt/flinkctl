deploy:
	./flink-cluster/deploy.sh apply

undeploy:
	./flink-cluster/deploy.sh delete

install:
	go build -o flinkctl && mv flinkctl ~/bin/