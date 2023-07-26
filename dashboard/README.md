# Singularity Dashboard

## Commit
To embed the built website into Singularity, make sure to run below steps from the **root** before committing:
```shell
npm run build
git add ./dashboard/build
```

## Development
You should first run the dashboard backend server by running below from the **root**:
```
make build
./singularity run api
```

This stands up the backend server on [http://localhost:9090](http://localhost:9090)

Then run the dashboard frontend server by running below from the **root**:
```
npm start
```
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.\
You will also see any lint errors in the console.
