# Firebase Auth Server
## Getting Started
### Setup environment
1. Create a `.env` file
   ```sh
   cp .env.default .env
   ```
1. Edit the `.env` file
    
### Run Server
1. Run containers
    ```
    $ docker-compose up
    ```

## API
### POST /signup
```sh
curl --location --request POST 'localhost:8888/signup' \
--form 'email=test@sample.com' \
--form 'password=test12345' \
--form 'username=user1'
```

### POST /login
```sh
curl --location --request POST 'localhost:8888/login' \
--form 'email=test@sample.com' \
--form 'password=test12345'

```

### POST /v1/me
```sh
curl --location --request GET 'localhost:8888/v1/me' \
--header 'Authorization: Bearer {your_jwt_token}'
```
