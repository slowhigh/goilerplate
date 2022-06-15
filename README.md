# TypeGoMongo (Server)
## Languages
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

## Servers
![Nginx](https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white)

## Databases
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)

<br>
<br>

# Heroku Deploy

## Local Deployment
#### "Profile" File
```
(ubuntu) web: bin/{Project Name}           ex) web: bin/TypeGoMongo-Server
(window) web: bin\{Project Name}.exe       ex) web: bin\TypeGoMongo-Server.exe
```
#### Build
```
(ubuntu) go build -o bin/TypeGoMongo-Server -v .
(windows) go build -o bin/TypeGoMongo-Server.exe -v .
```
#### Deploy
```
(ubuntu) ~\TypeGoMongo-Server$ heroku local
(windows) PS ~\TypeGoMongo-Server> heroku local
```

<br>

## Remote Deployment
#### "Profile" File
```
(ubuntu) web: bin/{project name}           ex) web: bin/TypeGoMongo-Server
(window) web: bin\{project name}.exe       ex) web: bin\TypeGoMongo-Server.exe
```

#### Deploy
```
(ubuntu)
...

(windows)
PS ~\TypeGoMongo-Server> git branch master
PS ~\TypeGoMongo-Server> git switch master
PS ~\TypeGoMongo-Server> git add .
PS ~\TypeGoMongo-Server> git commit -m "deploy"            => The message may be 'ver 1.0.0.0' or others.
PS ~\TypeGoMongo-Server> git push heroku master             => The heroku is name of remote.
```

#### Destroy
```
heroku destroy --confirm {App Name}         ex) heroku destroy --confirm type-go-mongo
```