CREATE TABLE Users (
id uuid default gen_random_uuid() NOT NULL PRIMARY KEY,
username varchar(50) NOT NULL,
password text NOT NULL,
first_name varchar(50),
last_name varchar(50),
middle_name varchar(50),
phone varchar(30),
email varchar(50),
create_time timestamp default (now() at time zone 'utc') NOT NULL,
update_time timestamp default (now() at time zone 'utc') NOT NULL,
version int NOT NULL default 0
);

CREATE TABLE external_systems (
id uuid default gen_random_uuid() NOT NULL PRIMARY KEY,
name varchar(100) UNIQUE NOT NULL,
ip varchar(50),
port varchar(50),
path varchar(300),
driver text,
create_time timestamp default (now() at time zone 'utc') NOT NULL,
update_time timestamp default (now() at time zone 'utc') NOT NULL,
version int NOT NULL default 0
);

CREATE TABLE external_users (
id uuid default gen_random_uuid() NOT NULL PRIMARY KEY,
name varchar(100) UNIQUE NOT NULL,
system uuid NOT NULL REFERENCES systems (id),
username varchar(50),
password varchar(50),
token varchar(100),
create_time timestamp default (now() at time zone 'utc') NOT NULL,
update_time timestamp default (now() at time zone 'utc') NOT NULL,
version int NOT NULL default 0
);