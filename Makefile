# Binary name
BINARY=main

# Run project locally
run: 
	go run cmd/server/main.go

# Build binary
build:
	go build -o $(BINARY) cmd/server

# Run docker container
up:
	docker-compose up -d

# Stop docker container
down:
	docker-compose down

# Rebuild docker container
rebuild:
	docker-compose up --build -d

# Run database with docker
db:
	docker-compose up -d db

# Run the database in the shell
psql:
	docker exec -it schedule_db psql -U schedule_user -d schedule

# Show logs
logs:
	docker logs -f schedule_backend

# Clean binary and dependencies
clean:
	rm -f $(BINARY)
	go clean