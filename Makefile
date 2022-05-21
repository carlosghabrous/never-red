.PHONY: usage build coverage docs docs-serve down down-clean-db test up

DC_RUN=docker-compose run --rm web $(1)

usage: 
	@echo "Targets:"
	@echo "- build			-----> 		Builds services listed in the compose file"
	@echo "- down 			-----> 		Stops the never-red application, supporting services and removes their containers"
	@echo "- down-clean-db 	-----> 		Same as down, but also removes data volumes"
	@echo "- up 			-----> 		Runs the never-red application and supporting services"

build: 
	./scripts/composer-build.sh

coverage:
	PKG_LIST=$(go list ./... | grep -v /vendor/ | tr '\n' ' ')
	$(call DC_RUN, go test -covermode=count -coverprofile coverage $(PKG_LIST) && go tool cover -func=coverage)

docs:
	swagger generate spec -o ./swagger/swagger.json --scan-models

docs-serve:
	swagger serve ./swagger/swagger.json

down: 
	docker-compose down

down-clean-db:
	docker-compose down --volumes

test:
	$(call DC_RUN, go test -v $(if $(ARGS),$(ARGS),./...))

up: 
	docker-compose up
