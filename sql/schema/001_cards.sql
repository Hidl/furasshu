-- +goose Up
CREATE TABLE cards (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    lesson INT,
    romaji TEXT NOT NULL UNIQUE,
    hiragana TEXT,
    katakana TEXT,
    kanji TEXT,
    notes TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE cards;