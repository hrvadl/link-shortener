# Simple link shortener

Application to shorten links.

## How to run?

Make sure you have [taskfile](https://taskfile.dev/) and [Go](https://go.dev/) installed.

1. Fill the `.env` file with variables stated in the `.env.example`
2. From the root of the repo run `task docker-run`

## How to use it?

Send a `POST` request to `/` with the following body:

```json
{
  "url: "http://example.com"
}
```

For example:

```sh
❯ curl -X POST -d '{"url": "http://example.com"}' -H "Content-Type: application/json" http://localhost:3000

http://localhost:3000/qbnwQzbOAYGgjndOARE7MQ==%

```

You'll receive the short version of your desired URL. Then you can paste it in your browser and application will redirect you
to the desired URL.
