-- +goose Up
-- +goose StatementBegin
ALTER TABLE chat_users DROP CONSTRAINT chat_users_pkey;
ALTER TABLE chat_users DROP COLUMN user_name;
ALTER TABLE chat_users ADD COLUMN user_id INT NOT NULL;
ALTER TABLE chat_users ADD CONSTRAINT chat_users_pkey PRIMARY KEY (chat_id, user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE chat_users DROP CONSTRAINT chat_users_pkey;
ALTER TABLE chat_users DROP COLUMN user_id;
ALTER TABLE chat_users ADD COLUMN user_name VARCHAR(255) NOT NULL;
ALTER TABLE chat_users ADD CONSTRAINT chat_users_pkey PRIMARY KEY (chat_id, user_name);
-- +goose StatementEnd
