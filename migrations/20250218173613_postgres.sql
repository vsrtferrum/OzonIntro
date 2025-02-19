-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts(
    id BIGSERIAL PRIMARY KEY, 
    text TEXT NOT NULL,
    comments BOOLEAN
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
