CREATE TABLE IF NOT EXISTS firebase_authentications
(
    account_id INT          NOT NULL,
    uid        VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY (account_id, uid),
    FOREIGN KEY (account_id) REFERENCES accounts (id) ON DELETE CASCADE
);