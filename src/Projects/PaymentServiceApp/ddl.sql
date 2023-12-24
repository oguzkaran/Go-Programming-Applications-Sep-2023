create table users (
       username varchar(100) primary key,
       name varchar(300) not null,
       phone char(50) not null
    -- ...

);

create table products (
      code varchar(100) primary key,
      name varchar(300) not null,
      stock double precision default(0) not null,
      cost double precision not null,
      unit_price double precision not null
    -- ...
);

create table payments (
      payment_id bigserial primary key,
      username varchar(100) references users(username) not null,
      product_code varchar(100) references products(code) not null,
      amount double precision check(amount > 0) not null,
      unit_price double precision not null
    -- ...
);