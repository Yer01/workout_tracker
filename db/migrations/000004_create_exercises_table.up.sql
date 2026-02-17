CREATE TABLE IF NOT EXISTS exercises (
    id SERIAL PRIMARY KEY,
    link VARCHAR(255),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    duration DOUBLE PRECISION
);
