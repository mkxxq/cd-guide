
.PHONY: linux clean default build

output = _build
compile = go build -a -ldflags '-s -w --extldflags "-static -fpic"' -o
tag = v1

default:
	$(compile) $(output)/hello src/main.go
linux:
	GOOS=linux GOARCH=amd64 $(compile) $(output)/helloworld src/main.go
build: linux
	docker build -t helloworld -f docker/Dockerfile .
clean:
	rm hello