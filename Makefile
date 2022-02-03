test: 
	go test -short -v ./... -count=1

lint:
	golangci-lint run
	
docker:
	docker build -t matchescli .

.PHONY: test docker