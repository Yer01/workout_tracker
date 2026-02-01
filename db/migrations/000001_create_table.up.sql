CREATE TABLE IF NOT EXISTS workouts(
    workout_id serial PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    content TEXT,
    exercises TEXT,
    duration FLOAT
)