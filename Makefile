test: 
	go test -short -v ./... -count=1
docker:
	docker build -t matchescli .

.PHONY: test docker