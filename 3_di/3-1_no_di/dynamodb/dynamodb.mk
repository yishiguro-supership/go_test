### main command
dynamodb-start: ## dynamodb: start in local
	$(MAKE) aws-set-dummy-profile && sleep 2 && \
	$(MAKE) dynamodb-docker-run && sleep 2 && \
	$(MAKE) dynamodb-create-table && sleep 2 && \
	$(MAKE) dynamodb-batch-write-item

dynamodb-stop: ## dynamodb: stop in local
	$(MAKE) dynamodb-docker-stop

dynamodb-restart: ## dynamodb: restart in local
	$(MAKE) dynamodb-stop; \
	$(MAKE) dynamodb-start;


### docker

dynamodb-docker-run: ## docker run
	docker run --name dynamodb-local -d --rm -it \
		-p 8000:8000 amazon/dynamodb-local \
		-jar DynamoDBLocal.jar -inMemory -sharedDb

dynamodb-docker-stop: ## docker stop
	docker stop dynamodb-local


### dynamodb

TABLE_NAME = user
DYNAMODB_OPTIONS = --endpoint-url http://localhost:8000 --profile dummy

TABLE_ATTRIBUTE_DEFINITIONS = AttributeName=uid,AttributeType=S
TABLE_KEY_SCHEMA = AttributeName=uid,KeyType=HASH

dynamodb-create-table: ## create table
	aws dynamodb $(DYNAMODB_OPTIONS) create-table --table-name $(TABLE_NAME) \
		--attribute-definitions $(TABLE_ATTRIBUTE_DEFINITIONS) \
		--key-schema $(TABLE_KEY_SCHEMA) \
		--billing-mode PAY_PER_REQUEST \
		--no-cli-pager

dynamodb-describe-table: ## describe table
	aws dynamodb $(DYNAMODB_OPTIONS) describe-table --table-name $(TABLE_NAME) \
		--no-cli-pager

dynamodb-delete-table: ## delete table
	aws dynamodb $(DYNAMODB_OPTIONS) delete-table --table-name $(TABLE_NAME) \
		--no-cli-pager

dynamodb-batch-write-item: ## batch write item
	aws dynamodb $(DYNAMODB_OPTIONS) batch-write-item \
		--request-items file://dynamodb/items.json --no-cli-pager

dynamodb-list-tables: ## list tables
	aws dynamodb $(DYNAMODB_OPTIONS) list-tables \
	--no-cli-pager

dynamodb-scan: ## scan
	aws dynamodb $(DYNAMODB_OPTIONS) scan --table-name $(TABLE_NAME) \
	--no-cli-pager


### aws
aws-set-dummy-profile: ## set dummy profile
	cat dynamodb/aws_dummy_profile.txt | aws configure --profile dummy
