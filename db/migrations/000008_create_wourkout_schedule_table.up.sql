CREATE TABLE IF NOT EXISTS workouts_schedule (
    workout_id INT NOT NULL REFERENCES workout_plans(id),
    user_id INT NOT NULL REFERENCES users(id),
    scheduled_date TIMESTAMPTZ NOT NULL,
    status VARCHAR(50) NOT NULL,
    PRIMARY KEY (workout_id, user_id, scheduled_date)
);
