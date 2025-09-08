all: clean build docker

clean:
	rm rehabber-discord-integration -f

build:
	go build -v

docker: build
	docker build -t jbuchbinder/rehabber-discord-integration .
	docker push jbuchbinder/rehabber-discord-integration
