CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    space_id INTEGER REFERENCES spaces(id) ON DELETE CASCADE,
    parent_id INTEGER REFERENCES directories(id) ON DELETE SET NULL,
    name TEXT NOT NULL,
    type TEXT CHECK (type IN ('page', 'table', 'task', 'er')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_files_deleted_at ON files (deleted_at);
