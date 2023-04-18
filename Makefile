server:
	go run cmd/main.go

build:
	go build -o bin/server cmd/main.go

swagger:
	swag init -g cmd/main.go

d.up:
	docker-compose up 

d.down:
	docker-compose down

d.up.build:
	docker-compose up --build