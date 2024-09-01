.PHONY: run stop
run:
	@echo "Building and running bomberman-dom..."
	docker build -t bomber-backend-image ./backend
	docker run -d -p 3000:3000 --name bomber-backend-container bomber-backend-image
	docker build -t bomber-frontend-image ./frontend
	docker run -d -p 8080:8080 --name bomber-frontend-container bomber-frontend-image
stop:
	@echo "Stopping and removing bomberman-dom..."
	docker stop bomber-backend-container bomber-frontend-container
	docker rm bomber-backend-container bomber-frontend-container
	docker rmi bomber-backend-image bomber-frontend-image