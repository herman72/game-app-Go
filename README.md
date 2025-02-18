# game-app-Go
a simple quiz game develop in Go


# Migrations
```bash
go install github.com/rubenv/sql-migrate/...@latest
sql-migrate status -env="production" -config=dbconfig.yml 
sql-migrate upd -env="production" -config=dbconfig.yml 
sql-migrate down -env="production" -config=dbconfig.yml 
```