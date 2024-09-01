.PHONY: run stop

run:
	@echo "Building and running social network..."
	docker build -t social-backend-image ./backend
	docker run -d -p 3000:3000 --name social-backend-container social-backend-image

	docker build -t social-frontend-image ./frontend
	docker run -d -p 8080:8080 --name social-frontend-container social-frontend-image

stop:
	@echo "Stopping and removing social network..."
	docker stop social-backend-container social-frontend-container
	docker rm social-backend-container social-frontend-container
	docker rmi social-backend-image social-frontend-image
