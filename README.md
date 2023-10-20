# Bot detector

This is a bot detector project with frontend and backend.

## Prerequisites

Make sure you have Docker and Docker Compose installed on your machine.

- Docker installation guide: https://docs.docker.com/get-docker/
- Docker Compose installation guide: https://docs.docker.com/compose/install/

## Getting Started

1. Run the database:
```bash
docker-compose up clickhouse
```
2. Run the migrations:
```bash
make migrate
```
3. Run the backend project:
```bash
docker-compose up backend
```
4. Run the frontend project:
```bash
docker-compose up frontend
```

The backend project will be available at the 10658 and the frontend project at the port 5173. 

## Backend tests
To run the backend tests use:
```bash
make test
```
