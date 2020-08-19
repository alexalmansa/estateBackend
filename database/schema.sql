DROP TABLE IF EXISTS base_table CASCADE ;
CREATE TABLE base_table (
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE user_account (
    id SERIAL,
    email VARCHAR (255) UNIQUE NOT NULL,
    password VARCHAR (255) NOT NULL,
    PRIMARY KEY (id)
) INHERITS (base_table);

CREATE TABLE building(
    id SERIAL,
    name VARCHAR (255),
    address VARCHAR (255),
    longitude DECIMAL(9,6),
    latitude DECIMAL(9,6),
	PRIMARY KEY(id)
) INHERITS (base_table);

CREATE TABLE renter(
    id SERIAL,
    name VARCHAR (255),
    age INTEGER,
	PRIMARY KEY(id)
) INHERITS (base_table);

CREATE TABLE flat(
    id SERIAL,
    building_id INTEGER REFERENCES building (id) ON DELETE CASCADE,
    asked_price INTEGER,
    number_door VARCHAR (255),
    area integer,
	PRIMARY KEY(id)
) INHERITS (base_table);

CREATE TABLE lease(
    id SERIAL,
    flat_id INTEGER REFERENCES flat (id) ON DELETE CASCADE,
    renter_id INTEGER REFERENCES renter (id) ON DELETE CASCADE,
    price FLOAT,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    deposit FLOAT,
	PRIMARY KEY(id)
) INHERITS (base_table);

CREATE TABLE lease_historic_price(
    id SERIAL,
    lease_id INTEGER REFERENCES lease (id) ON DELETE CASCADE,
    price FLOAT,
	PRIMARY KEY(id)
) INHERITS (base_table);

CREATE TABLE alterations(
    id SERIAL,
    date TIMESTAMP,
    price INTEGER,
    description VARCHAR (255),
    flat_id INTEGER REFERENCES flat (id) ON DELETE CASCADE

) INHERITS (base_table);

CREATE TABLE files(
    id SERIAL,
    flat_id INTEGER REFERENCES flat (id) ON DELETE CASCADE,
    file_path VARCHAR (255)
) INHERITS (base_table);