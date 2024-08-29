create table users (
    id int unsigned not null auto_increment,
    email varchar(150) unique not null,
    password varchar(200) not null,

    primary key(id)
);
    
