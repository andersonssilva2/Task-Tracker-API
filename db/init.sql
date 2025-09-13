CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE
);

-- Optional: seed data
INSERT INTO tasks (title, completed) VALUES
('Learn Docker', false),
('Build Task Tracker API', true);
