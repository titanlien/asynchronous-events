# Overview of Milestone 4 Implementation 

# What Was Created?
In this milestone:
1. a Go Kafka consumer was created called [*Warehouse*](./warehouse/main.go)
1. a Go Kafka consumer was created called [*Notification*](./notification/main.go)
1. a Go Kafka consumer was created called [*Shipper*](./shipper/main.go)

# How to Test?
I was able to test all of the code created in this milestone on my local machine. The instructions below assume you are running on your local machine. I implemented this on a Mac, so references to the command-line will show as a UNIX shell.

1. Kafka and Zookeeper need to be running
    1. The *OrderReceived* topic should be created
    1. The *OrderPickedAndPacked* topic should be created
    1. The *Notification* topic should be created
    1. The *DeadLetterQueue* topic should be created
1. The Postgress database needs to be running
    1. The right database, schema and table need to be created: [click here for more information](./db/db.go)
1. The *Order* service needs to be running (assumes you are in the `milestone5/code` folder)
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
1. The *Inventory* consumer needs to be running (assumes you are in the `milestone5/code` folder)
    1. If this is the first time you are running this code, you will need to setup Go modules
        1. In your `~/.bash_profile`, make sure you have the following ENV var set: `export GO111MODULE=on` and make sure the file is sourced.
        1. Initialize Go modules
            ```shell
            $ go mod init
            ```
            ```shell
            $ go mod tidy
            ```
    1. Run the *Inventory* consumer service
        ```shell
        $ go run inventory/main.go
        ```
1. The *Warehouse* consumer needs to be running (assumes you are in the `milestone5/code` folder)
    1. If this is the first time you are running this code, you will need to setup Go modules
        1. In your `~/.bash_profile`, make sure you have the following ENV var set: `export GO111MODULE=on` and make sure the file is sourced.
        1. Initialize Go modules
            ```shell
            $ go mod init
            ```
            ```shell
            $ go mod tidy
            ```
    1. Run the *Warehouse* consumer service
        ```shell
        $ go run warehouse/main.go
        ```
1. The *Notification* consumer needs to be running (assumes you are in the `milestone5/code` folder)
    1. If this is the first time you are running this code, you will need to setup Go modules
        1. In your `~/.bash_profile`, make sure you have the following ENV var set: `export GO111MODULE=on` and make sure the file is sourced.
        1. Initialize Go modules
            ```shell
            $ go mod init
            ```
            ```shell
            $ go mod tidy
            ```
    1. Run the *Notification* consumer service
        ```shell
        $ go run notification/main.go
        ```
1. The *Shipper* consumer needs to be running (assumes you are in the `milestone5/code` folder)
    1. If this is the first time you are running this code, you will need to setup Go modules
        1. In your `~/.bash_profile`, make sure you have the following ENV var set: `export GO111MODULE=on` and make sure the file is sourced.
        1. Initialize Go modules
            ```shell
            $ go mod init
            ```
            ```shell
            $ go mod tidy
            ```
    1. Run the *Shipper* consumer service
        ```shell
        $ go run shipper/main.go
        ```
1. Send a HTTP request to the order service:
    ```shell
    $ curl -v -H "Content-Type: application/json" -d '{"id":"6e042f29-350b-4d51-8849-5e36456dfa48","products":[{"productCode":"12345","quantity":2}],"customer":{"firstName":"Tom","lastName":"Hardy","emailAddress":"tom.hardy@email.com","shippingAddress":{"line1":"123 Anywhere St","city":"Anytown","state":"AL","postalCode":"12345"}}}' http://localhost:8080/orders
    ```
1. You should see output in the console of the order service, and no errors. You can also check the contents of the *OrderReceived* topic in Kafka.
    ```shell
    $ $KAFKA_HOME/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic OrderReceived --from-beginning
    ```
1. You can also check the contents of the *Notification* topic in Kafka.
    ```shell
    $ $KAFKA_HOME/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic Notification --from-beginning
    ```
1. You should see output in the console of the inventory consumer, and no errors.
1. You should see output in the console of the warehouse consumer, and no errors.
1. You should see output in the console of the notification consumer, and no errors.
1. To check the the shipping consumer, you will need to manually publish a message to the *OrderPickedAndPacked* topic in Kafka.
    ```shell
    $ $KAFKA_HOMEbin/kafka-console-producer.sh --bootstrap-server localhost:9092 --topic OrderPickedAndPacked
    ```
    Then you can past a payload like this
    ```json
    {"EventBase":{"EventID":"4a651ef8-a851-4d77-a58b-3d8af748a570","EventTimestamp":"2020-08-16T16:03:05.258542-04:00"},"EventBody":{"id":"c6b37316-b4da-4b25-94c8-14c08bad95e6","products":[{"productCode":"12345","quantity":2}],"customer":{"firstName":"Tom","lastName":"Hardy","emailAddress":"tom.hardy@email.com","shippingAddress":{"line1":"123 Anywhere St","city":"Anytown","state":"AL","postalCode":"12345"}}}}
    ```
1. You should see output in the console of the shipper consumer, and no errors.