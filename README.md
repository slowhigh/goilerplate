<h1 align="center">
  <b>Goilerplate</b>
</h1>

<p align="center">
  For the perfect structure. 🔥🔥🔥
</p>

## Contents
- [Contents](#contents)
- [Quick start](#quick-start)
- [Default API](#default-api)
  -   [SignUp](#signup)
  -   [SignIn](#signin)
  -   [User Info](#user-info)
- [Implemented Features](#implemented-features)

<br>

## Quick start
- Make sure you have docker installed.
- Copy `/env/stage.env` to `.env`
- Run `docker-compose up -d`
  - API Host: localhost:5000
  - DB Viewer: localhost:8080 - Email: user@goilerplate.com - PW: 1234

<br>

## Default API
### SignUp
```HTTP
POST /auth/signup HTTP/1.1
Host: localhost:5000
Content-Type: application/json
Content-Length: 103

{
    "Email": "test@gmail.com",
    "Password": "1234",
    "Name": "test",
    "Role": "admin"
}
```

### SignIn
```HTTP
POST /auth/signin HTTP/1.1
Host: localhost:5000
Content-Type: application/json
Content-Length: 58

{
    "Email":"test@gmail.com",
    "Password":"1234"
}
```

### User Info
```HTTP
GET /user/info HTTP/1.1
Host: localhost:5000
Cookie: access_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...; refresh_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

<br>

## Implemented Features
-   <b>Routing</b> - Gin Web Framework [#Docs](https://gin-gonic.com/docs) [#GitHub](https://github.com/gin-gonic/gin)

-   <b>CLI</b> - Cobra [#Docs](https://cobra.dev) [#GitHub](https://github.com/spf13/cobra)

-   <b>Dependency Injection</b> - Fx [#Docs](https://uber-go.github.io/fx/get-started) [#GitHub](https://github.com/uber-go/fx)

-   <b>Environment</b> - Viper [#GitHub](https://github.com/spf13/viper)

-   <b>Logging</b> - Zap [#GitHub](https://github.com/uber-go/zap)

-   <b>PostgreSQL</b> - GORM [#Docs](https://gorm.io/docs) [#GitHub](https://github.com/go-gorm/gorm)

-   <b>Redis</b> - Go-Redis [#Docs](https://redis.uptrace.dev/guide) [#GitHub](https://github.com/go-redis/redis)

-   <b>DB Viewer</b> - pgAdmin4 (Web) [#DockerHub](https://hub.docker.com/r/dpage/pgadmin4)

-   <b>Authentication</b> - JWT (Access + refresh) [#GitHub](https://github.com/golang-jwt/jwt)
