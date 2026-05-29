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
	  -d '{"to":"a@b.com","subject":"hi","body":"hello"}'
