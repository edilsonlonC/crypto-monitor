run:
	go run src/infrastructure/app/main.go
# Require local nodemon
develop:
	nodemon --exec go run src/infrastructure/app/main.go --signal SIGTERM
run-migration: 
	go run  src/resources/migrations/main.go
