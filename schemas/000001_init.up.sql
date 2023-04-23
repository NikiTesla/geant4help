CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    username varchar(255) NOT NULL,
    password_hash varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS users_info (
    id int UNIQUE NOT NULL,
    job varchar(255),
    age int CHECK(age > 0),
    salary decimal
);