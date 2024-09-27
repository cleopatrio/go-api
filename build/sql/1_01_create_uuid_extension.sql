SET SCHEMA 'notes';

DROP EXTENSION IF EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" with schema notes;
