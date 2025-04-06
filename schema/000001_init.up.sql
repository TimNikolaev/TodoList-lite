CREATE TABLE users
(
  id serial not null unique,
  name varchar(255) not null,
  email varchar(255) not null unique,
  password_hash varchar(255) not null
);

CREATE TABLE todo_tasks
(
  id serial not null unique,
  title varchar(255) not null,
  description varchar(255),
  done boolean not null default false
);

CREATE TABLE users_tasks
(
  id serial not null unique,
  user_id int references users (id) on delete cascade not null,
  task_id int references todo_tasks (id) on delete cascade not null
);