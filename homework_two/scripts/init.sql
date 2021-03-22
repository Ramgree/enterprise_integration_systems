CREATE TABLE plant (
    plant_id INT GENERATED ALWAYS AS IDENTITY,
    PRIMARY KEY(plant_id),
    plant_daily_rental_price numeric,
    plant_name VARCHAR(100) NOT NULL,
    plant_type_name VARCHAR(100) NOT NULL
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

INSERT INTO plant
    (plant_name, plant_daily_rental_price, plant_type_name)
VALUES
    ('excavator', 1250,'rune'),
    ('bulldozer', 5000, 'adamant'),
    ('crane', 62500, 'mithril'),
    ('dumper', 2500, 'rune');

INSERT INTO booking
    (plant_id, total_rental_cost, start_date, end_date)
VALUES
    ((SELECT plant_id FROM plant WHERE plant_name ILIKE '%crane%'),
    ((SELECT (plant_daily_rental_price * (SELECT(EXTRACT(DAY FROM ('2021-11-21 00:00:00'::timestamp - '2021-11-16 00:00:00'::timestamp))))) FROM plant WHERE plant_name ILIKE '%crane%')),
    '2021-11-16 00:00:00'::timestamp,
    '2021-12-21 00:00:00'::timestamp);
