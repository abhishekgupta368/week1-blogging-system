# week1-blogging-system

## Description:
Built scalable blogging system, which allow user to add blog, search blog and delete blog at scale.

## Architecture(HLD):
<img width="922" alt="Screenshot 2022-12-05 at 1 41 01 PM" src="https://user-images.githubusercontent.com/39022530/205586378-2bffc913-b547-48d6-ab51-5d2249bf5ecd.png">

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
