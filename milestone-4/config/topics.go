package config

const (
	// OrderReceivedTopicName is the name of the topic that handles OrderReceived events
	OrderReceivedTopicName = "OrderReceived"

	// OrderConfirmedTopicName is the name of the topic that handles OrderConfirmed events
	OrderConfirmedTopicName = "OrderConfirmed"

	// ErrorsTopicName is the name of the topic that handles Error events
	ErrorsTopicName = "DeadLetterQueue"
)
