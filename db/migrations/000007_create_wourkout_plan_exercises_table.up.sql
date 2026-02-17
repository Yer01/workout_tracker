CREATE TABLE IF NOT EXISTS workout_plan_exercises (
    workout_id INT NOT NULL REFERENCES workout_plans(id),
    exercise_id INT NOT NULL REFERENCES exercises(id),
    order_index INT NOT NULL,
    sets INT,
    reps INT,
    weight DOUBLE PRECISION,
    notes TEXT,
    PRIMARY KEY (workout_id, order_index)
);
