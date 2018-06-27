echo "select sum(total_crawl),sum(total_fail),1-sum(total_fail)/sum(total_crawl) from proxy;" | mysql -h127.0.0.1 -uroot -phe110 proxy
