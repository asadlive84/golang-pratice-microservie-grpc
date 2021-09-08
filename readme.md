### how to run this project?
First open your terminal and run this commmand in your terminal.

`go mod tidy`

Now we run our `server`

`go run main.go`

Now, Open another `terminal` tab and make sure our `server` is run.

How to send `GET` method with `CURL`

`curl -v localhost:9000 -XGET  | jq `

Results:

```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying 127.0.0.1:9000...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 9000 (#0)
> GET / HTTP/1.1
> Host: localhost:9000
> User-Agent: curl/7.68.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Wed, 08 Sep 2021 17:50:57 GMT
< Content-Length: 88
< Content-Type: text/plain; charset=utf-8
< 
{ [88 bytes data]
100    88  100    88    0     0  44000      0 --:--:-- --:--:-- --:--:-- 44000
* Connection #0 to host localhost left intact
[
  {
    "id": 1,
    "name": "Samsung",
    "price": "$200"
  },
  {
    "id": 2,
    "name": "Nokia",
    "price": "$900"
  }
]

```



How to POST `GET` method with `CURL`

`curl -v localhost:9000 -d '{"name":"vivo","price":"95"}' -XPOST  | jq `


Now if you hit `curl -v localhost:9000 -XGET  | jq ` you will see all products in your terminal.

