# URL Shortener API

This is a simple REST API for shortening URL. It allows client to shorten the given URL.

## Endpoints Overview

| Method | Endpoint       | Description                                |
| ------ | -------------- | ------------------------------------------ |
| GET    | `/api/health`  | Check readiness of API                     |
| POST   | `/api/shorten` | Shorten the long URL provided by user      |
| GET    | `/{alias}`     | Redirect the shortened URL to the long URL |

## Endpoints

### API Health

<details>
<summary>Check readiness of API</summary>
Request:

- Method: `GET`
- URL: `/api/health`
- On Success Status Code: `200`

- Response:

  - `200 OK`

</details>

### Shorten URL

<details>
<summary>Shorten a long URL provided by user</summary>
Request:

- Method: `POST`
- URL: `/api/shorten`
- Headers:
  - `Content-Type`: `application/json`
- Request Body:
  - custom_alias `string` `unique`: Alias which will be used for redirection.
  - original_url `string`: URL shortening target.
- On Success Status Code: `201`
- Request Example:

  ```json
  {
    "custom_alias": "lol",
    "original_url": "https://leagueoflegends.com"
  }
  ```

- Response Example:

  ```json
  {
    "generated_link": "https://url-shortener-836536565354.asia-southeast2.run.app/lol"
  }
  ```

</details>

### Redirect URL

<details>
<summary>Redirect the shortened URL to the long URL</summary>
Request:

- Method: `GET`
- URL: `/{alias}`
- On Success Status Code: `200`

- Response:

  - `200 OK`

</details>
