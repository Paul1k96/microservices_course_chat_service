-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE chat_users (
    chat_id INT NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    PRIMARY KEY (chat_id, user_name),
    FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE CASCADE
);

CREATE TABLE messages (
    id UUID PRIMARY KEY,
    chat_id INT NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE messages;
DROP TABLE chat_users;
DROP TABLE chats;
-- +goose StatementEnd
