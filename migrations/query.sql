create table users(id int primary key auto_increment,
name  varchar(255)  , 
telegram_id int,    
first_name varchar(255), 
last_name varchar(255) , 
chat_id int,
created_at datetime,
updated_at datetime,
deleted_at datetime );