up:
	docker compose up --build

down:
	docker compose down

rabbit:
	docker compose up -d rabbitmq

rabbit_down:
	docker compose down rabbitmq

send:
	curl -X POST localhost:8080/emails \
	  -H "Content-Type: application/json" \
	  -d '{"to":"aakash19here@gmail.com","subject":"hi","body":"hello"}'

run_producer:
	go run ./cmd/producer
run_consumer:
	go run ./cmd/consumer
