CREATE TABLE IF NOT EXISTS users (
  id bigserial,
  first_name text NOT NULL ,
  last_name text NOT NULL ,
  middle_name text, 
  birthdate date not NULL,
  iin bigint not NULL, 
  phone text not NULL, 
  address text not NULL, 
  email text UNIQUE NOT NULL ,
  
  _created_at timestamp DEFAULT now() NOT NULL ,
  _modified_at timestamp DEFAULT now() NOT NULL ,
  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS patients (
  id bigserial,
  blood_type smallint not NULL, 
  emergency_contact_id bigint, 
  marital_status text not null,
  FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE
  FOREIGN KEY (emergency_contact_id) REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS doctors (
  id bigserial,
  department_id bigint NOT NULL,
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

