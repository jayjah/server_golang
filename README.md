### GoLang Server Example Project with GIN
    Simple GoLang Server to learn Golang for myself

#### Basic commands:

##### Build
```
    go build main.go
```

##### Run
```
    go run main.go
```

##### Run with Hot-Reload
```
    air
```

##### Create/Recreate db models in `./models` || `./gorm_models`
```
    sqlboiler psql -c sqlboiler.toml 
    OR 
    gentool -db postgres -dsn "host=localhost user=backend-db password=backend-db dbname=backend-db port=5432 sslmode=disable" -outPath "./gorm_models"
```


#### Todos:
- Docker [x]
- Docker-Compose [x]
- Migrations [x]
- Static File Handling [x]
- Data Seeding [x]
- Metrics (Prometheus) [x]
- CronJobs [x]
- Makefile
