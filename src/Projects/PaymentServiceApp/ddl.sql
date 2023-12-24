create table users (
       username varchar(100) primary key,
       password varchar(100) not null,
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

create or replace function get_all_users()
returns table (uname varchar(100), passwd varchar(100), fullname varchar(100), phone char(50))
as
$$
    begin
        return query select * from users;
    end
$$ language plpgsql;


create or replace function get_user_by_username(varchar)
returns table (uname varchar(100), passwd varchar(100), fullname varchar(100), phone char(50))
as
$$
    begin
        return query select * from users where username = $1;
    end
$$ language plpgsql;