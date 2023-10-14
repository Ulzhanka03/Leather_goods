CREATE TABLE IF NOT EXISTS leathergoods (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    type text NOT NULL,
    price integer NOT NULL,
    leather_type text NOT NULL,
    color text NOT NULL,
    version integer NOT NULL DEFAULT 1
    );
