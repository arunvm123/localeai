# Task Description

XRides, delivers about a 200 rides per minute or 288,000 rides per day. Now, they want to send
this data to your system via an API. Your task is to create this API and save the data into
PostgreSQL.
The API should be designed, keeping in mind the real-time streaming nature of data and the
burst of requests at peak times of the day. The user of this API expects an acknowledgment that
the data is accepted and a way to track if the request fails.

# Implementation

![Alt](https://i.imgur.com/zheO4Jt.jpg)

I decided to build a microservice architecture with Golang for this task because of it's concurrency model, which is relatively easy to implement, and its design as a static, compiled language would enable a distributed eventing systems that could perform at scale. Also, many companies have made the switch to Golang for their microservices, so it is a proven tool for this space.

I chose [NATS](https://nats.io/) for the message broker because it is a simple, secure, and high-performance open source messaging system for microservice architecture and is written in Golang. It also had some features that made it easier for me to send an acknowledgement after the data is processed.

The producer is the external facing service that XRides can communicate via REST API. It sends the booking detail to it and the producer pushes it to the message broker. The consumer reads it from the message broker and saves it into Postgres. Then the response is pushed into another message broker, which is picked up by the producer that initiated the request.

The response is picked up by the producer and then is relayed back to XRides. If the response is successful, then everything worked perfectly. If there is an error for some reason, XRides can retry the request. The error response will contain the ID of the booking detail and the error message

## Ideal Architecture

- The microservice uses REST and JSON encoding now. I would like to look into gRPC and protobuf
- Use Kubernetes to manage deployment, scaling, load balancing etc
- Store and properly manage logging
- Setup monitoring and alerting

# How to run?

Create a yaml file for configuration based on `config-example.yaml`

Setup nats messaging system with the help of docker by running the command

`docker run -p 4222:4222 -ti nats:latest`

To run the producer, move to `cmd/producer/` and run the following command

` go run . -config-path=path/to/config `

Similarly for the consumer, move to `cmd/consumer/` and run the following command

` go run . -config-path=path/to/config `

Now the data can be send as a POST request in the body. The url is `http://localhost:<PORT>/save/booking/detail`