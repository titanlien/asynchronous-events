#!/bin/bash

# Set this to the location where Kafka has been installed
KAFKA_HOME=~/Documents/reading/kafka-asynchronous/kafka_2.13-2.8.0

# Start Kafka
$KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/server.properties
