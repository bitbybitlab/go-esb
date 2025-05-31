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

CREATE TABLE connection_types (
id uuid default gen_random_uuid() NOT NULL PRIMARY KEY,
name varchar(150) UNIQUE NOT NULL,
description varchar(300)
);

CREATE TABLE external_systems (
id uuid default gen_random_uuid() NOT NULL PRIMARY KEY,
name varchar(100) UNIQUE NOT NULL,
type uuid NOT NULL REFERENCES connection_types (id),
ip varchar(50),
port varchar(50),
path varchar(300),
driver text,
create_time timestamp default (now() at time zone 'utc') NOT NULL,
update_time timestamp default (now() at time zone 'utc') NOT NULL,
version int NOT NULL default 0
);

CREATE TABLE authentication_types (
id uuid default gen_random_uuid() NOT NULL PRIMARY KEY,
name varchar(150) UNIQUE NOT NULL,
description varchar(300)
);

INSERT INTO public.authentication_types(
	name, description)
	VALUES ('Basic', 'Auth with username and password');
INSERT INTO public.authentication_types(
	name, description)
	VALUES ('BearerToken', 'Bearer token (JWT-token)');

CREATE TABLE external_users (
id uuid default gen_random_uuid() NOT NULL PRIMARY KEY,
name varchar(100) UNIQUE NOT NULL,
system uuid NOT NULL REFERENCES external_systems (id),
type uuid NOT NULL REFERENCES authentication_types (id),
username varchar(50),
password varchar(50),
token varchar(100),
create_time timestamp default (now() at time zone 'utc') NOT NULL,
update_time timestamp default (now() at time zone 'utc') NOT NULL,
version int NOT NULL default 0
);