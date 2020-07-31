### AB Testing bucket.

#### WIP

### .abucket.yaml

```
GIN_MODE: release
SERVER_LISTEN: ":8989"

DB_DRIVER: mysql
DB_CONNECTION: "root:admin@tcp(127.0.0.1:3306)/abucket?charset=utf8mb4&parseTime=True&loc=Local"
MYSQL_MAX_IDLE: 10
MYSQL_MAX_OPEN_CONNS: 100
MYSQL_LOG_MODE: false

LOG_PATH: "/tmp/abucket.log"

CORS_ALLOW_ORIGINS:
    - "http://localhost"

```
