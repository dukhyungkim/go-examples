build:
	go clean
	go build -trimpath -ldflags='-s -w'

sql:
	sqlc generate

graphql:
	gqlgen generate
