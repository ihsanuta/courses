CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  username varchar(40) UNIQUE NOT NULL,
  password varchar(100) NOT NULL,
  type varchar(10) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  is_deleted BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS category (
  id serial PRIMARY KEY,
  name varchar(40) UNIQUE NOT NULL,
  created_at TIMESTAMP NOT NULL
);


CREATE TABLE IF NOT EXISTS courses (
  id serial PRIMARY KEY,
  category_id integer NOT NULL,
  name varchar(40) UNIQUE NOT NULL,
  image_url varchar(100) NOT NULL DEFAULT '',
  price integer NOT NULL,
  qty integer NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL
);

INSERT INTO users(username,password,type,created_at,is_deleted) VALUES ('admin','8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918','admin',NOW(),false);
INSERT INTO category(name,created_at) VALUES ('pemograman',NOW());