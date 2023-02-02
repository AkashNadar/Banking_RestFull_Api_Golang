create database banking;
use banking;

create table `customers`(
	`customer_id` int(11) not null auto_increment,
	`name` varchar(100) not null,
	`date_of_birth` date not null,
	`city` varchar(100) not null,
	`pincode` varchar(10) not null,
	`status` tinyint(1) not null default '1',
	primary key (`customer_id`)
) engine=InnoDB auto_increment=2006 default CHARSET=latin1;


insert into `customers` values
(2000, 'Steve', '1978-12-15', 'Delhi', '110075', 1);


create table `accounts`(
	`account_id` int(11) not null auto_increment,
	`customer_id` int(11) not null,
	`opening_date` datetime not null default current_timestamp,
	`account_type` varchar(10) not null,
	`pin` varchar(10) not null,
	`status` tinyint(4) not null default '1',
	primary key (`account_id`),
	key `accounts_FK` (`customer_id`),
	constraint `accounts_FK` foreign key (`customer_id`) references `customers` (`customer_id`)
) engine=InnoDB auto_increment=95476 default CHARSET=latin1;

insert into `accounts` values 
	(95470, 2000, '2020-08-22 10:20:06', 'Saving', '1075', 1);
	
select * from accounts;

create table `transactions`(
	`transaction_id` int(11) not null auto_increment,
	`account_id` int(11) not null,
	`amount` int(11) not null,
	`transaction_type` varchar(10) not null,
	`transaction_date` datetime not null default current_timestamp,
	primary key (`transaction_id`),
	key `transactions_FK` (`account_id`),
	constraint `transactions_FK` foreign key (`account_id`) references `accounts` (`account_id`)
) engine=InnoDB default CHARSET=latin1;