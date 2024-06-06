CREATE TABLE party(
    id UUID PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    slogan TEXT NOT NULL,
    opened_date TIMESTAMP DEFAULT NOW(),
    description TEXT
);

CREATE TABLE public(
    id UUID PRIMARY KEY,
    first_name VARCHAR(32) NOT NULL,
    last_name VARCHAR(32) NOT NULL,
    birthday TIMESTAMP DEFAULT NOW(),
    gender VARCHAR(1),
    nation VARCHAR(64),
    party_id UUID references party(id)
);

