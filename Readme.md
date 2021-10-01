# Go Concurrent test project

## 1. Pre-requisition

- Kafka single node
- MongoDB
- Go >= 1.16

## 2. Environment Variables

Create `.env` file with following variables

```sh
MONGO_URI=mongodb://localhost:27017/
DBNAME=test_sync
KAFKA_URLS=localhost:9092
KAFKA_TOPIC=test_sync
WORKER_GROUP=with_buffer
MAX_WORKERS=5
BUFFER_LENGTH=2
```

## 3. How to run

1. Feed Kafka with data using `feed`
2. Update WORKER_GROUP for running each tests:

- `serial` - Run without concurrent
- `single_pipeline` - One goroutine per stage
- `pipelines` - Multiple Goroutine per stage
- `with_buffer` - Multiple Goroutine per stage, with buffered channels
