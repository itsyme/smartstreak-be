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

CREATE TABLE daily_questions (
    date DATE PRIMARY KEY,
    question_ids INTEGER[] NOT NULL
);

-- Creates row in users table when a new user is created in auth.users
CREATE OR REPLACE FUNCTION public.handle_new_user()
RETURNS trigger AS $$
BEGIN
  INSERT INTO public.users (id, name, email, subscription_status)
  VALUES (
    NEW.id,
    NEW.raw_user_meta_data->>'name',
    NEW.email,
    'Explorer' -- default subscription tier
  )
  ON CONFLICT (id) DO NOTHING; -- prevents duplicate inserts just in case

  RETURN NEW;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;

-- Trigger on auth.users
CREATE TRIGGER on_auth_user_created
AFTER INSERT ON auth.users
FOR EACH ROW
EXECUTE FUNCTION public.handle_new_user();