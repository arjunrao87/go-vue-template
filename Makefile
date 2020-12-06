build-frontend:
	cd frontend && yarn run build

build-server:
	cd main && go build server.go

run: build-frontend build-server
	 ./main/server -static=frontend/