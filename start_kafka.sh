#!/bin/bash

# Set this to the location where Kafka has been installed
KAFKA_HOME=~/Tools/kafka_2.12-2.5.0

# Start Kafka
$KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/server.properties
