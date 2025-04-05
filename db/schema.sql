CREATE TYPE question_type AS ENUM ('multiple-choice', 'open-ended');
CREATE TYPE subscription_tier AS ENUM ('Explorer', 'Scholar');

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    subscription_status subscription_tier NOT NULL
);

CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    type question_type NOT NULL,
    question TEXT NOT NULL,
    answer TEXT NOT NULL,
    source TEXT,
    explanation TEXT,
    options JSONB
);

CREATE TABLE answered_questions (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    question_id INTEGER NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    correct BOOLEAN NOT NULL,
    date TIMESTAMP NOT NULL,
    time_taken INTEGER NOT NULL,
    user_answer TEXT NOT NULL
);
