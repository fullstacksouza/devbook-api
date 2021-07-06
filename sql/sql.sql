DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id varchar(36) primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(50) not null,
 craeted_at timestamp default current_timestamp,
   updated_at timestamp default current_timestamp,
);

CREATE TABLE followers(
    user_id varchar(36) not null,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,
     craeted_at timestamp default current_timestamp,
   updated_at timestamp default current_timestamp,

    follower_id varchar(36) not null,
    FOREIGN KEY (follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE
    
    primary key(user_id,follower_id)

)

CREATE TABLE posts(
    id varchar(36) primary key,
    title varchar(50) not null,
    content varchar(300) not null,
    author_id varchar(36) not null,
   craeted_at timestamp default current_timestamp,
   updated_at timestamp default current_timestamp,

    FOREIGN KEY (author_id)
    REFERENCES users(id)
    ON DELETE CASCADE
)