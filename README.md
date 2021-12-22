# go-restful-server-ping
api for pinging a server using go-restful


## API v1

| URL | REST Verb | Action | Resource |
| --- | --- | --- | --- |
| `/ping` | GET | Ping the server | |
| `/v1/train` (details as JSON) | POST | Create  | Train |
| `/v1/station` (details as JSON) | POST | Create | Station |
| `/v1/train/id` | GET | Read | Train |
| `/v1/station/id` | GET | Read | Station |
| `/v1/schedule` (source and destination as JSON) | POST | Create | Route |


## Test

```
curl -X GET http://localhost:8000/ping

2021-12-22 21:34:52.243239244 +0100 CET m=+14.830391099
```

POST request to create a new train
```
curl -X POST \
    http://localhost:8000/v1/trains \
    -H 'cache-control: no-cache' \
    -H 'content-type: application/json' \
    -d '{"driverName":"Menaka", "operatingStatus":true}'

{
 "ID": 1,
 "DriverName": "Menaka",
 "OperatingStatus": true
}
```

GET request of the created train resource
```
curl -X GET http://localhost:8000/v1/trains/1

{
 "ID": 1,
 "DriverName": "Menaka",
 "OperatingStatus": true
}
```

DELETE the train resource 1
```
curl -X DELETE http://localhost:8000/v1/trains/1
```

If you send a GET request again, the train resource 1 is not found:
```
curl -X GET http://localhost:8000/v1/trains/1

Train could not be found
```