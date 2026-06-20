CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hashed TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sender_id INTEGER NOT NULL,
    receiver_id_user INTEGER,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    rec_type TEXT NOT NULL,
    receiver_id_group INTEGER,
    FOREIGN KEY (sender_id) REFERENCES users(id),
    FOREIGN KEY (receiver_id_user) REFERENCES users(id),
    FOREIGN KEY (receiver_id_group) REFERENCES groups(id),

    CHECK (
        (receiver_id_group IS NULL AND receiver_id_user IS NOT NULL)
        OR
        (receiver_id_group IS NOT NULL AND receiver_id_user IS NULL)
    )
);

CREATE TABLE IF NOT EXISTS groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    creator INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(creator) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS group_members (
    member_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,

    PRIMARY KEY (member_id, group_id),
    FOREIGN KEY (member_id) REFERENCES users(id),
    FOREIGN KEY (group_id) REFERENCES groups(id)
);
