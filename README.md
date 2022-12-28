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
-   <b>Routing</b> - Gin Web Framework ----------------------------[📚](https://gin-gonic.com/docs) [:octocat:](https://github.com/gin-gonic/gin)

-   <b>CLI</b> - Cobra -------------------------------------------------[📚](https://cobra.dev) [:octocat:](https://github.com/spf13/cobra)

-   <b>Dependency Injection</b> - Fx --------------------------------[📚](https://uber-go.github.io/fx/get-started) [:octocat:](https://github.com/uber-go/fx)

-   <b>Environment</b> - Viper ---------------------------------------[:octocat:](https://github.com/spf13/viper)

-   <b>Logging</b> - Zap ----------------------------------------------[:octocat:](https://github.com/uber-go/zap)

-   <b>PostgreSQL</b> - GORM ---------------------------------------[📚](https://gorm.io/docs) [:octocat:](https://github.com/go-gorm/gorm)

-   <b>Redis</b> - Go-Redis -------------------------------------------[📚](https://redis.uptrace.dev/guide) [:octocat:](https://github.com/go-redis/redis)

-   <b>DB Viewer</b> - pgAdmin4 (Web) -----------------------------[🐳](https://hub.docker.com/r/dpage/pgadmin4)

-   <b>Authentication</b> - JWT (Access + refresh) ------------------[:octocat:](https://github.com/golang-jwt/jwt)
