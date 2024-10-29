PROJECT_NAME=ip_info_service
.PHONY: build clean

build:
	go build -o $(PROJECT_NAME) ./cmd/main/main.go

build_static:
	CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o $(PROJECT_NAME)_static ./cmd/main/main.go

clean:
	rm -f $(PROJECT_NAME)

run: build
	./$(PROJECT_NAME)
