CREATE TABLE
    IF NOT EXISTS users (
        id UUID PRIMARY KEY,
        name text NOT NULL UNIQUE,
        nickname text,
        line_id text UNIQUE
    )
