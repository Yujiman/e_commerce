# REST API

### Pagination request
#### Arguments

| Option   | Type     | Required | Description                         |
|----------|----------|----------|-------------------------------------|
| `p`      | `string` | no       | page                                |
| `limit`  | `string` | no       | limit(if value = -1, then no limit) |
| `offset` | `string` | no       | offset                              |

## Available routes


## Policy

### Get client by iin

```
GET /v1/clients
```

#### Arguments

| Option | Type     | Required | Description |
|--------|----------|----------|-------------|
| `iin`  | `string` | yes      | uuid        |


#### Response
**Code: 200**

```json
{
    "id": "32941413",
    "first_name": "МАРАЛ",
    "last_name": "КУЛЖАБАЕВА",
    "middle_name": "",
    "document": false,
    "license": false,
    "class": "3"
}
```
_Error Response Codes_

`409` | `422` | `400` | `405` | `401` | `500` | `503`

