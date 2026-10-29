#! /usr/bin/make
GIT_COMMIT_COUNT = $(shell git rev-list --count HEAD)
VERSION =v0.0.$(GIT_COMMIT_COUNT)

.DEFAULT_GOAL := help
help:  ## Show this help
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk -F ':.*## ' '{printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

version: ## Show software version
	@echo $(VERSION)

release: ## Release Go module version via new git tag
	@git checkout master
	@git pull
	@echo "Tagging new release of web/pkg as $(VERSION)"
	@git tag $(VERSION)
	@git push --tags
