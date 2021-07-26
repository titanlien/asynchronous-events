package handlers

import (
	"time"

	"github.com/titanlien/asynchronous-events/milestone-4/config"
	"github.com/titanlien/asynchronous-events/milestone-4/events"
	"github.com/titanlien/asynchronous-events/milestone-4/publisher"

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
