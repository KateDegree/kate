CREATE TABLE directories (
    id SERIAL PRIMARY KEY,
    space_id INTEGER REFERENCES spaces(id) ON DELETE CASCADE,
    parent_id INTEGER REFERENCES directories(id) ON DELETE SET NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_directories_deleted_at ON directories (deleted_at);
