dev-up:
	docker-compose up -d
dev-down:
	docker-compose down

run:
	cd cmd/ordersystem/ && go run main.go wire_gen.go
