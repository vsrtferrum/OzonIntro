-- +goose Up
-- +goose StatementBegin
CREATE Table comments(
    comment_id BIGSERIAL PRIMARY KEY,
    post_id BIGINT REFERENCES posts(id), 
    text VARCHAR(2000) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE comments;
-- +goose StatementEnd
