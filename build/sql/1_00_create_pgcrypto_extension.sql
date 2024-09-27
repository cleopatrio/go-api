SET SCHEMA 'notes';

DROP EXTENSION IF EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS pgcrypto with schema notes;
