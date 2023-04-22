CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    name varchar(255),
    job varchar(255) DEFAULT 'No job',
    age int CHECK(age > 0),
    salary decimal
);
