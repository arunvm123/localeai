# Task Description

XRides, delivers about a 200 rides per minute or 288,000 rides per day. Now, they want to send
this data to your system via an API. Your task is to create this API and save the data into
PostgreSQL.
The API should be designed, keeping in mind the real-time streaming nature of data and the
burst of requests at peak times of the day. The user of this API expects an acknowledgment that
the data is accepted and a way to track if the request fails.

# Implementation

![Alt](https://i.imgur.com/zheO4Jt.jpg)

The service is dependant on 
*[NATS](https://nats.io/) (Message broker)
*Postgres

# How to run?

Create a yaml file for configuration based on `config-example.yaml`

Setup nats messaging system with the help of docker by running the command

`docker run -p 4222:4222 -ti nats:latest`

To run the producer, move to `cmd/producer/` and run the following command

` go run . -config-path=path/to/config `

Similarly for the consumer, move to `cmd/consumer/` and run the following command

` go run . -config-path=path/to/config `

Now the data can be send as a POST request in the body. The url is `http://localhost:<PORT>/save/booking/detail`