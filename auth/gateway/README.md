# Server Auth

Access Token lifetime - **1 Hour**.
Refresh Token lifetime - **24 Hours**.
All settings you can set in .env files

###### Available routes
### Sign In
```
POST /v1/oauth/login
```

#### Arguments

| Option            | Type     | Required | Description                                |
|-------------------|----------|----------|--------------------------------------------|
| `username`           | `string` | yes      | Login or Email or Phone. Length min 3.     |
| `password`        | `string` | yes      | Length min 6.                              |
| `domain`        | `string` | yes      | Length min 3.                              |

#### Response

    HTTP/1.1 200 OK
    Status: 200 OK
    Content-Type: application/json
    AccessToken: <JWT chank 1>
    X-Satrap-2: <JWT chank 2>

```
{
    "email": "",
    "id": "6494e273-85a2-494b-b1af-45ec9acd6144",
    "login": "admin",
    "phone": ""
}
```

#### Response

    HTTP/1.1 200 OK
    Status: 200 OK
    Content-Type: application/json


### Check Access Token (X-Satrap-1)

```
GET /v1/oauth/check
```

With Headers:

```
AccessToken: <JWT chank 1>
```

#### Response

    HTTP/1.1 200 OK
    Status: 200 OK
    Content-Type: application/json
    AccessToken: <JWT chank 1>

```
{
    "email": "",
    "id": "9ff09fc8-e937-4ae0-b90a-da0ec788f056",
    "login": "",
    "phone": "770511112233"
}
```

### Refresh Token

```
POST /v1/oauth/refresh
```

With Headers:

```
X-Satrap-2: <JWT chank 2>
```

#### Response

    HTTP/1.1 200 OK
    Status: 200 OK
    Content-Type: application/json
    AccessToken: <New JWT chank 1>
    X-Satrap-2: <New JWT chank 2>
