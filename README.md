# go_demo

- 提供簡單 API GET/POST 功能
- 具 config/log

## git at local initial setting

- 1.link project

```
git init
git remote add origin https://github.com/yumememooo/simple_service.git
```

- 2.check commit user

```
//check local user
git config user.name
git config user.email
//change local user
git config user.name "yumememooo"
git config user.email "yumememooo@gmil.com"
* 此為指定本專案，請勿加上--global全域設定
```

## start project

- you can run diff. config

```
go run main.go -confDir configs --confEnv dev
```

- http_server.StartHttpServer
  //start http & swagger
  http://127.0.0.1:56888/api/swagger/index.html

- swaggo Usage

```
 go get -u github.com/swaggo/swag/cmd/swag
 git init
 see more:https://github.com/swaggo/gin-swagger

 if you can't find swag tool,make sure your env setting is correct

      "terminal.integrated.env.windows": {
        "GOPATH":"D:\\goWorkSpace",
        "GOBIN":"${env:GOPATH}\\bin",
        "PATH": "${env:PATH};${env:GOPATH};D:\\goWorkSpace\\bin"
    },
    then open cmd, check  echo %PATH%
##

```

- find pet by animal_kind
  http://127.0.0.1:56888/api/v1/pet?animal_kind=%E8%B2%93
- post animal
  write to json
