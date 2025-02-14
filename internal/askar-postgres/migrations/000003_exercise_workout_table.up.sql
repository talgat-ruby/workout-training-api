CREATE TABLE IF NOT EXISTS exercise_workout (
    id bigserial PRIMARY KEY,
    exercise_id BIGINT NOT NULL,
    workout_id BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_exercise
        FOREIGN KEY (exercise_id)
        REFERENCES exercises (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_workout
        FOREIGN KEY (workout_id)
        REFERENCES workouts (id)
        ON DELETE CASCADE
);
