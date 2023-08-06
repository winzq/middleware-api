# Create middleware API

## Task
you should create a handler which sends how many days left until 1 Jan 2025 and response with HTTP 200 OK status code.
Secondly, build a middleware, which checks HTTP header User-Role presents and contains admin and prints
"Admin has logged at" and current date time with any layout to the console (using default log package or any 3rd party) if so.
## Test

```shell
curl --location '127.0.0.1:8080/status' \
--header 'User-Role: admin'
```

## Author

Semin Vadim