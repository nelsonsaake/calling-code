.PHONY: run ncommit temp get pull

run:
	cls
	go run .

ncommit:
	go mod tidy
	git add .
	git commit -m "ncommit"
	git push origin main

temp: 
	go get github.com/alediaferia/gocountries

get: 
	cls
	go get github.com/nelsonsaake/go-ns@v0.0.12
	go get github.com/nelsonsaake/go-ns@latest

pull:
	git pull origin main