create table users (
                       username varchar(100) primary key,
                       password varchar(100) not null,
                       name varchar(300) not null,
                       phone char(50) not null,
                       birth_date date not null,
                       register_date_time timestamp default(current_timestamp) not null
    -- ...
);

create table products (
                          code varchar(100) primary key,
                          name varchar(300) not null,
                          stock double precision default(0) not null,
                          cost double precision not null,
                          unit_price double precision not null,
                          expiry_date date
    -- ...
);

create table payments (
                          payment_id bigserial primary key,
                          username varchar(100) references users(username) not null,
                          product_code varchar(100) references products(code) not null,
                          amount double precision check(amount > 0) not null,
                          unit_price double precision not null,
                          pay_date_time timestamp
    -- ...
);

create or replace procedure sp_insert_user(varchar(100), varchar(100), varchar(100), char(50), date)
language plpgsql
as
$$
    begin
    insert into users (username, password, name, phone, birth_date) values ($1, $2, $3, $4, $5);
    end
$$;

create or replace procedure sp_insert_user(varchar(100), varchar(100), varchar(100), char(50), date, timestamp)
language plpgsql
as
$$
begin
insert into users (username, password, name, phone, birth_date, register_date_time) values ($1, $2, $3, $4, $5, $6);
end
$$;

create or replace function get_all_users()
returns table (uname varchar(100), passwd varchar(100), fullname varchar(100), phone char(50), bdate date, rdatetime timestamp)
as
$$
begin
return query select * from users;
end
$$ language plpgsql;


create or replace function get_user_by_username(varchar)
returns table (uname varchar(100), passwd varchar(100), fullname varchar(100), phone char(50), bdate date, rdatetime timestamp)
as
$$
begin
return query select * from users where username = $1;
end
$$ language plpgsql;


create or replace function get_user_count()
returns table (n_users bigint)
as
$$
begin
return query select count(*) from users;
end
$$ language plpgsql;

create or replace function user_exists(varchar(100))
returns table (u_exits bool)
as
$$
begin
return query select exists (select * from users where username = $1);
end
$$ language plpgsql;