## AB Testing bucket.

### Install

#### Go

```
go get -u github.com/outman/abucket
```

#### Config

.abucket.yaml

```yaml
GIN_MODE: release
SERVER_LISTEN: ":8989"

DB_DRIVER: mysql
DB_CONNECTION: "root:admin@tcp(127.0.0.1:3306)/abucket?charset=utf8mb4&parseTime=True&loc=Local"
MYSQL_MAX_IDLE: 10
MYSQL_MAX_OPEN_CONNS: 100
MYSQL_LOG_MODE: false

CRYPT_KEY: e2e11d1b6d1ac295ffec01a4b1938fcb
LOG_PATH: "/tmp/abucket.log"

CORS_ALLOW_ORIGINS:
    - "http://localhost"

ADMIN:
  admin: testpassword


```

#### Database table

```sql
create database abucket default charset utf8mb4;
```

```
./abucket automigrate  // create tables.
```

#### Run

```
./abucket server
```

### Usage
```
./abucket

A simple AB Tesing traffic bucket system.

Usage:
  abucket [command]

Available Commands:
  automigrate Auto migrate database tables.
  help        Help about any command
  server      Run application.
  token       Make access token with .abucket.yaml's ADMIN node.
  version     Display version.

Flags:
      --config string   config file (default is $HOME/.abucket.yaml)
  -h, --help            help for abucket
  -t, --toggle          Help message for toggle

Use "abucket [command] --help" for more information about a command.
```