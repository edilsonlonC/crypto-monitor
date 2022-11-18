run:
	go run infrastructure/app/main.go
# Require local nodemon
develop:
	nodemon --exec go run infrastructure/app/main.go --signal SIGTERM
