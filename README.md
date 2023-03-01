# GinProject
請參照以下步驟來進行啟動服務

## postgresql
* 透過 docker hub 將 postgresql 的 image pull 至本機
```
docker pull postgres
```

* 執行 script/psql_init.sh 以建立 postgresql 的 container
* 執行 script/psql_create_users_table.sh 以在 db 中建立 users table

## app server
* 執行 script/jwt_key_creator.sh 以建立所需的 key file
* 在 cmd 下建立 .env 檔案，內容如下
```
DB_CONFIG="host=localhost user=test password=testpwd dbname=test port=5432"
JWT_PRIVATE_KEY="../dev-keys/jwt_RS256.key"
JWT_PUBLIC_KEY="../dev-keys/jwt_RS256.key.pub"
```
* 至 cmd 下執行
```
go run .
```
此服務預設會在 8080 port 運行，亦可使用 localhost:8080/swagger/index.html 檢視 API 文件