# link-shortener

Application to shorten link

## How to run?

You can use docker compose to run both the database & server in an isolated environment. Type `docker compose up -d` in your terminal from the project's root.

## How to use it?

Send a `POST` request to `/`  with the following body

```json
{
  "url: "http://example.com"
}
```

You'll receive the short version of your desired URL.
