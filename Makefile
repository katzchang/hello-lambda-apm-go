
build: fmt
	go build main.go

fmt:
	go fmt .

NEW_RELIC_LINKED_ACCOUNT_NAME=hello-lambda-apm
install:
	newrelic-lambda integrations install \
	 --aws-profile $(AWS_PROFILE) \
	 --nr-account-id $(NEW_RELIC_ACCOUNT_ID) \
	 --linked-account-name $(NEW_RELIC_LINKED_ACCOUNT_NAME) \
	 --nr-api-key $(NEW_RELIC_API_KEY)

install-cli:
	pip3 install newrelic-lambda-cli
