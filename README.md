# prosigliere-rest-api


## How to execute the api locally
- Ensure Docker (and Docker Compose) is installed on your machine
- From the project root, run `make start`. This builds the images, starts the containers, and prepares the local 
environment so the API is ready to serve requests.

## cur examples:
```
curl --location --request GET 'http://localhost:3000/api/posts
```
```
curl --location --request POST 'http://localhost:3000/api/posts' \
--header 'Content-Type: application/json' \
--data-raw '{
"title": "mock title",
"content": "this is a mock comment"
}
```
```
curl --location --request GET 'http://localhost:3000/api/posts/post_id
```
```
curl --location --request POST 'http://localhost:3000/api/posts/b69dd84a-c9e5-4867-92fd-9179123f67c7/comments' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "someuser",
    "Content": "it is a mock comment"
}'
```

## Next steps
- Create unit tests.
- Implement a better error-handling layer that wraps errors (`%w`) and inspects them in higher layers with `errors.Is`.
- Return richer response bodies—for example, a JSON object with an `error` key for failure cases instead of a plain text message.
- Make comment and post pagination dynamic; avoid hard-coding `offset` and `limit` values as is done now.
- Set the config to be taken from a file; avoid hard-coding config like `url` for postgres client.