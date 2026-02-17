CREATE TABLE IF NOT EXISTS session_sets (
    set_num INT NOT NULL,
    session_id INT NOT NULL REFERENCES workout_sessions(id),
    exercise_id INT NOT NULL REFERENCES exercises(id),
    reps INT,
    weight DOUBLE PRECISION,
    PRIMARY KEY (session_id, set_num)
);
