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