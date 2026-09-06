BEGIN;

ALTER TABLE car_models
    DROP COLUMN IF EXISTS production_end_year,
    DROP COLUMN IF EXISTS production_start_year;

ALTER TABLE car_makes
    DROP COLUMN IF EXISTS logo_url,
    DROP COLUMN IF EXISTS country_code;

COMMIT;
