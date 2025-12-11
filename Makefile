run:
	CompileDaemon -directory=./app -command="./app/app.exe"

migrate: 
	go run migration/main.go

swagger: 
	swag init -g app/main.go -o app/docs