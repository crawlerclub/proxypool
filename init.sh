mysql -h127.0.0.1 -uroot -phe110 <<EOF
    CREATE DATABASE IF NOT EXISTS proxy;
    use proxy;
    CREATE TABLE IF NOT EXISTS proxy (
      id int(11) NOT NULL auto_increment,
      ip_port varchar(50) NOT NULL,
      update_time datetime NOT NULL default CURRENT_TIMESTAMP,
      last_fail_time datetime NOT NULL default '0000-00-00 00:00:00',
      total_crawl int(11) NOT NULL default '0',
      total_fail int(11) NOT NULL default '0',
      PRIMARY KEY  (id),
      UNIQUE KEY ip_port (ip_port)
    ) ENGINE=InnoDB AUTO_INCREMENT=3389 DEFAULT CHARSET=utf8;
    exit
EOF
