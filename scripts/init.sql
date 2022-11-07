CREATE TABLE IF NOT EXISTS users (
  id bigserial,
  first_name text NOT NULL ,
  last_name text NOT NULL ,
  middle_name text, 
  birthdate text not NULL,
  iin bigint not NULL, 
  phone text not NULL, 
  address text not NULL, 
  email text UNIQUE ,
  
  _created_at timestamp DEFAULT now() NOT NULL ,
  _modified_at timestamp DEFAULT now() NOT NULL ,
  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS patients (
  id bigserial,
  blood_type smallint not NULL, 
  emergency_contact text not NULL, 
  marital_status text not NULL,
  user_id bigint not NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS doctors (
  id bigserial,
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
  user_id bigint not NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

