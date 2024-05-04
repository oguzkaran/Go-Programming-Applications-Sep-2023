create database postal_codes_info_db

create table postal_codes (
    code int primary key
);

create table postal_code_info (
    postal_code_info_id bigserial primary key,
    postal_code int references postal_codes(code) not null,
    admin_name1 varchar(100),
    admin_name2 varchar(100),
    longitude real not null,
    latitude real not null,
    country_code char(5) not null,
    place_name varchar(200) not null,
    ISO3166_2 varchar(10)
);

create table postal_code_search_info (
    postal_code_search_info_id bigserial primary key,
    postal_code int references postal_codes(code) not null,
    search_date_time timestamp default(current_timestamp) not null
);

