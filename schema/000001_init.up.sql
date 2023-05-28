CREATE TABLE Address(
                        address_id SERIAL PRIMARY KEY,
                        country VARCHAR(100),
                        city VARCHAR(100),
                        state_province VARCHAR(100),
                        zip_postal_code VARCHAR(10),
                        address_line VARCHAR(150)
);

CREATE TABLE Photo(
                      photo_id SERIAL PRIMARY KEY ,
                      photo BYTEA
);

CREATE TABLE Vaccine(
                        vaccine_id SERIAL PRIMARY KEY ,
                        name VARCHAR(50),
                        virus VARCHAR(50),
                        price REAL
);

CREATE TABLE Vet(
                    vet_id SERIAL PRIMARY KEY ,
                    name VARCHAR(150),
                    surname VARCHAR(150),
                    email VARCHAR(150),
                    phone_number VARCHAR(15)
);


CREATE TABLE Clinic(
                       clinic_id SERIAL PRIMARY KEY,
                       name VARCHAR(50),
                       phone_number VARCHAR(15),
                       address_id INT,
                       FOREIGN KEY(address_id)
                           REFERENCES Address(address_id)
);

CREATE TABLE Pet_hotel(
                          hotel_id SERIAL PRIMARY KEY,
                          name VARCHAR(50),
                          address_id INT,
                          FOREIGN KEY (address_id)
                              REFERENCES Address(address_id)
);

CREATE TABLE Pet_shop(
                         shop_id SERIAL PRIMARY KEY ,
                         name VARCHAR(50),
                         address_id INT,
                         FOREIGN KEY (address_id)
                             REFERENCES Address(address_id)
);

CREATE TABLE Owner(
                      user_id SERIAL PRIMARY KEY ,
                      username VARCHAR(100),
                      name VARCHAR(150),
                      surname VARCHAR(150),
                      email VARCHAR(150),
                      phone_number VARCHAR(15),
                      address_id INT,
                      photo_id INT,
                      FOREIGN KEY (address_id)
                          REFERENCES Address(address_id)
                          ON UPDATE CASCADE,
                      FOREIGN KEY (photo_id)
                          REFERENCES Photo(photo_id)
                          ON DELETE SET NULL
);

CREATE TABLE Pet(
                    pet_id SERIAL PRIMARY KEY,
                    name VARCHAR(20),
                    species VARCHAR(40),
                    breed VARCHAR(40),
                    color VARCHAR(20),
                    date_of_birth DATE,
                    sex BOOLEAN,
                    tattoo VARCHAR(20),
                    issued_organization VARCHAR(20),
                    photo_id_1 INT,
                    user_id INT,
                    FOREIGN KEY (photo_id_1)
                        REFERENCES Photo(photo_id)
                        ON DELETE SET NULL,
                    FOREIGN KEY (user_id)
                        REFERENCES Owner(user_id)
);


CREATE TABLE Articles(
                         article_id SERIAL PRIMARY KEY,
                         text VARCHAR(500),
                         date DATE,
                         address_id INT,
                         photo_id INT,
                         user_id_1 INT,
                         pet_id INT,
                         FOREIGN KEY (address_id)
                             REFERENCES Address(address_id)
                             ON UPDATE CASCADE ,
                         FOREIGN KEY (photo_id)
                             REFERENCES Photo(photo_id)
                             ON DELETE SET NULL,
                         FOREIGN KEY (user_id_1)
                             REFERENCES Owner(user_id),
                         FOREIGN KEY (pet_id)
                             REFERENCES Pet(pet_id)
);

CREATE TABLE Found(
                      found_id SERIAL PRIMARY KEY,
                      article_id INT,
                      expenses REAL,
                      FOREIGN KEY (article_id)
                          REFERENCES Articles(article_id)
                          ON DELETE CASCADE
);

CREATE TABLE Lost(
                     lost_id SERIAL PRIMARY KEY ,
                     article_id INT,
                     reward REAL,
                     FOREIGN KEY (article_id)
                         REFERENCES Articles(article_id)
                         ON DELETE CASCADE
);

CREATE TABLE Article_comments(
                                 comment_id SERIAL PRIMARY KEY ,
                                 text TEXT,
                                 article_id_1 INT,
                                 FOREIGN KEY (article_id_1)
                                     REFERENCES Articles(article_id)
                                     ON DELETE CASCADE
);

CREATE TABLE "Order"(
                        order_id SERIAL PRIMARY KEY ,
                        article_id INT,
                        price REAL,
                        FOREIGN KEY (article_id)
                            REFERENCES Articles(article_id)
                            ON DELETE CASCADE
);

CREATE TABLE Analysis(
                         analysis_id SERIAL PRIMARY KEY ,
                         date DATE,
                         diagnosis TEXT,
                         results TEXT,
                         clinic_id INT,
                         vet_id INT,
                         pet_id INT,
                         price REAL,
                         FOREIGN KEY (clinic_id)
                             REFERENCES Clinic(clinic_id),
                         FOREIGN KEY (vet_id)
                             REFERENCES Vet(vet_id),
                         FOREIGN KEY (pet_id)
                             REFERENCES Pet(pet_id)
);