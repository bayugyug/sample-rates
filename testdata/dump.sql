

create database rates;
create user rates;
grant all privileges on rates.* to rates@localhost identified by 'rat3s';
grant all privileges on rates.* to rates@127.0.0.1 identified by 'rat3s';
flush privileges;




CREATE TABLE IF NOT EXISTS `rates` (
	`id`          int(11) NOT NULL AUTO_INCREMENT,
	`base`        varchar(20) NOT NULL,
	`currency`    varchar(3)  NOT NULL,
	`rate`        decimal(16,8) default '0.0',
	`rate_dt`     date not null,
	`created_dt`  datetime,
	`modified_dt` datetime,
	PRIMARY KEY (`id`),
	UNIQUE  KEY idx_bz(`base`,`rate_dt`,`currency`)
) ENGINE=InnoDB;


drop table rates;




