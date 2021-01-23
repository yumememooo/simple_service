# go_demo


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
```
// you can run diff. config
go run main.go -confDir configs --confEnv dev
```

