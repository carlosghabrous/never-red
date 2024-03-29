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

checks: build \
	vet \
	test \
	coverage

coverage:
	$(call DC_RUN, go test --count=1 -coverprofile=cover.out ./...)

docs:
	@echo "TBD"

down: 
	docker-compose down --remove-orphans

down-clean-db:
	docker-compose down --volumes

test:
	$(call DC_RUN, go test -v --count=1 $(if $(ARGS),$(ARGS),./...))

up:
	docker-compose up

vet:
	$(call DC_RUN, go vet ./...)
