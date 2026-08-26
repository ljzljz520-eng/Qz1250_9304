# Store 55 Account Desk

## Build
`CGO_ENABLED=0 go build ./...`

## Run
`STORE55_DATA=store55.db go run ./cmd/server`

## Test
`go test -count=1 ./...`
