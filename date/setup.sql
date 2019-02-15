drop table account;

create table account(
    id  SERIAL PRIMARY KEY,
    aname varchar(100) not null unique,
    apassword varchar(100) not null  default '123456', 
    amoney double precision not null default 10.0,
    lasttime timestamp(0) without time zone
);