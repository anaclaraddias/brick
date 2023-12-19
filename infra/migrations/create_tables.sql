CREATE TABLE vehicle (
    id UUID PRIMARY KEY NOT NULL,
    brand VARCHAR(60) NOT NULL,
    model VARCHAR(60) NOT NULL,
    vehicle_year VARCHAR(4) NOT NULL,
    license_plate VARCHAR(7) NOT NULL,
    renavam VARCHAR(11) NOT NULL,
    vehicle_value FLOAT NOT NULL,
    cargo FLOAT NOT NULL,
    vehicle_height FLOAT NOT NULL,
    vehicle_width FLOAT NOT NULL,
    vehicle_length FLOAT NOT NULL,
    vehicle_type VARCHAR(12) CHECK (vehicle_type IN ('personal', 'business')),
    update_date DATE NULL,
    creation_date DATE NOT NULL
);

CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(80) NOT NULL,
    email VARCHAR(80) NOT NULL,
    cellphone VARCHAR(14) NOT NULL,
    cpf VARCHAR(11) NOT NULL,
    cnpj VARCHAR(14) NULL,
    update_date DATE NULL,
    creation_date DATE NOT NULL
);

CREATE TABLE policy (
    id UUID PRIMARY KEY NOT NULL,
    status VARCHAR(12) CHECK (status IN ('pending', 'active', 'canceled', 'renovation')),
    coverage_limit FLOAT  NULL,
    value FLOAT  NULL,
    start_date DATE NULL,
    end_date DATE NULL,
    user_id UUID NOT NULL,
    update_date DATE NULL,
    creation_date DATE NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE coverage (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(80) NOT NULL, 
    description VARCHAR(200) NULL, 
    rate_value FLOAT NOT NULL,
    creation_date DATE NOT NULL
);

INSERT INTO coverage (
    name,
    descricao,
    rate_value,
    creation_date
) VALUES 
('Theft and Robbery', 'protects against the theft or robbery of insured property', 1000.00, '2023-12-18'),
( 'Collision', 'cover the cost of repairs or replacement of the insured vehicle in case of a collision with another vehicle or object', 800.00, '2023-12-18'),
('Third Parties', 'provides protection for the insured against claims made by third parties', 1200.00, '2023-12-18');

CREATE TABLE linked_policy_vehicle (
    id UUID PRIMARY KEY NOT NULL,
    vehicle_id UUID NOT NULL,
    policy_id UUID NOT NULL,
    creation_date DATE NOT NULL,
    FOREIGN KEY (vehicle_id) REFERENCES vehicle (id),
    FOREIGN KEY (policy_id) REFERENCES policy (id)
);


CREATE TABLE linked_policy_coverage (
    id UUID PRIMARY KEY NOT NULL,
    coverage_id SERIAL NOT NULL,
    policy_id UUID NOT NULL,
    creation_date DATE NOT NULL,
    FOREIGN KEY (coverage_id) REFERENCES coverage (id),
    FOREIGN KEY (policy_id) REFERENCES policy (id)
);
