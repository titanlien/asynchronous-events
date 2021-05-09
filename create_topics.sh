#!/bin/bash

# Set this to the location where Kafka has been installed
KAFKA_HOME=~/Tools/kafka_2.12-2.5.0

# Create the OrderReceived topic
$KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderReceived

# Create the OrderConfirmed topic
$KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderConfirmed

# Create the OrderPickedAndPacked topic
$KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderPickedAndPacked

# Create the Notification topic
$KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic Notification

# Create the DeadLetterQueue topic
$KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic DeadLetterQueue