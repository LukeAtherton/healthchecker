.PHONY:	build push publish launch

build:	
	docker build --no-cache -t docker.hivebase.io:5000/healthchecker .

push:	
	docker push docker.hivebase.io:5000/healthchecker

publish:	
	make build;
	make push;

launch:	
	-fleetctl stop healthchecker.service; 
	fleetctl unload healthchecker.service; 
	-fleetctl destroy healthchecker.service;
	fleetctl submit ./systemd/healthchecker.service; 
	fleetctl start healthchecker.service;
	fleetctl journal -f healthchecker;

deploy:
	make publish;
	make launch;