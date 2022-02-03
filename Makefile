api-build:
	go build ./cmd/api/main.go
	
api-run:
	go run ./cmd/api/main.go

postgres:
	docker run --name postgres-web -p 5432:5432 -e POSTGRES_USER=monitor -e POSTGRES_PASSWORD=monitor -d postgres

swagger:
	swagger generate spec -o ./internal/services/api/docs/swagger.yaml --scan-models

create-db:
	docker exec -it postgres-web createdb --username=monitor --owner=monitor website-monitor

drop-db:
	docker exec -it postgres-web dropdb --username=monitor website-monitor

migrate-up:
	migrate -path  ./internal/infrastructure/data_imp/migrations -database "postgresql://monitor:monitor@localhost:5432/website-monitor?sslmode=disable" -verbose up

migrate-down:
	migrate -path  ./internal/infrastructure/data_imp/migrations -database "postgresql://monitor:monitor@localhost:5432/website-monitor?sslmode=disable" -verbose down

