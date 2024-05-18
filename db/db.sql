CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE attendance_type AS ENUM ('CHECKIN', 'CHECKOUT');

CREATE TABLE config (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    field VARCHAR(255) NOT NULL UNIQUE,
    value TEXT
)

CREATE TABLE attendant (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    employee_id VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    badge_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE record_attendance (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    employee_id VARCHAR(255) NOT NULL,
    time TIMESTAMP NOT NULL,
    attendance_type attendance_type NOT NULL
);