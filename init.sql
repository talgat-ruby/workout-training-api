CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
);

INSERT INTO users (email, password_hash) VALUES
('test@example.com', 'hashedpassword');
