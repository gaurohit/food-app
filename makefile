UNAME := $(shell uname)
PWD := $(shell pwd)


.PHONY: dev clean


dev: 
		docker build --tag app .
		docker run -d -p 8081:8083 --name app-server app

clean:
		docker ps -q --filter "name=app" | xargs docker container stop;
		docker ps -a -q --filter "name=app" | xargs docker container rm;
rerun:	clean dev