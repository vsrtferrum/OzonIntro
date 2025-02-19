-- +goose Up
-- +goose StatementBegin
CREATE TABLE comment_ref(
    original_comment BIGINT REFERENCES comments(comment_id),
    reference_comment BIGINT REFERENCES comments(comment_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE comment_ref;
-- +goose StatementEnd
