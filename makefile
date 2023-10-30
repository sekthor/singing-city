all: frontend backend

frontend: ui/Dockerfile
	docker build ./ui -t sbsc-frontend

backend: Dockerfile
	docker build . -t sbsc-backend
