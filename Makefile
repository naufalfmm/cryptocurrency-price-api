SHELL:/bin/bash

DBFILENAME=
DBPATH=
SQLPATH=
ENVFILENAME?=.env
VERSION=
PORT=

db_migrate:
	docker run --name cryptocurrency-price-api-migration -v $(SQLPATH):/usr/src/migrations/sql -v $(DBPATH)/$(DBFILENAME):/usr/src/$(DBFILENAME) --env-file $(ENVFILENAME) --rm naufalfmm/cryptocurrency-price-api-migration ./migrations/cryptocurrency-price-api-migration migrate

db_rollback:
	docker run --name cryptocurrency-price-api-migration -v $(SQLPATH):/usr/src/migrations/sql -v $(DBPATH)/$(DBFILENAME):/usr/src/$(DBFILENAME) --env-file $(ENVFILENAME) --rm naufalfmm/cryptocurrency-price-api-migration ./migrations/cryptocurrency-price-api-migration rollback $(if $(strip $(VERSION)), --version $(VERSION))

db_create:
	docker run --name cryptocurrency-price-api-migration -v $(SQLPATH):/usr/src/migrations/sql --env-file $(ENVFILENAME) --rm naufalfmm/cryptocurrency-price-api-migration ./migrations/cryptocurrency-price-api-migration create --name $(NAME)

db_init:
	docker build -t naufalfmm/cryptocurrency-price-api-migration:latest -f .\dockerfile\Dockerfile.migration .

db:
	db_init && db_migrate

app_init:
	docker build -t naufalfmm/cryptocurrency-price-api:latest -f .\dockerfile\Dockerfile.app .

app_run:
	docker run --name cryptocurrency-price-api -p $(PORT):$(PORT) -v $(DBPATH)/$(DBFILENAME):/usr/src/$(DBFILENAME) --env-file $(ENVFILENAME) --rm naufalfmm/cryptocurrency-price-api

app:
	app_init && app_run

run:
	db && app