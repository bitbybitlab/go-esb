CREATE TABLE Users (
ref uuid default gen_random_uuid() NOT NULL PRIMARY KEY,
username varchar(50) NOT NULL,
password varchar(50) NOT NULL,
first_name varchar(50),
last_name varchar(50),
middle_name varchar(50),
phone varchar(30),
email varchar(50),
create_time timestamp default (now() at time zone 'utc') NOT NULL,
update_time timestamp default (now() at time zone 'utc') NOT NULL,
version int NOT NULL default 0
)