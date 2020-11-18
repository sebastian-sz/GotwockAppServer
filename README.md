# GotwockAppServer
Otwock App Server rewritten in Go.

### Example request:
(Assuming you started the server on port 9100)
## curl
curl \
 -H "Content-Type: application/json" \
 --request POST \
 --data '{"Latitude":52.0989711, "Longitude": 21.2715719, "maxDistanceFromUser": 5.1}' \
 http://127.0.0.1:9100/


## Python
```python
import requests
#import json

server = "http://127.0.0.1:9100"
headers = {"Content-Type": "application/json"}
payload = {"Latitude": 52.0989711, "Longitude": 21.2715719, "maxDistance": 5.1}

r = requests.post(server, json=payload, headers=headers)

print(r.text)
print(r.status_code)
```
