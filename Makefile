all :
	go build -o minisearch src/main.go

dev :
	go run src/main.go --dev
