CREATE TABLE locales (
    locale VARCHAR(5) NOT NULL,
    name VARCHAR(50) NOT NULL,
    is_default TINYINT NULL
);

CREATE TABLE translations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    locale VARCHAR(5) NOT NULL,
    entity_type VARCHAR(100) NOT NULL,
    entity_id INT NOT NULL,
    field VARCHAR(100) NOT NULL,
    translation TEXT NULL
);

INSERT INTO locales VALUES('pt', 'Português', 1);
INSERT INTO locales VALUES('en', 'English', 0);
