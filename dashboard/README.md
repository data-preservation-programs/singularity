# Singularity Dashboard

## Development
First, setup a test database using below command from **root**:
```
go run ./testdb
```
This will generate a database `test.db`
Then run the dashboard backend API by running below from the **root**:
```
make build
./singularity --database-connection-string "sqlite:test.db" run api
# Or 
# DATABASE_CONNECTION_STRING="sqlite:test.db" ./singularity run api
```

Then in another terminal, run the dashboard frontend server:
```
npm start
```

Now go to http://127.0.0.1:3000/ to view the dashboard. All API calls will be made to http://127.0.0.1:9090/api which is the API server

## Commit
To embed the built-in website into Singularity, make sure to run below steps from the **root** before committing:
```shell
npm run build
git add ./dashboard/build
```

All built js needs to be checked in to the repo, this makes sure the compiled JS is embedded into the golang binary. You can see how this works using below command:
```shell
make build
./singularity --database-connection-string "sqlite:test.db" run api
# Or 
# DATABASE_CONNECTION_STRING="sqlite:test.db" ./singularity run api
```
Goto http://127.0.0.1:9090 to verify the new change is reflected in the Singularity binary.

