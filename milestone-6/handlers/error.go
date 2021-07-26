package handlers

import (
	"time"

	"github.com/bpmericle/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/milestone6/code/config"
	"github.com/bpmericle/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/milestone6/code/events"
	"github.com/bpmericle/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/milestone6/code/publisher"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// HandleError will publish an error event to Kafka
func HandleError(event events.Event) {
	var err error

	e := translateToErrorEvent(event)
	if err = publisher.PublishEvent(e, config.ErrorsTopicName); err != nil {
		log.WithField("error", err).
			WithField("topic", config.ErrorsTopicName).
			Error("an issue ocurred publishing an error event to Kafka")
	}
}

func translateToErrorEvent(event events.Event) events.Event {
	return events.Error{
		EventBase: events.BaseEvent{
			EventID:        uuid.New(),
			EventTimestamp: time.Now(),
		},
		EventBody: event,
	}
}