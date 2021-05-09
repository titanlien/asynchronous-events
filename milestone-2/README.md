# Overview of Milestone 2 Implementation 

# What Was Created?
In this milestone:
1. a Go microservice was created called [*Order*](./order/main.go)
1. a Go function for publishing events to a Kafka topic was created called [*PublishEvent*](./publisher/publisher.go)

# How do I Test?
I was able to test all of the code created in this milestone on my local machine. The instructions below assume you are running on your local machine. I implemented this on a Mac, so references to the command-line will show as a UNIX shell.

1. Kafka and Zookeeper need to be running
1. The *OrderReceived* topic should be created
1. The *Order* service needs to be running (assumes you are in the `milestone2/code` folder)
    1. If this is the first time you are running this code, you will need to setup Go modules
        1. In your `~/.bash_profile`, make sure you have the following ENV var set: `export GO111MODULE=on` and make sure the file is sourced.
        1. Initialize Go modules
            ```shell
            $ go mod init
            ```
            ```shell
            $ go mod tidy
            ```
    1. Run the *Order* service
        ```shell
        $ go run order/main.go
        ```
1. Send a HTTP request to the order service:
    ```shell
    $ curl -v -H "Content-Type: application/json" -d '{"id":"6e042f29-350b-4d51-8849-5e36456dfa48","products":[{"productCode":"12345","quantity":2}],"customer":{"firstName":"Tom","lastName":"Hardy","emailAddress":"tom.hardy@email.com","shippingAddress":{"line1":"123 Anywhere St","city":"Anytown","state":"AL","postalCode":"12345"}}}' http://localhost:8080/orders
    ```
1. You should see output in the console of the order service, and no errors. You can also check the contents of the *OrderReceived* topic in Kafka.
    ```shell
    $ $KAFKA_HOME/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic OrderReceived --from-beginning
    ```