docker run --name first -p 3306:3306 -v /Users/baidu/proxypool/mysql/data:/var/lib/mysql -v /Users/baidu/proxypool/mysql/mysql-files:/var/lib/mysql-files -e MYSQL\_ROOT\_PASSWORD=he110 -d mysql
