
DROP TABLE IF EXISTS alterations;
DROP TABLE IF EXISTS files;
DROP TABLE IF EXISTS lease_historic_price;
DROP TABLE IF EXISTS lease;
DROP TABLE IF EXISTS renter;
DROP TABLE IF EXISTS flat;

DROP TABLE IF EXISTS building;

DROP TABLE IF EXISTS user_account;

CREATE TABLE user_account
(
    id int NOT NULL AUTO_INCREMENT,
    created_at TIMESTAMP ,
    PRIMARY KEY (id),
    updated_at TIMESTAMP,
    email    VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role int NOT NULL

) ;

CREATE TABLE building
(
    id        int NOT NULL AUTO_INCREMENT,
    created_at TIMESTAMP ,
    updated_at TIMESTAMP,
    name      VARCHAR(255),
    address   VARCHAR(255),
    longitude FLOAT,
    latitude  FLOAT,
    PRIMARY KEY (id)
) ;


CREATE TABLE renter
(
    id   int NOT NULL AUTO_INCREMENT,
    created_at TIMESTAMP ,
    updated_at TIMESTAMP,
    name VARCHAR(255),
    nif  VARCHAR(255),
    PRIMARY KEY (id)
) ;


CREATE TABLE flat
(
    id          int NOT NULL AUTO_INCREMENT,
    created_at TIMESTAMP ,
    updated_at TIMESTAMP,
    building_id INTEGER,
    FOREIGN KEY (building_id) REFERENCES building (id) ON DELETE CASCADE,
    asked_price INTEGER,
    floor INTEGER ,
    door_number INTEGER,
    area        integer,
    boiler_date VARCHAR(255),
    boiler_description VARCHAR(255),
    price_index FLOAT,
    PRIMARY KEY (id)
) ;


CREATE TABLE lease
(
    id         int NOT NULL AUTO_INCREMENT,
    created_at TIMESTAMP ,
    updated_at TIMESTAMP,
    flat_id    INTEGER,
    renter_id  INTEGER,
    FOREIGN KEY (flat_id) REFERENCES flat (id) ON DELETE CASCADE,
    FOREIGN KEY (renter_id) REFERENCES renter (id) ON DELETE CASCADE,
    price      FLOAT,
    start_date TIMESTAMP,
    end_date   TIMESTAMP,
    deposit    FLOAT,
    PRIMARY KEY (id)
) ;


CREATE TABLE alterations
(
    id          int NOT NULL AUTO_INCREMENT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    alter_date        TIMESTAMP,
    price       INTEGER,
    description VARCHAR(255),
    flat_id INTEGER,
    FOREIGN KEY (flat_id) REFERENCES flat (id) ON DELETE CASCADE,
    PRIMARY KEY (id)

);


CREATE TABLE files
(
    id int NOT NULL AUTO_INCREMENT,
    created_at TIMESTAMP ,
    updated_at TIMESTAMP,
    flat_id   INTEGER,
    FOREIGN KEY (flat_id) REFERENCES flat (id) ON DELETE CASCADE,
    file_path VARCHAR(255),
    PRIMARY KEY (id)
);