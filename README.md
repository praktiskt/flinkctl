# flinkctl
A command line utility for managing remote Flink clusters and applications with a familiar api. 

# Status of project
Lots of work to be done still, but currently it can do rudementary tasks. **Use this in production at your own risk, it's probably not stable.**

* [x] Add and manage multiple different (remote/local) clusters (`config add-cluster <url>`, `config use-cluster <existing url>`) 
* [ ] Respect basic auth from config (some code is there, haven't really tested anything)
* [x] Describe the current cluster (`describe cluster`)
* [x] Describe a single job (`describe job <job id>`)
* [x] List running jobs (`get jobs`, `get jobs -d`)
* [x] Submit fat jars (with args) (`submit jar /path/to/fat.jar --parallelism 2`)
* [x] List uploaded jars (`get jars`)
* [x] Submit an uploaded jar to cluster (`submit job <uploaded file id>`)
* [x] Stop a running job (`stop job <job id>`)
* [x] Stop running cluster jobmanager (`stop cluster`)
* [x] Trigger job rescale (but it's disabled in newer Flink versions :sadpanda:) (`scale job <job id> --parallelism=4`)
* [x] Generate autocomplete for binary (`generate-autocomplete <shell>`)
* [ ] Json or table output on all prints / consistent printing (some deviations exists)
* [ ] Create job savepoint
* [ ] Monitor misc triggerids
* [ ] Clean up `.flinkctl.yaml`-management (at the moment it's pretty garbage)
* [ ] Clean up help
* [ ] Tests... Tests everywhere.