docker run -d \
	--name postgres \
	-e POSTGRES_PASSWORD=testpwd \
  -e POSTGRES_USER=test \
  -e POSTGRES_DB=test \
  -p 5432:5432 \
postgres