DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id varchar(36) primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(50) not null,
    craeted_at timestamp default current_timestamp
);

CREATE TABLE followers(
    user_id string not null,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    follower_id string not null,
    FOREIGN KEY (follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE

    primary key(user_id,follower_id)

)