#!/bin/bash
YAML_PATH="flink-cluster"
for f in $(ls $YAML_PATH | grep yaml); do
    kubectl $@ -f $YAML_PATH/$f
done