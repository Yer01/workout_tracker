CREATE TABLE IF NOT EXISTS workout_plans (
    id SERIAL PRIMARY KEY,
    total_duration DOUBLE PRECISION,
    name VARCHAR(200) NOT NULL,
    user_id INT NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    description TEXT
);
