up: ## Up docker container
	sudo rm -rf devops/local/db-data/*
	@docker-compose up --build

down: ## Down docker container
	@docker-compose down

