#!/bin/bash
YAML_PATH="flink-cluster"
for f in $(ls $YAML_PATH | grep yaml); do
    kubectl delete -f $YAML_PATH/$f
done