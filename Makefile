all:run

run:
	@go run ./main.go $(args)

build:
	@go build -o ./app.exe ./main.go

app:
	@./app.exe

clean:
	@go clean -cache

.PHONY: all run build clean