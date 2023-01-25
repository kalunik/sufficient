## Run the service

1) `cp .env_sample .env`

2) `make`

## Send json

3) `cd sendJSON`

4) `go run publisher.go -json order2.json` you can use it without `-json` flag (default json: order1.json).

## Searching order data

5) Now you can go to [adminer](localhost:8080) to see Postgres records. Login and find table `orders`. 
Select all from it, copy `order_uid`.
6) And open [the page](localhost:8000) to search order's json by `order_uid` that you have.
