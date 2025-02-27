CREATE TABLE IF NOT EXISTS exercises (
    id bigserial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    repetitions INTEGER NOT NULL,
    sets INTEGER NOT NULL,
    weight INTEGER NOT NULL,
    muscle_groups text [] NOT NULL,
    categories text [] NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);