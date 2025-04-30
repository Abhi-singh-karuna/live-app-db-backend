-- Create test table
CREATE TABLE IF NOT EXISTS test (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

-- Create index on name column for faster lookups
CREATE INDEX IF NOT EXISTS idx_test_name ON test(name);

-- Add any additional tables or modifications here 