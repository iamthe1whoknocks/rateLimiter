build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./cmd/app/app ./cmd/app/

test:
	go test -v -race ./handler

run:build
	docker build -t rate-limiter-scratch -f Dockerfile.scratch .
	docker run -it -p 8083:8083 rate-limiter-scratch
