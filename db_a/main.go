package main

//create table lend_book(
//	book_id  int primary key auto_increment,
//	is_login bool NOT NULL ,
//	is_register bool NOT NULL ,
//	book_name varchar(50) NOT NULL ,
//	is_lend bool NOT NULL,
//	lend_name varchar(20),
//	backTime datetime
//)
//
//b:
//	预处理是提前将命令部分发给数据库进行提前处理，可以避免数据过大带来速度的降低，提升性能，优化速度
//	还可以比避免sql注入 因为提前处理好了命令部分 数据部分用占位符代替
