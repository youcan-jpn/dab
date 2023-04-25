.PHONY: __init-db-args
__init-db-args:
ifndef DB_HOST
	$(warning DB_HOST was not set; localhost is used)
	$(eval DB_HOST := localhost)
endif
ifndef DB_PORT
	$(warning DB_PORT was not set; 3396 is used)
	$(eval DB_PORT := 13306)
endif
ifndef DB_USER
	$(warning DB_USER was not set; root is used)
	$(eval DB_USER := root)
endif
ifndef DB_PASS
	$(warning DB_PASS was not set; passw0rd is used)
	$(eval DB_PASS := passw0rd)
endif
ifndef DB_NAME
	$(warning DB_NAME was not set; aot_image is used)
	$(eval DB_NAME := dab_db)
endif

.PHONY: db-migrate
db-migrate: __init-db-args
	go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
	migrate -source "file://ddl" -database "mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" down
	migrate -source "file://ddl" -database "mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" up
	$(MAKE) gen-xo

.PHONY: gen-xo
gen-xo: __init-db-args
	go install github.com/xo/xo@42b11c7999bc6ac5be620949723f44bd0ec63e02
	xo schema --out "gen/dbschema" -t json 'mysql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)'

.PHONY: gen-daocore
gen-daocore:
	go run ./script/dbgen ./gen/dbschema/xo.xo.json

.PHONY: gen-from-oas
gen-from-oas:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4
	oapi-codegen -package oapi -generate server -o ./gen/oapi/oapi_server.gen.go ../documents/api/openapi.yaml
	oapi-codegen -package oapi -generate types -o ./gen/oapi/oapi_type.gen.go ../documents/api/openapi.yaml
	oapi-codegen -package oapi -generate spec -o ./gen/oapi/oapi_spec.gen.go ../documents/api/openapi.yaml

.PHONY: go-fix-lint
go-fix-lint:
	go install golang.org/x/tools/cmd/goimports@latest
	find . -print | grep --regex '.*\.go$$' | xargs goimports -w -local "github.com/youcan-jpn/dab"

.PHONY: seed-test-data
seed-test-data: __init-db-args
	for file in $$(find testing/data/db -type f -name '*.sql'); do mysql -u$(DB_USER) -p$(DB_PASS) -h$(DB_HOST) -P$(DB_PORT) --protocol='tcp' --database=$(DB_NAME) < $$file; done

.PHONY: go-test
go-test:
	$(MAKE) seed-test-data
	go test -race ./... -cover -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
