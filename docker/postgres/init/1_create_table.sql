CREATE TABLE my_user (
    id bigserial primary key,
    name varchar(255) NOT NULL,
    email varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);