.PHONY: run ncommit temp get 

run:
	cls
	go run .

ncommit:
	git add .
	git commit -m "ncommit"
	git git push origin main

temp: 
	go get github.com/alediaferia/gocountries

get: 
	cls
	go get github.com/nelsonsaake/go-ns@v0.0.9