# 3. Добавление организации (POST /ranges/add/organization)
curl -X POST "http://localhost:8080/ranges/add/organization" \
  -H "Content-Type: application/json" \
  -d '{
    "ip_start": "8.8.8.0",
    "ip_end": "8.8.8.255",
    "name": "Google"
  }'