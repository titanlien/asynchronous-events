#!/bin/bash

# Set this to the location where Kafka has been installed
KAFKA_HOME=~/Documents/reading/kafka-asynchronous/kafka_2.13-2.8.0

# Start Zookeeper
$KAFKA_HOME/bin/zookeeper-server-start.sh $KAFKA_HOME/config/zookeeper.properties
