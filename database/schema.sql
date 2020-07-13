
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
    adress VARCHAR (255),
    longitude INTEGER,
    latitude INTEGER,
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
    renter_id INTEGER REFERENCES renter (id) ON DELETE CASCADE,
    price INTEGER,
	PRIMARY KEY(id)
) INHERITS (base_table);

CREATE TABLE alterations(
    id SERIAL,
    date TIMESTAMP,
    price INTEGER,
    description VARCHAR (255),
    flat_id INTEGER REFERENCES flat (id) ON DELETE CASCADE

) INHERITS (base_table);