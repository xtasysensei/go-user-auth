CREATE TABLE IF NOT EXISTS options (
    option_id SERIAL PRIMARY KEY,
    poll_id INTEGER REFERENCES polls(poll_id),
    text VARCHAR(255) NOT NULL
);