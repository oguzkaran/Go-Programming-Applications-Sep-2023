-- chatapdb (PostgreSQL)

create table users (
    nickname varchar(100) primary key,
    username varchar(256) not null,
    password varchar(100) not null,
    register_date_time timestamp default (current_date) not null,
    -- ...
);

create table login_info (
    login_info_id bigserial primary key,
    nickname varchar(100) references users(nickname) not null,
    login_date_time timestamp default(current_date) not null
    -- ...
);

create table guest_info (
    guest_info_id bigserial primary key,
    nickname varchar(100) unique not null,
    password varchar(100) not null,
    login_date_time timestamp default(current_date) not null
    -- ...
);

-- ..