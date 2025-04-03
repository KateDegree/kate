CREATE TABLE space_users (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    space_id INTEGER REFERENCES spaces(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, space_id)
);
