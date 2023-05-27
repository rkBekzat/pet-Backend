CREATE TABLE Articles(
                         aid SERIAL PRIMARY KEY,
                         username VARCHAR(100),
                         id INT,
                         species VARCHAR(150),
                         breed VARCHAR(150),
                         type VARCHAR(150),
                         photo BYTEA,
                         date DATE
);

CREATE TABLE Owner(
                      username VARCHAR(100) PRIMARY KEY,
                      name VARCHAR(150),
                      surname VARCHAR(150),
                      address_id INT,
                      email VARCHAR(150),
                      phone_number VARCHAR(15)
);


CREATE TABLE Users(
                      id SERIAL NOT NULL UNIQUE ,
                      username VARCHAR(100) PRIMARY KEY UNIQUE ,
                      name VARCHAR(150) NOT NULL ,
                      password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE Address(
                        address_id SERIAL PRIMARY KEY,
                        country VARCHAR(100),
                        city VARCHAR(100),
                        State_Province VARCHAR(100),
                        zip_postal_code VARCHAR(10),
                        address_line VARCHAR(150)
);

CREATE TABLE Pet(
                    id SERIAL PRIMARY KEY,
                    username VARCHAR(100),
                    species VARCHAR(40),
                    breed VARCHAR(40),
                    name VARCHAR(20),
                    date_of_birth DATE,
                    color VARCHAR(20),
                    sex BOOLEAN,
                    tattoo VARCHAR(20),
                    issued_organization VARCHAR(20)
);

CREATE TABLE Found(
                      fid SERIAL PRIMARY KEY,
                      username VARCHAR(100),
                      address_id INT,
                      id INT,
                      photo BYTEA,
                      expenses REAL
);

CREATE TABLE Lost(
                     lid SERIAL PRIMARY KEY ,
                     username VARCHAR(100),
                     address_id INT,
                     id INT,
                     photo BYTEA,
                     reward REAL,
                     date DATE
);
