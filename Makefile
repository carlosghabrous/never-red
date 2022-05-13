.PHONY: usage up down down-clean-db test

DC_RUN=docker-compose run --rm web $(1)

usage: 
	@echo "Targets:"
	@echo "- build			-----> 		Builds services listed in the compose file"
	@echo "- down 			-----> 		Stops the never-red application, supporting services and removes their containers"
	@echo "- down-clean-db 	-----> 		Same as down, but also removes data volumes"
	@echo "- up 			-----> 		Runs the never-red application and supporting services"

build: 
	./scripts/composer-build.sh

down: 
	docker-compose down

down-clean-db:
	docker-compose down --volumes

test:
	$(call DC_RUN, go test -v $(if $(ARGS),$(ARGS),./...))

up: 
	docker-compose up
