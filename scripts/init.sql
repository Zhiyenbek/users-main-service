CREATE TABLE IF NOT EXISTS users (
  id bigserial,
  first_name text NOT NULL ,
  last_name text NOT NULL ,
  middle_name text, 
  birthdate text not NULL,
  iin text not NULL, 
  phone text not NULL, 
  address text not NULL, 
  email text,
  
  _created_at timestamp DEFAULT now() NOT NULL ,
  _modified_at timestamp DEFAULT now() NOT NULL ,
  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS patients (
  id bigint,
  blood_type smallint not NULL, 
  emergency_contact text not NULL, 
  marital_status text not NULL,
  FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS doctors (
  id bigint,
  department_id int NOT NULL,
  spec_id int not NULL, 
  experience int not NULL, 
  photo text not NULL, 
  category text not NULL, 
  price int not NULL, 
  schedule text not NULL, 
  degree text not NULL, 
  rating int not NULL,
  website_url text,
  FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS specs (
  id bigserial,
  name text NOT NULL,
  description text NOT NULL
);

CREATE TABLE IF NOT EXISTS departments (
  id bigserial,
  name text NOT NULL
);

CREATE TABLE IF NOT EXISTS auth (
  user_id bigserial,
  login text not NULL,
  password text not NULL
);

CREATE INDEX departments_idx_id ON "departments" ("id");
CREATE INDEX doctors_idx_id ON "doctors" ("id");
CREATE INDEX users_idx_id ON "users" ("ID");

INSERT INTO auth VALUES('admin', '$2a$12$aJdxplWONcENmkOWWNahce6io526pkprBg9ZVE2jhozd2IgSrFKJe')