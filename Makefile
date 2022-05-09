clean:
	docker-compose -p sqlsprout-backend down
	docker-compose -p sqlsprout-backend up -d

run:
	go run ./main.go

load-schemas:
	go run ./schemas/load.go
