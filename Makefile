build:
	GOOS=linux go build
	docker build -t cgbaker/nomad-dag-hack .
