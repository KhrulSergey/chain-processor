CREATE TABLE IF NOT EXISTS users
(
    id                  UUID     DEFAULT generateUUIDv4(),
    login               String,
    password            String,
    first_name          String,
    second_name         String,
    gender              String,
    date_of_birth       Date,
    country             String,
    city                String,
    address             String,
    email               String,
    phone               String,
    registration_date   Date,
    registration_ip     String,
    registration_device String,
    is_active           UInt8    DEFAULT 0,
    activation_code     String,
    created_at          DateTime DEFAULT now(),
    updated_at          DateTime DEFAULT now(),
    PRIMARY KEY (id)
) ENGINE = MergeTree()
      ORDER BY id;

INSERT INTO users (id, login, password, first_name, second_name, gender, date_of_birth, country, city, address, email,
                   phone, registration_date, registration_ip, registration_device, is_active, activation_code,
                   created_at, updated_at)
VALUES ('13d4a949-b9fc-438a-8fef-e60a02abdc3d', 'adminPWG',
        '$2a$12$Xq9ZRPaXT51H.lY3d459/eesYHvz3UBui/9dQYbUExgpm08lQFPxy',
        'super admin', 'PayWithGlass', 'Male',
        '2001-01-01',
        'USA',
        'Washington',
        '1600 Pennsylvania Avenue',
        'admin@paywithglass.com',
        '+312223334455',
        today(),
        '127.0.0.1',
        'pg_console',
        1,
        '',
        now(),
        now());
