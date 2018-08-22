# UpWorkDemo

## Running the App

To run the app, make sure you have **go** and **go get** installed.

Auth0 credentials are needed as follows.

```bash
# .env

AUTH0_CLIENT_ID=YL6wNwYgl1kz2fnDoscdByFj8MoQnzwK
AUTH0_DOMAIN=programminggeek.auth0.com
AUTH0_CLIENT_SECRET=zUqlGi-2Yg13prnxEeaB0t1FOsskujJzbS1tJ-DcnUs7jMUXoamNsaRPfKVxZvNo
AUTH0_CALLBACK_URL=http://localhost:3000/callback
AUTH0_AUDIENCE=
```


Once you've set your Auth0 credentials in the `.env` file, run `go get -d` to install the Go dependencies.

Run `go run main.go` to start the app and navigate to [http://localhost:3000/](http://localhost:3000/).