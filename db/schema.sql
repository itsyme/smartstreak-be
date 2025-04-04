CREATE TYPE question_type AS ENUM ('multiple-choice', 'open-ended');
CREATE TYPE subscription_tier AS ENUM ('Explorer', 'Scholar');

CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    type question_type NOT NULL,
    question TEXT NOT NULL,
    answer TEXT NOT NULL,
    source TEXT,
    explanation TEXT
);

CREATE TABLE multiple_choice_options (
    id SERIAL PRIMARY KEY,
    question_id INTEGER REFERENCES questions(id) ON DELETE CASCADE,
    value TEXT NOT NULL,
    text TEXT NOT NULL
);

CREATE TABLE answered_questions (
    id SERIAL PRIMARY KEY,
    question TEXT NOT NULL,
    correct BOOLEAN NOT NULL,
    date TIMESTAMP NOT NULL,
    time_taken INTEGER NOT NULL,
    user_answer TEXT NOT NULL,
    correct_answer TEXT NOT NULL,
    explanation TEXT,
    source TEXT
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    subscription_status subscription_tier NOT NULL
);
