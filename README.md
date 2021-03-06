# clean-architecture-go

## Start server

```
docker-compose up -d
```

## Migration

Uses https://github.com/golang-migrate/migrate.

Binary is must to be downloaded to `bin` directory.  
Make sure mysql server is running.

```
$ cd bin
$ curl -SL https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
```

Then, run migration command.

```shell
$ make migration_up
```

