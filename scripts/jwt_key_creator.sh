echo "create key pair for jwt in \"dev-keys\" under project folder"

dir_path="../dev-keys"

if [ ! -L "$dirname" ]
then
    mkdir $dir_path
fi

ssh-keygen -t rsa -b 4096 -m PEM -f $dir_path/jwt_RS256.key -q -N ""

openssl rsa -in $dir_path/jwt_RS256.key -pubout -outform PEM -out $dir_path/jwt_RS256.key.pub