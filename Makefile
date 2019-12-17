
GOOS=linux
build: fmt
	GOOS=$(GOOS) go build main.go

fmt:
	go fmt .

function-name=hello-lambda-apm
update-function-code: main.zip
	aws lambda update-function-code \
	 --function-name $(function-name) \
	 --zip-file fileb://./$<

main.zip: main
	rm -rf $@
	zip $@ $^

main: build

NEW_RELIC_LINKED_ACCOUNT_NAME=hello-lambda-apm
install:
	newrelic-lambda integrations install \
	 --aws-profile $(AWS_PROFILE) \
	 --nr-account-id $(NEW_RELIC_ACCOUNT_ID) \
	 --linked-account-name $(NEW_RELIC_LINKED_ACCOUNT_NAME) \
	 --nr-api-key $(NEW_RELIC_API_KEY)

install-cli:
	pip3 install newrelic-lambda-cli
