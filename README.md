# GotwockAppServer
Otwock App Server rewritten in Go!  
Otwock App Server was the idea of a simple backend server that given user coordinates
could display a list of nearby tourist or simply interesting locations.  
It probably would not top Google Maps in that way, but it was more of a programming exercise.  

The first version of this backend server was written in Python, but it was tempting to
use a different, new (to me) technology. Hence, came the idea of Gotwock App Server!

# Installation and running.
You need to have [Go installed](https://golang.org/doc/install).  

Then, you can simply clone this repository:  
`git clone https://github.com/sebastian-sz/GotwockAppServer.git`

build the package (this will also install dependencies)  
`make build` or `go build -o bin/GotwockAppServer` 

finally, run:  
`./bin/GotwockAppServer`

If all goes well you should see the message   
`Starting server at: http://127.0.0.1:9100`

# How to use?
The server provides locations via REST API. Here are example requests:  
(keep in mind that the server should be running on a separate terminal)
#### Curl
```bash
curl \
 -H "Content-Type: application/json" \
 --request POST \
 --data '{"Latitude":52.0989711, "Longitude": 21.2715719, "maxDistance": 5.1}' \
 http://127.0.0.1:9100/
```
#### Python
```python
import requests

server = "http://127.0.0.1:9100"
headers = {"Content-Type": "application/json"}
payload = {"Latitude": 52.0989711, "Longitude": 21.2715719, "maxDistance": 5.1}

r = requests.post(server, json=payload, headers=headers)

print(r.text)
print(r.status_code)

```

To get all available locations simply pass 0.0 for `maxDistance`, or don't pass it at all. `Latitude` and `Longitude` 
are mandatory.

# Running tests.

### Unit tests:
The package is tested with [testify](https://github.com/stretchr/testify/).
The test suite is automatically run with each PR to `main` or `develop` branch.  
In order to run the test manually run: `make test` or `go test ./...`.

### End to end tests:
This is a work in progress. As for now, you can manually hit the server with curl. 

# Backstory.
Yes, you can skip it if you want.  

Otwock is a small town in masovian voivodeship in Poland. Although being quite rough
around the edges there are quite a few gems, literally hidden in plain sight. Those kinds of gems that have a 
look, or a story that you cannot find in Google Maps.  

The idea was originally brought up by my friend [Jan](https://github.com/JanSzala).  After initial planning, we decided 
to split the work: I did backend server, he did the client.
This repository contains the server code for GotwockAppServer. 

