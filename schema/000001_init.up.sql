CREATE TABLE users
(
    id serial not null unique,
    usertype integer not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);