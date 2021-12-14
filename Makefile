.PHONY: $(MAKECMDGOALS)

setup: ## Install dependencies
	yarn install

generate: ## Rebuild client
	yarn run openapi-generator-cli generate -g go -o msgraph/ -i openapi.fixed.yaml