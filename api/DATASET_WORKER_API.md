# Dataset Worker API Documentation

This API allows you to manage the `dataset-worker` process from the backend via HTTP endpoints. You can start, stop, check status, view logs, and list jobs for the worker—all from the UI or other clients.

## Endpoints

### Start Worker
- **POST** `/api/dataset-worker/start`
- **Payload Example:**
```json
{
  "concurrency": 2,
  "enableScan": true,
  "enablePack": true,
  "enableDag": false,
  "exitOnComplete": false,
  "exitOnError": true,
  "minInterval": "5s",
  "maxInterval": "160s"
}
```
- **Response:**
```json
{ "status": "started", "pid": 12345 }
```

### Stop Worker
- **POST** `/api/dataset-worker/stop`
- **Response:**
```json
{ "status": "stopped" }
```

### Status
- **GET** `/api/dataset-worker/status`
- **Response:**
```json
{
  "running": true,
  "pid": 12345,
  "config": { ... },
  "startTime": "2025-07-18T12:00:00Z",
  "uptime": "2m30s"
}
```

### Logs
- **GET** `/api/dataset-worker/logs`
- **Response:**
```json
{
  "stdout": "Worker started...",
  "stderr": ""
}
```

### Jobs
- **GET** `/api/dataset-worker/jobs`
- **Response:**
```json
{
  "jobs": [
    { "id": 42, "type": "scan", "status": "processing", "started": "..." }
  ]
}
```

## Usage Notes
- Only one worker process can run at a time.
- All endpoints require admin privileges (add authentication as needed).
- The jobs endpoint lists recent jobs from the database.
- Logs are captured from the worker process's stdout/stderr.

## Error Codes
- `409 Conflict`: Worker already running
- `404 Not Found`: Worker not running (for stop/logs)
- `500 Internal Server Error`: Unexpected error

## Example Workflow
1. Start the worker with desired config.
2. Monitor status and logs.
3. Stop the worker when needed.
4. List jobs to track progress.

## Help Text
- Each endpoint returns a clear status or error message.
- Invalid config or requests will return a helpful error.

---
For more details, see the code in `api/dataset_worker.go` and tests in `api/dataset_worker_test.go` and `api/dataset_worker_integration_test.go`.
