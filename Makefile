deploy:
	./flink-cluster/deploy.sh apply

undeploy:
	./flink-cluster/deploy.sh delete

build:
	go build -o flinkctl 

install: build
	mv flinkctl ~/bin/