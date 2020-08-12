CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email VARCHAR(255),
    password_hash VARCHAR(255),
    name VARCHAR(255),
    user_token VARCHAR(255),
    email_confirmation_token VARCHAR(255),
    email_confirmed TINYINT,
    reset_password_token VARCHAR(255),
    reset_password_requested_at DATETIME,
    created_at DATETIME
);

CREATE TABLE users_login_tokens (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INT,
    user_login_token VARCHAR(255),
    user_login_token_expires_at DATETIME
);
