# link-shortener

Application to shorten link

## How to run?

You can use docker compose to run both database & server in isolated environment. Type `docker compose up -d` in your terminal from the root of the project.

## How to use it?

Send a `POST` request to `/`  with the following body

```json
{
  "url: "http://example.com"
}
```

You'll recevive the srot version of your desired URL.
