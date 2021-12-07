postgres:
	docker run --name postgres-web -p 5432:5432 -e POSTGRES_USER=monitor -e POSTGRES_PASSWORD=monitor -d postgres

swagger:
	swagger generate spec -o ./internal/services/api/docs/swagger.yaml --scan-models

createdb:
	docker exec -it postgres-web createdb --username=monitor --owner=monitor website-monitor

dropdb:
	docker exec -it postgres-web dropdb --username=monitor website-monitor

migrateup:
	migrate -path  ./internal/infrastructure/data_imp/migrations -database "postgresql://monitor:monitor@localhost:5432/website-monitor?sslmode=disable" -verbose up

migratedown:
	migrate -path  ./internal/infrastructure/data_imp/migrations -database "postgresql://monitor:monitor@localhost:5432/website-monitor?sslmode=disable" -verbose down