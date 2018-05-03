mysql -h$1 -u$2 -p$3 <<EOF
    CREATE DATABASE IF NOT EXISTS proxy;
    use proxy;
    CREATE TABLE IF NOT EXISTS proxy (
      id int(11) NOT NULL auto_increment,
      ip_port varchar(50) NOT NULL,
      update_time timestamp NOT NULL default CURRENT_TIMESTAMP,
      last_fail_time timestamp NOT NULL default '1970-01-01 00:00:01',
      total_crawl int(11) NOT NULL default '0',
      total_fail int(11) NOT NULL default '0',
      PRIMARY KEY  (id),
      UNIQUE KEY ip_port (ip_port)
    ) ENGINE=InnoDB AUTO_INCREMENT=3389 DEFAULT CHARSET=utf8;
    exit
EOF
