-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE chat_users (
    id SERIAL PRIMARY KEY,
    chat_id INT NOT NULL,
    user_id INT NOT NULL
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    chat_id INT NOT NULL,
    user_id INT NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE messages;
DROP TABLE chat_users;
DROP TABLE chats;
-- +goose StatementEnd
