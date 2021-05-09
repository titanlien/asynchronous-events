#!/bin/bash

# Set this to the location where Kafka has been installed
KAFKA_HOME=~/Tools/kafka_2.12-2.5.0

# Start Zookeeper
$KAFKA_HOME/bin/zookeeper-server-start.sh $KAFKA_HOME/config/zookeeper.properties