# Overview of Milestone 3 Implementation 

# What Was Created?
In this milestone:
1. a Go Kafka consumer was created called [*Inventory*](./inventory/main.go)
1. a Postgres database is used as a mechanism to handle duplicate message processing: [click here for more information](./db/db.go)

# How to Test?
I was able to test all of the code created in this milestone on my local machine. The instructions below assume you are running on your local machine. I implemented this on a Mac, so references to the command-line will show as a UNIX shell.

1. Kafka and Zookeeper need to be running
    1. The *OrderReceived* topic should be created
    1. The *DeadLetterQueue* topic should be created
1. The Postgress database needs to be running
    1. The right database, schema and table need to be created: [click here for more information](./db/db.go)
1. The *Order* service needs to be running (assumes you are in the `milestone4/code` folder)
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
1. The *Inventory* consumer needs to be running (assumes you are in the `milestone4/code` folder)
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
1. Send a HTTP request to the order service:
    ```shell
    $ curl -v -H "Content-Type: application/json" -d '{"id":"6e042f29-350b-4d51-8849-5e36456dfa48","products":[{"productCode":"12345","quantity":2}],"customer":{"firstName":"Tom","lastName":"Hardy","emailAddress":"tom.hardy@email.com","shippingAddress":{"line1":"123 Anywhere St","city":"Anytown","state":"AL","postalCode":"12345"}}}' http://localhost:8080/orders
    ```
1. You should see output in the console of the order service, and no errors. You can also check the contents of the *OrderReceived* topic in Kafka.
    ```shell
    $ $KAFKA_HOME/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic OrderReceived --from-beginning
    ```
1. You should see output in the console of the inventory consumer, and no errors.