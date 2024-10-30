# Car Rent

## Client-Server application for making car rents

### Stack

- Backend: Go
  - Web: Gin
  - DB: Postgres
  - ORM: Gorm
- Frontend: JS
  - Render: React
  - Components: MUI + MUI plugins

### To run

1) Change secrets in `secrets/` folder
2) Execute

```shell
docker compose --profile prod up -d --build 
```

If you want some test data to be generated, add `--profile test` to command
