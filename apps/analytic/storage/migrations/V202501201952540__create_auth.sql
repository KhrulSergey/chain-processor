CREATE TABLE IF NOT EXISTS open_id_clients
(
    id UUID DEFAULT generateUUIDv4(),
    secret String NOT NULL,
    created_at DateTime DEFAULT now(),
    updated_at DateTime DEFAULT now(),
    PRIMARY KEY (id)
) ENGINE = MergeTree()
      ORDER BY id;

CREATE TABLE IF NOT EXISTS open_id_authentication_requests
(
    id UUID DEFAULT generateUUIDv4(),
    user_id UUID NOT NULL,
    client_id UUID NOT NULL,
    code String NOT NULL,
    scope String NOT NULL,
    redirect_uri String NOT NULL,
    created_at DateTime DEFAULT now(),
    updated_at DateTime DEFAULT now(),
    PRIMARY KEY (id)
) ENGINE = MergeTree()
      ORDER BY id;

CREATE TABLE IF NOT EXISTS sessions
(
    id UUID DEFAULT generateUUIDv4(),
    user_id UUID NOT NULL,
    session_id String NOT NULL,
    version Int32 NOT NULL,
    access_token String NOT NULL,
    refresh_token String NOT NULL,
    login_date Date NOT NULL,
    login_ip String NOT NULL,
    login_device String NOT NULL,
    expires_at DateTime NOT NULL,
    is_valid UInt8 NOT NULL,
    created_at DateTime DEFAULT now(),
    updated_at DateTime DEFAULT now(),
    PRIMARY KEY (id)
) ENGINE = MergeTree()
      ORDER BY id;

INSERT INTO open_id_clients (id, secret, created_at, updated_at)
VALUES ('6feb7b69-cd99-4e07-b491-132ae38d5ee6', '$2a$12$Io5yBhBdg1W2gYavJrqknOP2pXcmMf7zJcNm0LgUv7qTTo4QcSTty',
        now(), now());