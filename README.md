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