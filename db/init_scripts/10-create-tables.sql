CREATE TYPE status AS ENUM ('ACTIVE', 'IN_ACTIVE');

CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    email VARCHAR(150),
    status status NOT NULL DEFAULT 'ACTIVE',
    time_created timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    time_modified timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX users_email_index ON users(email);