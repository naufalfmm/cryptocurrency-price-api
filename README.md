# Cryptocurrency Price API

## Getting started

Service for track the latest price of cryptocurrency asset

## Prerequisities
1. Docker
2. Golang >=1.19

## How To Run
1. Create `.env` file
    ```sh
    NAME="Cryptocurrency Price API" # The service name
    PORT=PORT # The service port. The default is 8080

    DB_PATH=DB_PATH # The location of Sqlite database

    DB_MAX_IDLE_CONNECTION=DB_MAX_IDLE_CONNECTION # The max idle connection of the db. The default is 10
    DB_MAX_OPEN_CONNECTION=DB_MAX_OPEN_CONNECTION # The max open connection of the db. The default is 10
    DB_CONNECTION_MAX_LIFE_TIME=DB_CONNECTION_MAX_LIFE_TIME # The connection maximum lifetime. The default is 60s

    DB_LOG_MODE=DB_LOG_MODE # Log mode of the DB. If true, the query is printed. The default is false
    DB_LOG_SLOW_THRESHOLD=DB_LOG_SLOW_THRESHOLD # Threshold of the query printed as slow. The default is 1s

    DB_RETRY=DB_RETRY # The number of connection retrying. The default is 3
    DB_WAIT_SLEEP=DB_WAIT_SLEEP # The sleep time between retry. The default is 1s

    LOG_MODE=LOG_MODE # The log mode. If true, log is printed

    BCRYPT_COST=BCRYPT_COST # The password bcrypt cost. The default is 5

    JWT_PUBLIC_KEY=JWT_PUBLIC_KEY
    JWT_ALG=JWT_ALG # The default is HS256
    JWT_EXPIRES=JWT_EXPIRES # The default is 1h

    COINCAP_BASE_PATH=COINCAP_BASE_PATH # The base path of coincap API
    COINCAP_PRICE_SYNC_MODE=COINCAP_PRICE_SYNC_MODE # The mode of asset price sync via websocket. The default is false
    ```
2. Migrate the database by
    ```sh
    make db DOCKER= SQLPATH= DBPATH= DBFILENAME= ENVFILENAME=

    # Note:
    # DOCKER= // If the migrating process use docker, please set as true.
    # SQLPATH= // Set the value if the DOCKER parameter is true
    # DBPATH= // Set the value if the DOCKER parameter is true
    # DBFILENAME= // Set the value if the DOCKER parameter is true
    # ENVFILENAME= // Set the value if the DOCKER parameter is true
    ```
    Note:
    The command `make db` automatically create the docker image (`make db_init`) and run the migration (`make db_migrate`) if the `DOCKER` parameter is `true`

4. Run the service by
    ```sh
    make app DOCKER= PORT= DBPATH= DBFILENAME= ENVFILENAME=

    # Note:
    # DOCKER= // If the migrating process use docker, please set as true.
    # PORT= // Port the service is running. Please set is as stated in the .env file.
    # DBPATH= // Set the value if the DOCKER parameter is true
    # DBFILENAME= // Set the value if the DOCKER parameter is true
    # ENVFILENAME= // Set the value if the DOCKER parameter is true
    ```
    Note:
    The command `make app` automatically create the docker image (`make app_init`) and run the migration (`make app_run`) if the `DOCKER` parameter is `true`

## Testing

If you want to test the code, run
```sh
go test ./... -count=1 -failfast
```

## Make Command

### Initialize migration
The command create the docker image of migration `naufalfmm/cryptocurrency-price-api-migration`

```sh
make db_init DOCKER=

# Note
# DOCKER= // If the parameter is `true`, it will build the docker image
```

### Create SQL file
The command create the sql file

```sh
make db_create DOCKER= SQLPATH= DBPATH= DBFILENAME= ENVFILENAME=

# Note:
# DOCKER= // If the file migration creation use docker, please set as true.
# SQLPATH= // Set the value if the DOCKER parameter is true
# DBPATH= // Set the value if the DOCKER parameter is true
# DBFILENAME= // Set the value if the DOCKER parameter is true
# ENVFILENAME= // Set the value if the DOCKER parameter is true
```

### Migrate SQL files
The command migrates all sql files in `./migrates/sql` or `./sql`

```sh
make db_migrate DOCKER= SQLPATH= DBPATH= DBFILENAME= ENVFILENAME=

# Note:
# DOCKER= // If the migrating process use docker, please set as true.
# SQLPATH= // Set the value if the DOCKER parameter is true
# DBPATH= // Set the value if the DOCKER parameter is true
# DBFILENAME= // Set the value if the DOCKER parameter is true
# ENVFILENAME= // Set the value if the DOCKER parameter is true
```

### Rollback SQL files
The command rollbacks all sql files in `./migrates/sql` or `./sql`

```sh
make db_migrate DOCKER= SQLPATH= DBPATH= DBFILENAME= ENVFILENAME= VERSION=

# Note:
# DOCKER= // If the migrating process use docker, please set as true.
# SQLPATH= // Set the value if the DOCKER parameter is true
# DBPATH= // Set the value if the DOCKER parameter is true
# DBFILENAME= // Set the value if the DOCKER parameter is true
# ENVFILENAME= // Set the value if the DOCKER parameter is true
# VERSION= // The value is optional. If the VERSION parameter set, the sql files will be rollback-ed until the VERSION value
```

## API Docs

If the JWT inputted in `Authorization` header is not valid, the endpoint return

```json
// Response (401)
{
    "ok": false,
    "message": "invalid token",
    "data": {
        "error": "invalid token"
    }
}
```

### Sign Up (POST /v1/auth/signup)
#### Description
Endpoint to sign up new user
#### Request
##### Body
```json
{
    "email": "", // user's email. Requred
    "password": "", // user's pasword. Required
    "password_confirmation": "" // Required. The password confirmation must be same as password
}
```
#### Response (200)
```json
{
    "ok": true,
    "message": "Success",
    "data": {
        "id": 0,
        "email": ""
    }
}
```

#### Response (403)
##### Email Has Been Used
```json
{
    "ok": false,
    "message": "email has been used",
    "data": {
        "error": "email has been used"
    }
}
```

#### Response (500)
```json
{
    "ok": false,
    "message": "internal server error",
    "data": {
        "error": "internal server error"
    }
}
```

### Sign In (POST /v1/auth/signin)
#### Description
Endpoint to sign in
#### Request
##### Body
```json
{
    "email": "", // user's email. Requred
    "password": "", // user's pasword. Required
}
```
#### Response (200)
```json
{
    "ok": true,
    "message": "Success",
    "data": {
        "jwt": "xxxxx.yyyyy.zzzzz", // the JWT of user
        "user": {
            "id": 0,
            "email": ""
        }
    }
}
```

#### Response (403)
##### Wrong Password
```json
{
    "ok": false,
    "message": "wrong password",
    "data": {
        "error": "wrong password"
    }
}
```
##### Email Missing
```json
{
    "ok": false,
    "message": "email missing",
    "data": {
        "error": "email missing"
    }
}
```

#### Response (500)
```json
{
    "ok": false,
    "message": "internal server error",
    "data": {
        "error": "internal server error"
    }
}
```

### Track Coin Asset (POST /user-coins/track)
#### Description
Endpoint for add the coin to the user's tracker. Coin cannot be added to the tracker more than once.
#### Request
##### Headers
1. Authorization: `Bearer <JWT>`
##### Body
```json
{
    "coin_symbol": "ETH" // the coin symbol of coincap (coincap.io)
}
```
#### Response (200)
```json
{
    "ok": true,
    "message": "Success",
    "data": {
        "track_id": 0, // The track id of the user coin
        "code": "ETH", // The coincap symbol of the coin
        "name": "Ethereum", // The name of the coin
        "latest_price": 2919.4101049055726, // The latest price of the coin
        "latest_price_currency": "USD", // The currency of coin latest price
        "added_at": "2024-05-12T10:14:26.5455845+07:00", // The time of coin added to the tracker
        "added_by": ""
    }
}
```

#### Response (400)
##### Coin Has Been Added for The User
```json
{
    "ok": false,
    "message": "coin has been added for the user",
    "data": {
        "error": "coin has been added for the user"
    }
}
```

#### Response (500)
```json
{
    "ok": false,
    "message": "internal server error",
    "data": {
        "error": "internal server error"
    }
}
```

### Untrack Coin Asset (POST /user-coins/untrack)
#### Description
Endpoint for remove the coin to the user's tracker
#### Request
##### Headers
1. Authorization: `Bearer <JWT>`
##### Body
```json
{
    "coin_symbol": "ETH" // the coin symbol of coincap (coincap.io)
}
```
#### Response (200)
```json
{
    "ok": true,
    "message": "Success"
}
```

#### Response (400)
##### Coin Missing
```json
{
    "ok": false,
    "message": "coin is missing",
    "data": {
        "error": "coin is missing"
    }

}
```
##### Tracking Coin is Missing (If we untrack unadded coin)
```json
{
    "ok": false,
    "message": "tracking coin is missing",
    "data": {
        "error": "tracking coin is missing"
    }

}
```

#### Response (500)
```json
{
    "ok": false,
    "message": "internal server error",
    "data": {
        "error": "internal server error"
    }
}
```

### Get All Tracking Coins (GET /user-coins)
#### Request
##### Headers
1. Authorization: `Bearer <JWT>`
2. X-Currency: // Example: IDR, USD, CLP

#### Response (200)
```json
{
    "ok": true,
    "message": "Success",
    "data": {
        "items": [
            {
                "track_id": 12,
                "code": "ETH",
                "name": "Ethereum",
                "latest_price": 2931.13,
                "latest_price_currency": "USD",
                "added_at": "2024-05-12T10:14:26.5455845+07:00",
                "added_by": ""
            }
        ]
    }
}
```

#### Response (500)
```json
{
    "ok": false,
    "message": "internal server error",
    "data": {
        "error": "internal server error"
    }
}
```