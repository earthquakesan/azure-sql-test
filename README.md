# azure-sql-test
Test util for azure SQL written in golang

## Usage

```
export SQL_USERNAME=sqladmin
export SQL_PASSWORD='password#$!@#'
export SQL_HOSTNAME=hostname.database.windows.net
export SQL_PORT=1433
export SQL_DATABASE=projectsqldatabase

make build
# For MacOS
./azure-sql-test-darwin-amd64
# For Windows
./azure-sql-test.exe
# For Linux
./azure-sql-test-linux-amd64
```

On the successful connection application will return exit code 0. 