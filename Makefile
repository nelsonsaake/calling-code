.PHONY: run ncommit

run:
	go run .

ncommit:
	git add .
	git commit -m "ncommit"
	git git push origin main