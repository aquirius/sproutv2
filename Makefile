clean:
	docker-compose -p sqlsprout-backend down
	docker-compose -p sqlsprout-backend up -d

run:
	SPRT_MYSQL="sprout:sprout@tcp(localhost:3311)/sprout?collation=utf8mb4_0900_ai_ci" \
	go run ./main.go

client:
	go run ./cmd/client -email=admin0@doe.de -password=test1234

cron:
	go run ./scripts/cron

load-schemas:
	go run ./schemas/load.go

create-accounts:
	go run ./scripts/dev-accounts

create-items:
	go run ./scripts/dev-create-items

create-timezones:
	go run ./scripts/timezones > ./internal/systems/static/timezones.go

create-countries:
	cd ./scripts/countries/ && go build && go run ./countries.go > ../../internal/systems/static/countries.go

generate:
	go install ./cmd/fbapigen
	go generate ./...

generateAPI:
	go run ./cmd/fbapigen/cli/cli.go -ts -dart -config=".generate.json"

generateDart:
	go run ./cmd/fbapigen/cli/cli.go -dart -config=".generate.json"

generateTS:
	go run ./cmd/fbapigen/cli/cli.go -ts -config=".generate.json"

test:
	go test -timeout 20m -p 1 -v -failfast -coverpkg ./internal/... -coverprofile profile.cov ./internal/testing/tests -args integration

paypal-totp:
	cd ./scripts/totp && go install ./
	totp -secret="YQZS RYCR 5ETQ DGJJ"