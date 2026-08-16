CREATE SCHEMA IF NOT EXISTS chatgo;

CREATE TABLE IF NOT EXISTS chatgo.users (
    id            SERIAL PRIMARY KEY,
    username      TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS chatgo.messages (
    id          SERIAL PRIMARY KEY,
    sender_id   INTEGER NOT NULL REFERENCES chatgo.users(id) ON DELETE CASCADE,
    receiver_id INTEGER NOT NULL REFERENCES chatgo.users(id) ON DELETE CASCADE,
    content     TEXT NOT NULL,
    sent_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    is_read     BOOLEAN DEFAULT FALSE
);