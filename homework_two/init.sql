CREATE TABLE plant_type (
    plant_type_id INT GENERATED ALWAYS AS IDENTITY,
    PRIMARY KEY(plant_type_id),
    plant_type_name VARCHAR(100) NOT NULL
);

CREATE TABLE plant (
    plant_id INT GENERATED ALWAYS AS IDENTITY,
    PRIMARY KEY(plant_id),
    plant_type_id int,
    plant_daily_rental_price numeric,
    CONSTRAINT fk_plant_type
        FOREIGN KEY (plant_type_id)
        REFERENCES plant_type(plant_type_id),
    plant_name VARCHAR(100) NOT NULL
);

CREATE TABLE booking (
    plant_id int,
    CONSTRAINT fk_plant_id
        FOREIGN KEY (plant_id)
        REFERENCES plant(plant_id),
    total_rental_cost numeric,
    start_date TIMESTAMP,
    end_date TIMESTAMP
);

INSERT INTO plant_type
    (plant_type_name)
VALUES
    ('mithril'),
    ('adamant'),
    ('rune'),
    ('dragon');

INSERT INTO plant
    (plant_type_id, plant_daily_rental_price, plant_name)
VALUES
    ((SELECT plant_type_id FROM plant_type WHERE plant_type_name ILIKE '%rune%'),
     1250,
     'excavator'),
    ((SELECT plant_type_id FROM plant_type WHERE plant_type_name ILIKE '%adamant%'),
     5000,
     'bulldozer'),
    ((SELECT plant_type_id FROM plant_type WHERE plant_type_name ILIKE '%mithril%'),
     62500,
     'crane'),
    ((SELECT plant_type_id FROM plant_type WHERE plant_type_name ILIKE '%rune%'),
     2500,
     'dumper'),
    ((SELECT plant_type_id FROM plant_type WHERE plant_type_name ILIKE '%dragon%'),
     5000,
     'forklift'),
    ((SELECT plant_type_id FROM plant_type WHERE plant_type_name ILIKE '%adamant%'),
     1000,
     'mewp'),
    ((SELECT plant_type_id FROM plant_type WHERE plant_type_name ILIKE '%rune%'),
     2500,
     'sweeper'),
    ((SELECT plant_type_id FROM plant_type WHERE plant_type_name ILIKE '%mithril%'),
     100000,
     'road roller');

INSERT INTO booking
    (plant_id, total_rental_cost, start_date, end_date)
VALUES
    ((SELECT plant_id FROM plant WHERE plant_name ILIKE '%road roller%'),
    ((SELECT (plant_daily_rental_price * (SELECT(EXTRACT(DAY FROM ('2021-10-21 00:00:00'::timestamp - '2021-10-19 00:00:00'::timestamp))))) FROM plant WHERE plant_name ILIKE '%road roller%')),
    '2021-10-19 00:00:00'::timestamp,
    '2021-10-21 00:00:00'::timestamp);
