CLIENT_DIR = client
SERVER_DIR = server

.DEFAULT_GOAL := client 

server:
	@echo "Running the server..."
	cd $(SERVER_DIR) && go run .

client:
	@echo "Running the client..."
	cd $(CLIENT_DIR) && go run .

build: server_build client_build

server_build:
	@echo "Building the server..."
	cd $(SERVER_DIR) && go build -o ../server_app .

client_build:
	@echo "Building the client..."
	cd $(CLIENT_DIR) && go build -o ../client_app .

clean:
	@echo "Cleaning up binaries..."
	rm -f server_app client_app

# Prevent targets from conflicting with files of the same name
.PHONY: server client build server_build client_build clean