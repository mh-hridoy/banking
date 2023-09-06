mup:
	migrate  -path db/migrations -database "postgresql://hridoy:2543@localhost:5432/go1?sslmode=disable" -verbose up

mdown:
	migrate  -path db/migrations -database "postgresql://hridoy:2543@localhost:5432/go1?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

generate:
	sqlc generate

run:
	go run main.go

.PHONY:
	mup mdown test generate
