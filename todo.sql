CREATE TYPE todo_status_enum AS ENUM (
	'not_started',
	'done',
	'deleted'
)

CREATE TABLE users (
	id serial primary key
	name varchar
	email varchar
	password varchar
)

CREATE TABLE todos (
	id serial primary key,
	user_id int references "users"(id),
	description varchar,
	todo_status todo_status_enum default 'not_started'
)