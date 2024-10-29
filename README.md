# ip_info_service

Service for searching information about ip in local database (country, organization).

## API Endpoints

### 1. Lookup IP Information

```
GET /lookup?ip=<ip_address>
```

```bash
curl "http://localhost:8080/lookup?ip=8.8.8.8"
```

### 2. Add Country Range

```
POST /ranges/add/country
```

```bash
curl -X POST "http://localhost:8080/ranges/add/country" \
-H "Content-Type: application/json" \
-d '{
"ip_start": "192.168.0.0",
"ip_end": "192.168.255.255",
"name": "RU"
}'
```

### 3. Add Organization Range

```
POST /ranges/add/organization
```

```bash
curl -X POST "http://localhost:8080/ranges/add/organization" \
-H "Content-Type: application/json" \
-d '{
"ip_start": "8.8.8.0",
"ip_end": "8.8.8.255",
"name": "Google"
}'
```