BINPATH=/build
SERVER_BIN=$(BINPATH)/server
SERVER_SRC=cmd/api/main.go
GOCMD=go
GOTEST=$(GOCMD) test -covermode=count -coverprofile=coverage.out ./pkg/...
GOCOVER=$(GOCMD) tool cover -html=coverage.out
GOBUILD_SERVER=$(GOCMD) build -o $(SERVER_BIN) -v ./$(SERVER_SRC)

test:
	$(GOTEST)
cover: test
	$(GOCOVER)
build: test
	$(GOBUILD_SERVER)
run: build
	$(SERVER_BIN)
clear:
	rm -fv $(SERVER_BIN)
build-db:
	docker cp ./sql/schema.sql cashmachine-db:/schema.sql
	docker cp ./sql/table.sql cashmachine-db:/table.sql
	docker exec -it cashmachine-db psql cashmachine -U cash_machine -f /schema.sql -f /table.sql

#docker 
docker-build:
	docker-compose build
docker-up:
	docker-compose up