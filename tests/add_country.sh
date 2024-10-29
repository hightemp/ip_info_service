# 2. Добавление страны (POST /ranges/add/country)
curl -X POST "http://localhost:8080/ranges/add/country" \
  -H "Content-Type: application/json" \
  -d '{
    "ip_start": "192.168.0.0",
    "ip_end": "192.168.255.255",
    "name": "RU"
  }'