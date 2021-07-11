# Firebase Auth Server
## Getting Started
### Setup environment
1. Create a `.env` file
   ```sh
   cp build/local/server/.env.default build/local/server/.env
   ```
1. Download [service-account-file.json](https://firebase.google.com/docs/admin/setup)
1. Edit the `.env` file
1. Add login page
   ```sh
   cp ./config.default.js ./public/javascripts/config.js
   ```
1. Set Firebase configuration at `public/javascripts/config.js`
    
### Run Server
1. Run containers
   ```
   make run
   ```
1. Run migration
   ```
   make migrate-up
   ```
1. Open http://localhost:8888/
## API
### POST /authenticate
```sh
curl --location --request POST 'localhost:8888/authenticate' \
--form 'id_token={your_id_token}'
```

### GET /v1/me
```sh
curl --location --request GET 'localhost:8888/v1/me' \
--header 'Authorization: Bearer {your_id_token}'
```
