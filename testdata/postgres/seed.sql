create extension if not exists "uuid-ossp";
CREATE TABLE IF NOT EXISTS message
(
    id                      uuid           NOT NULL DEFAULT uuid_generate_v4(),
    content                 varchar(255)   NOT NULL,
    seen                    bool           NOT NULL DEFAULT false,
    CONSTRAINT message_pkey PRIMARY KEY (id)
);
