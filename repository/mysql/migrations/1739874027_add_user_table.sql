-- +migrate Up
CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL UNIQUE,
    created_at datetime DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down