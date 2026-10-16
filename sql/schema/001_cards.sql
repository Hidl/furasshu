-- +goose Up
CREATE TABLE cards (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    lesson INT,
    dutch TEXT NOT NULL,
    hiragana TEXT,
    katakana TEXT,
    kanji TEXT,
    notes TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE cards;