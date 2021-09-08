### how to run this project?

`go mod tidy`

`go run main.go`

### How to send `GET` method with `CURL`

`curl -v localhost:9000 -XGET  | jq `

### How to POST `GET` method with `CURL`

`curl -v localhost:9000 -d '{"name":"vivo","price":"95"}' -XPOST  | jq `


Now if you hit `curl -v localhost:9000 -XGET  | jq ` you will see all products in your terminal.

