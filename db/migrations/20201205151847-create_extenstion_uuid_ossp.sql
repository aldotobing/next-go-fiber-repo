
-- +migrate Up
CREATE extension IF NOT EXISTS "uuid-ossp";

-- +migrate Down
DROP extension  IF EXISTS "uuid-ossp";