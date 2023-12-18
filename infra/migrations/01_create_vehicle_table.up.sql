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
