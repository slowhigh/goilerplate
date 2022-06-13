# TypeGoMongo (Server)
### Languages
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

### Servers
![Nginx](https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white)

### Databases
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)

### Heroku로 배포하는 방법

#### "Profile" 파일의 내용
```
(ubuntu) web: bin/server
(window) web: bin\server.exe
```

#### go build
```
(ubuntu) go build -o bin/server -v .
(windows) go build -o bin/server.exe -v .
```