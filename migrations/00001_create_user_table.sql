-- +goose Up
-- +goose StatementBegin
CREATE TYPE role AS ENUM ('user', 'maintainer', 'admin');

CREATE TABLE users (
  id integer PRIMARY KEY,
  username integer,
  email varchar(255),
  password varchar(255),
  role role
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users
DROP TYPE role
-- +goose StatementEnd
