BEGIN;

CREATE TABLE car_makes (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name text NOT NULL CHECK (btrim(name) <> ''),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX car_makes_name_unique_idx
    ON car_makes (lower(name));

CREATE TABLE car_models (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    make_id bigint NOT NULL REFERENCES car_makes (id) ON DELETE RESTRICT,
    name text NOT NULL CHECK (btrim(name) <> ''),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX car_models_make_id_name_unique_idx
    ON car_models (make_id, lower(name));

CREATE TABLE cars (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    model_id bigint NOT NULL REFERENCES car_models (id) ON DELETE RESTRICT,
    vin varchar(17),
    year smallint NOT NULL CHECK (year BETWEEN 1886 AND 2100),
    current_mileage integer NOT NULL DEFAULT 0 CHECK (current_mileage >= 0),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CHECK (vin IS NULL OR btrim(vin) <> '')
);

CREATE INDEX cars_model_id_idx
    ON cars (model_id);

CREATE UNIQUE INDEX cars_vin_unique_idx
    ON cars (upper(vin))
    WHERE vin IS NOT NULL;

COMMIT;
