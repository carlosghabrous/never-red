.PHONY: usage bash build coverage docs down down-clean-db test up

DC_RUN=docker-compose run --rm web $(1)

usage: 
	@echo "Targets:"
	@echo "- bash			-----> 		Opens a bash shell in the app container"
	@echo "- build			-----> 		Builds services listed in the compose file"
	@echo "- coverage   	----->		Calls go test with coverage tools"
	@echo "- docs			----->		Generates swagger json file"
	@echo "- down 			-----> 		Stops the never-red application, supporting services and removes their containers"
	@echo "- down-clean-db 	-----> 		Same as down, but also removes data volumes"
	@echo "- test			-----> 		Runs tests in never-red packages"
	@echo "- up 			-----> 		Runs the never-red application and supporting services"

bash: 
	$(call DC_RUN, bash)

build: 
	./scripts/composer-build.sh

coverage:
	PKG_LIST=$(go list ./... | grep -v /vendor/ | tr '\n' ' ')
	$(call DC_RUN, go test -covermode=count -coverprofile coverage $(PKG_LIST) && go tool cover -func=coverage)

docs:
	@echo "TBD"

down: 
	docker-compose down

down-clean-db:
	docker-compose down --volumes

test:
	$(call DC_RUN, go test -v $(if $(ARGS),$(ARGS),./...))

up: 
	docker-compose up
