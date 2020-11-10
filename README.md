# flinkctl
A command line utility for managing remote Flink clusters and applications with a familiar api. 

## Why not use existing binary?
In all honesty you probably should. I wanted to learn what's available from the Flink REST API so I figured I'd build my own client. In addition, `flinkctl` enables you to connect to remove clusters which is something I find useful.

## Status of project
Lots of work to be done still, but currently it can do rudementary tasks. **Use this in production at your own risk, it's probably not stable.**

* [x] Add and manage multiple different (remote/local) clusters (`config add-cluster <url>`, `config use-cluster <existing url>`) 
* [ ] Respect basic auth from config (some code is there, haven't really tested anything)
* [x] Describe the current cluster (`describe cluster`)
* [x] Describe a single job (`describe job <job id>`)
* [x] List running jobs (`get jobs`, `get jobs -d`)
* [x] Submit fat jars (with args) (`submit jar -f /path/to/fat.jar --parallelism 2`)
* [x] List uploaded jars (`get jars`)
* [x] Submit an uploaded jar to cluster (`submit job <uploaded file id>`)
* [x] Stop a running job (`stop job <job id>`)
* [x] Stop running cluster jobmanager (`stop cluster`)
* [x] Trigger job rescale (but it's disabled in newer Flink versions :sadpanda:) (`scale job <job id> --parallelism=4`)
* [x] Generate autocomplete for binary (`generate-autocomplete <shell>`)
* [ ] Json or table output on all prints / consistent printing (some deviations exists)
* [ ] Create job savepoint
* [ ] Monitor misc triggerids
* [ ] List job vertexes
* [ ] Measure backpressure of different vertexes
* [ ] Metrics
* [ ] Clean up `.flinkctl.yaml`-management (at the moment it's pretty garbage)
* [ ] Tests... Tests everywhere.

## Development / Examples
There's a dummy app that runs forever with minimum utilization in `flink-cluster`. In addition, the `Makefile` bootstraps a Flink cluster to your `kubectl config current-context`.

### Build the binary
Since this is a Cobra application (Golang), you'll need Go to build it. Once that's out of your way, you can use `make build` to generate the binary. 

Using `make install` will also move that binary to `~/bin/`.

### Bootstrap a Flink cluster locally
```
# Start Kubernetes locally
minikube start 

# Make sure your current context is minikube
kubectl config current-context #if it's not minikube, change.

# Deploy a Flink cluster
make deploy 

# Port forward the jobmanager to your machine
minikube service flink-jobmanager # Opens up 3 browser tabs, find the Jobmanager master (it's the one with the regular Flink dashboard).

# Configure flinkctl to use your cluster
flinkctl config add-cluster <url to jobmanager, e.g. https://localhost:54321>

# Validate flinkctl can see cluster
flinkctl describe cluster
```

### Build a fat jar
If you have a Flink job packaged already, skip this.

```
# Go to the dummy app
cd flink-cluster/DummyApp

# Use Maven to build your app
mvn clean package
```

### Start using flinkctl features
Output from a freshly deployed (example) cluster.

```
flinkctl describe cluster -o json
# Outputs:
# {
#     "flink-commit": "d04872d",
#     "flink-version": "1.11.0",
#     "jobs-cancelled": 0,
#     "jobs-failed": 0,
#     "jobs-finished": 0,
#     "jobs-running": 0,
#     "slots-available": 4,
#     "slots-total": 4,
#     "taskmanagers": 2
# }

flinkctl submit jar -f flink-cluster/DummyApp/target/DummyApp-1.0-SNAPSHOT.jar
# Outputs:
#  FILENAME                                                         STATUS   
# ---------------------------------------------------------------- --------- 
#  066d0a1e-6e9d-44b6-af36-4a72727bc7b6_DummyApp-1.0-SNAPSHOT.jar   success  

flinkctl submit job 066d0a1e-6e9d-44b6-af36-4a72727bc7b6_DummyApp-1.0-SNAPSHOT.jar
# Outputs:
#  JOBID                             
# ---------------------------------- 
#  b68454ec5d7ca9e13669687871cea629  

flinkctl get jobs
# Outputs:
#  ID                                 STATUS   
# ---------------------------------- --------- 
#  b68454ec5d7ca9e13669687871cea629   RUNNING  

flinkctl describe job b68454ec5d7ca9e13669687871cea629
# Outputs a lot, doesn't fit README. Even more with -o json

flinkctl stop job b68454ec5d7ca9e13669687871cea629
# Outputs:
# Successfully cancelled job b68454ec5d7ca9e13669687871cea629

flinkctl get jobs
# Outputs: 
#  ID                                 STATUS    
# ---------------------------------- ---------- 
#  b68454ec5d7ca9e13669687871cea629   CANCELED  
```

## Contributing
Feel free to contribute as much as you want. I'm very open to suggestions and improvements. There's a lot of unattended TODO's that needs work, and that's just the tip of the ice berg.