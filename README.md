# Idempotency-Gateway Diagrams, Instructions and Documentation

## Architecture Diagram

![Alt text](flowchart.png)

## Setup Instructions

## Prerequisites

- Go 1.26.2 or later installed
- Git installed

## Clone the repository

```bash
git clone https://github.com/amalitechglobaltraining/Idempotency-Gateway.git
cd Idempotency-Gateway
```

## Install dependencies

```bash
go mod tidy
```

## Run locally

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

## Test the API

Send a request with a unique `Idempotency-Key` header.

```bash
curl -X POST http://localhost:8080/process-payment \
  -H "Content-Type: application/json" \
  -H "Idempotency-Key: test-key-123" \
  -d '{"amount":100,"currency":"GHS"}'
```

## Expected behavior

- Requests must include the `Idempotency-Key` header.
- Endpoint: `POST /process-payment`
- First request with a new key processes the payment and stores the response.
- Repeated requests with the same key and identical body return the cached response.
- Repeated requests with the same key but different body return a `422 Unprocessable Entity` error.

## Notes

- This project uses an in-memory map for idempotency storage, so stored responses are lost when the server restarts.
- The API simulates processing with a 2-second delay on the first request.
