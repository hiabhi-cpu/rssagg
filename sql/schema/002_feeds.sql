-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    feed_name VARCHAR(255) UNIQUE NOT NULL,
    feed_url VARCHAR(255) UNIQUE NOT NULL,
    user_id UUID NOT NULL,
    constraint fk foreign key (user_id) references users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;