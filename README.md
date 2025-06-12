# Cognitive API Backend

This repository contains a simple Go backend exposing APIs for running various cognitive tests.

## Available Endpoints

- `/api/memory` – basic visual memory grid test
- `/api/test` – advanced grid based on cognitive stages
- `/api/test/submit` – save user submission
- `/api/test/results` – list of submissions received
- `/api/stroop` – Stroop word/color test
- `/api/math` – simple arithmetic questions
- `/api/sequence` – numeric series reasoning task
- `/api/iq` – approximate IQ estimation from test scores
- `/api/pool` – generate a randomized pool of tests for a session

All endpoints return JSON data and support CORS for local development.

Start the server with:

```bash
go run .
```

The server listens on `http://localhost:8080` by default.
