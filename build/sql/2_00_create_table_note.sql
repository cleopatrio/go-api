SET SCHEMA 'notes';

CREATE TABLE note
(
    id          uuid        NOT NULL DEFAULT uuid_generate_v4(),
    description varchar     NOT NULL,
    status      varchar     NOT NULL,
    created_at  timestamptz NOT NULL DEFAULT NOW(),
    updated_at  timestamptz NOT NULL DEFAULT NOW(),
    CONSTRAINT payers_pkey PRIMARY KEY (id)
);
