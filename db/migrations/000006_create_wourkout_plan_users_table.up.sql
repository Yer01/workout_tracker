CREATE TABLE IF NOT EXISTS workout_plan_users (
    user_id INT NOT NULL REFERENCES users(id),
    workout_id INT NOT NULL REFERENCES workout_plans(id),
    added_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, workout_id)
);
