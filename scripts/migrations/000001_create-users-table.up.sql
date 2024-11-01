CREATE TABLE IF NOT EXISTS  users (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  email VARCHAR(250) NOT NULL UNIQUE,
  password VARCHAR(500) NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);