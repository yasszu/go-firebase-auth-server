# Firebase Auth Server
## Getting Started
### Setup environment
1. Create a `.env` file
   ```sh
   cp .env.default .env
   ```
1. Download [service-account-file.json](https://firebase.google.com/docs/admin/setup)
1. Edit the `.env` file
    
### Run Server
1. Run containers
    ```
    docker-compose up
    ```

## API
### POST /authenticate
```sh
curl --location --request POST 'localhost:8888/authenticate ' \
--form 'id_token={your_id_token}'
```

### POST /v1/me
```sh
curl --location --request GET 'localhost:8888/v1/me' \
--header 'Authorization: Bearer {your_id_token}'
```
