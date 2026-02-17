CREATE TABLE IF NOT EXISTS exercises_musclegroup (
    exercise_id INT NOT NULL REFERENCES exercises(id),
    musclegroup_id INT NOT NULL REFERENCES musclegroups(id),
    is_primary BOOLEAN NOT NULL,
    PRIMARY KEY (exercise_id, musclegroup_id)
);
