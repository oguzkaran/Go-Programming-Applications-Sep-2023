create table earthquakes (
    name varchar(100) primary key,
    date_time timestamp not null
);

create table earthquake_info (
    earthquake_info_id bigserial primary key,
    name varchar(100) references earthquakes(name) not null,
    query_date_time timestamp default(current_timestamp) not null
);