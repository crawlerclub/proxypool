DIR="/data/zliu/proxypool/mysql"
docker run --name first -u 1000:1000 -p 3306:3306 -v $DIR/data:/var/lib/mysql -v $DIR/mysql-files:/var/lib/mysql-files -e MYSQL\_ROOT\_PASSWORD=he110 -d mysql:5.6
