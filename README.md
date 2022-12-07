# week1-blogging-system

## Description:
Built scalable blogging system, which allow user to add blog, search blog and delete blog at scale.

## Architecture(HLD):
<img width="922" alt="Screenshot 2022-12-05 at 1 41 01 PM" src="https://user-images.githubusercontent.com/39022530/205586378-2bffc913-b547-48d6-ab51-5d2249bf5ecd.png">

This system is designed in such a way that, it can scale up to million of user. I have divide my project into two part, that is dashboard and consumer, here number of parttion will be equal to number of consumer. All the compoments are described below.

1. Blog Dashboard:
This compoment is responsible for pushing all the events from the user like, creation, reading and deletion. It will be isolated and standlone application that can scale without effecting all the consumer as both are decoupled.

2. ZooKeeper:
This is also known as service discovery, that will keep the record of all the broker in the kafka and verify. It will perform co-ordination between different node and store all the meta-data of the machine with their broker.

3. Kafka:
It is distributed streaming platform, that will collect all the event produced from the dashboard and store in the buffer. later it will be consumered by the consumer, running in the background. It will guarantee atleast once delivery and make available all the event to the consumers. It also decouple all the different applications and provide flexibility.

4. Blog Consumer:
It will collect all the events from the kafka and perform all the backgroud operations in A-Sync ways and provide high performance to the system with very low latency.

#### Blog DashBoard


## Tech Stack:
1. golang
2. kafka
3. zookeeper
4. elasticsearch

## Set Up:
1. Install elasticsearch
2. Install docker

## How to run project?
1. Invoke docker
```
  docker-compose up
```

2. Invoke BlogDashBoard
```
  go run .
```

3. Invoke BlogConsumer
```
  go run .
```

## Requests:
1. Create blog
```
curl --location --request POST 'localhost:8000/api/write_blog' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":"golang",
    "body":"golang is amazing for the backeng service",
    "author":"abhishek gupta"
}'
```
2. Read Blog:
```
curl --location --request GET 'localhost:8000/api/read_blog?title=golang'
```
3. Delete Blog:
```
curl --location --request DELETE 'localhost:8000/api/delete_blog?title=golang'
```
