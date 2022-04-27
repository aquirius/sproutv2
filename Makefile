clean:
	docker-compose -p sqlsprout-backend down
	docker-compose -p sqlsprout-backend up -d

run:
	SPRT_MYSQL="sprout:sprout@tcp(localhost:3311)/sprout?collation=utf8mb4_0900_ai_ci" \
	go run ./main.go

load-schemas:
	go run ./schemas/load.go
