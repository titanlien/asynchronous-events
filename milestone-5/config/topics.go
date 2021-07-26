package config

const (
	// OrderReceivedTopicName is the name of the topic that handles OrderReceived events
	OrderReceivedTopicName = "OrderReceived"

	// OrderConfirmedTopicName is the name of the topic that handles OrderConfirmed events
	OrderConfirmedTopicName = "OrderConfirmed"

	// NotificationTopicName is the name of the topic that handles Notification events
	NotificationTopicName = "Notification"

	// OrderPickedAndPackedTopicName is the name of the topic that handles OrderPickedAndPacked events
	OrderPickedAndPackedTopicName = "OrderPickedAndPacked"

	// ErrorsTopicName is the name of the topic that handles Error events
	ErrorsTopicName = "DeadLetterQueue"
)
