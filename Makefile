BINPATH=./bin
SERVER_BIN=$(BINPATH)/server
SERVER_SRC=cmd/main.go
GOCMD=go
GOTEST=$(GOCMD) test -covermode=count -coverprofile=coverage.out ./pkg/...
GOCOVER=$(GOCMD) tool cover -html=coverage.out
GOBUILD_SERVER=$(GOCMD) build -o ./$(SERVER_BIN) -v ./$(SERVER_SRC)

test:
	$(GOTEST)
cover: test
	$(GOCOVER)
build: test
	$(GOBUILD_SERVER)
run: build
	$(SERVER_BIN)
clear:
	rm -fv $(BINPATH)/*


#docker 
docker-build:
	docker-compose build
docker-up:
	docker-compose up