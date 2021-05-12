package publisher

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	"milestone-2/config"
	"milestone-2/events"
)

// PublishEvent will publish the specified event to the messaging system (currently running on localhost)
func PublishEvent(event events.Event, topic string) error {

	log.WithField("event", event).Info("attempting to publish event")

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":   config.BrokerAddress(),
		"socket.timeout.ms":   30000,
		"delivery.timeout.ms": 30000})
	if err != nil {
		return err
	}

	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)

	var value []byte
	if value, err = json.Marshal(event); err != nil {
		return err
	}
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	log.WithField("Name", *m.TopicPartition.Topic).
		WithField("Partition", m.TopicPartition.Partition).
		WithField("PartitionOffset", m.TopicPartition.Offset).
		Infof("Delivered message to topic")

	close(deliveryChan)

	return nil
}
