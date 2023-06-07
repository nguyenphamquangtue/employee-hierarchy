#!bin/bash
export ENV=develop
export SERVER=employeeHierarchy
export SECRET_KEY=secret
export PORT=:3000
export POSTGRES_HOST=localhost
export POSTGRES_PORT=5432
export POSTGRES_USER=fram
export POSTGRES_PASSWORD=fram
export POSTGRES_DB=employeeHierarchy

run:
	go run main.go