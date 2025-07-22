build:
	docker build -t budgit-back .

run-with-docker: build
	docker run -p 8080:8080 budgit-back

clean:
	docker rmi budgit-back

.PHONY: run-with-docker clean build
