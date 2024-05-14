SHELL:/bin/bash

DBPATH?=
ENVFILENAME?=.env
VERSION?=

db_migrate:
	docker run --name cryptocurrency-price-api-migration -v $(DBPATH):/migrations/sql --env-file $(ENVFILENAME) --rm naufalfmm/cryptocurrency-price-api-migration ./migrations/cryptocurrency-api/migration migrate

db_rollback:
	docker run --name cryptocurrency-price-api-migration -v $(DBPATH):/migrations/sql --env-file $(ENVFILENAME) --rm naufalfmm/cryptocurrency-price-api-migration ./migrations/cryptocurrency-api/migration rollback $(if $(strip $(VERSION)), --version $(VERSION))

db_create:
	docker run --name cryptocurrency-price-api-migration -v $(DBPATH):/migrations/sql --env-file $(ENVFILENAME) --rm naufalfmm/cryptocurrency-price-api-migration ./migrations/cryptocurrency-api/migration create --name $(NAME)

db_init:
	docker build -t naufalfmm/cryptocurrency-price-api-migration:latest -f .\dockerfile\Dockerfile.migration .

db:
	db_init && db_migrate

app_init:
	docker build -t naufalfmm/cryptocurrency-price-api:latest -f .\dockerfile\Dockerfile.app .

app_run:
	docker run --name cryptocurrency-price-api --env-file $(ENVFILENAME) --rm naufalfmm/cryptocurrency-price-api

app:
	app_init && app_run

run:
	db && app