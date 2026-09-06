BEGIN;

ALTER TABLE car_makes
    ADD COLUMN country_code char(2),
    ADD COLUMN logo_url text,
    ADD CONSTRAINT car_makes_country_code_check
        CHECK (country_code IS NULL OR country_code ~ '^[A-Z]{2}$'),
    ADD CONSTRAINT car_makes_logo_url_check
        CHECK (logo_url IS NULL OR btrim(logo_url) <> '');

ALTER TABLE car_models
    ADD COLUMN production_start_year smallint,
    ADD COLUMN production_end_year smallint,
    ADD CONSTRAINT car_models_production_start_year_check
        CHECK (
            production_start_year IS NULL
            OR production_start_year BETWEEN 1886 AND 2100
        ),
    ADD CONSTRAINT car_models_production_end_year_check
        CHECK (
            production_end_year IS NULL
            OR production_end_year BETWEEN 1886 AND 2100
        ),
    ADD CONSTRAINT car_models_production_years_check
        CHECK (
            production_start_year IS NULL
            OR production_end_year IS NULL
            OR production_end_year >= production_start_year
        );

COMMIT;
