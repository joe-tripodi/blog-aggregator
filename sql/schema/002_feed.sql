-- +goose Up
CREATE TABLE feed (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL UNIQUE NOT NULL,
  url TEXT NOT NULL UNIQUE NOT NULL,
  user_id UUID references users(id) ON DELETE CASCADE NOT NULL
);

-- +goose Down
DROP TABLE feed;
