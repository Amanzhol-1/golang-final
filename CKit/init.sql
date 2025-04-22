CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE shipments (
  id                UUID        PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id           TEXT        NOT NULL,
  from_address      TEXT        NOT NULL,
  to_address        TEXT        NOT NULL,
  pickup_time       TIMESTAMPTZ NOT NULL,
  delivery_price    NUMERIC(10,2) NOT NULL,
  price_negotiable  BOOLEAN     NOT NULL DEFAULT FALSE,
  weight            NUMERIC(10,2) NOT NULL,
  volume            NUMERIC(10,2) NOT NULL,
  cargo_type        VARCHAR(255) NOT NULL,
  sender_name       VARCHAR(255) NOT NULL,
  sender_phone      VARCHAR(20)  NOT NULL,
  receiver_name     VARCHAR(255) NOT NULL,
  receiver_phone    VARCHAR(20)  NOT NULL,
  additional_notes  TEXT
);

-- Indexes for faster lookups by user and by date
CREATE INDEX idx_shipments_user_id      ON shipments(user_id);
CREATE INDEX idx_shipments_pickup_time  ON shipments(pickup_time);