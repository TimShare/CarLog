INSERT INTO car_makes (id, name, country_code, logo_url)
OVERRIDING SYSTEM VALUE
VALUES
    (100, 'Toyota', 'JP', 'https://example.test/logos/toyota.svg'),
    (101, 'BMW', 'DE', 'https://example.test/logos/bmw.svg');

INSERT INTO car_models (
    id,
    make_id,
    name,
    production_start_year,
    production_end_year
)
OVERRIDING SYSTEM VALUE
VALUES
    (1000, 100, 'Camry XV70', 2017, 2024),
    (1001, 100, 'Corolla E210', 2018, NULL),
    (1100, 101, '3 Series G20', 2018, NULL);

INSERT INTO cars (id, model_id, vin, year, current_mileage, created_at, updated_at)
OVERRIDING SYSTEM VALUE
VALUES
    (
        10000,
        1000,
        'JTNB11HK5K3000001',
        2019,
        145000,
        '2026-01-10 10:00:00+00',
        '2026-01-10 10:00:00+00'
    ),
    (
        10001,
        1001,
        'JTDBR32E720000001',
        2020,
        82000,
        '2026-02-15 12:30:00+00',
        '2026-03-01 09:15:00+00'
    ),
    (
        10002,
        1100,
        NULL,
        2021,
        61000,
        '2026-04-20 08:45:00+00',
        '2026-04-20 08:45:00+00'
    );

SELECT setval(
    pg_get_serial_sequence('car_makes', 'id'),
    (SELECT max(id) FROM car_makes)
);
SELECT setval(
    pg_get_serial_sequence('car_models', 'id'),
    (SELECT max(id) FROM car_models)
);
SELECT setval(
    pg_get_serial_sequence('cars', 'id'),
    (SELECT max(id) FROM cars)
);
