# API Documentation

Base URL: `http://localhost:8000/api/v1`

## Tables

### Create Table

Creates a new table with specified columns.

**Endpoint:** `POST /tables/create-table`

**Request Body:**
```json
{
  "name": "users",
  "columns": [
    { "name": "id", "type": 0, "is_indexed": true },
    { "name": "name", "type": 1, "length": 50, "is_indexed": false },
    { "name": "age", "type": 0, "is_indexed": false }
  ]
}
```

**Column Types:**
| Type | Value | Description |
|------|-------|-------------|
| INT | 0 | Integer |
| VARCHAR | 1 | String (requires `length`) |
| DATE | 2 | Date (YYYY-MM-DD) |
| TIMESTAMP | 3 | Timestamp (RFC3339) |
| FLOAT | 4 | Floating point |
| JSON | 5 | JSON object |

**Example:**
```bash
curl -X POST http://localhost:8000/api/v1/tables/create-table \
  -H "Content-Type: application/json" \
  -d '{
    "name": "users",
    "columns": [
      { "name": "id", "type": 0, "is_indexed": true },
      { "name": "name", "type": 1, "length": 50, "is_indexed": false }
    ]
  }'
```

**Response:**
```json
{
  "status": "CREATED",
  "message": "Table created successfully!"
}
```

---

## Records

### Insert Record

Inserts a new record into a table.

**Endpoint:** `POST /records/insert`

**Request Body:**
```json
{
  "name": "users",
  "values": {
    "id": 1,
    "name": "John"
  }
}
```

**Example:**
```bash
curl -X POST http://localhost:8000/api/v1/records/insert \
  -H "Content-Type: application/json" \
  -d '{
    "name": "users",
    "values": { "id": 1, "name": "John" }
  }'
```

**Response:**
```json
{
  "status": "CREATED",
  "message": "Record inserted"
}
```

---

### Query Records

Retrieves records from a table with optional filtering and column selection.

**Endpoint:** `POST /records/query`

**Request Body:**
```json
{
  "name": "users",
  "select": ["id", "name"],
  "filter": [
    { "column": "id", "operator": "=", "value": 1 }
  ]
}
```

**Filter Operators:**
| Operator | Description |
|----------|-------------|
| `=` | Equal |
| `!=` | Not equal |

**Example — Get all records:**
```bash
curl -X POST http://localhost:8000/api/v1/records/query \
  -H "Content-Type: application/json" \
  -d '{ "name": "users", "select": [], "filter": [] }'
```

**Example — With filter:**
```bash
curl -X POST http://localhost:8000/api/v1/records/query \
  -H "Content-Type: application/json" \
  -d '{
    "name": "users",
    "select": ["id", "name"],
    "filter": [{ "column": "id", "operator": "=", "value": 1 }]
  }'
```

**Response:**
```json
{
  "status": "OK",
  "data": [
    { "id": 1, "name": "John" }
  ]
}
```

---

### Update Records

Updates records matching the filter criteria.

**Endpoint:** `POST /records/update`

**Request Body:**
```json
{
  "name": "users",
  "values": {
    "name": "Jane"
  },
  "filter": [
    { "column": "id", "operator": "=", "value": 1 }
  ]
}
```

**Example:**
```bash
curl -X POST http://localhost:8000/api/v1/records/update \
  -H "Content-Type: application/json" \
  -d '{
    "name": "users",
    "values": { "name": "Jane" },
    "filter": [{ "column": "id", "operator": "=", "value": 1 }]
  }'
```

**Response:**
```json
{
  "status": "OK",
  "message": "1 row(s) updated"
}
```

---

### Delete Records

Deletes records matching the filter criteria.

**Endpoint:** `POST /records/delete`

**Request Body:**
```json
{
  "name": "users",
  "filter": [
    { "column": "id", "operator": "=", "value": 1 }
  ]
}
```

**Example:**
```bash
curl -X POST http://localhost:8000/api/v1/records/delete \
  -H "Content-Type: application/json" \
  -d '{
    "name": "users",
    "filter": [{ "column": "id", "operator": "=", "value": 1 }]
  }'
```

**Response:**
```json
{
  "status": "OK",
  "message": "1 row(s) deleted"
}
```

---

## Error Responses

All endpoints return errors in this format:

```json
{
  "status": "BAD_REQUEST",
  "message": "error description"
}
```

**Status Codes:**
| Status | Description |
|--------|-------------|
| OK | Success |
| CREATED | Resource created |
| BAD_REQUEST | Invalid request |
| NOT_FOUND | Resource not found |
| INTERNAL_SERVER_ERROR | Server error |
