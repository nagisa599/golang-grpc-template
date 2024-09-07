-- +goose Up
ALTER TABLE users DROP COLUMN id;
ALTER TABLE users ADD COLUMN id INT AUTO_INCREMENT PRIMARY KEY FIRST;
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE users DROP COLUMN id;
ALTER TABLE users ADD COLUMN id INT AUTO_INCREMENT PRIMARY KEY FIRST;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
