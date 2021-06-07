DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id varchar(36) primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(50) not null,
    craeted_at timestamp default current_timestamp
);
