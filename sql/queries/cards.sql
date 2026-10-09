-- name: CreateCard :exec
INSERT INTO cards (id, lesson, romaji, hiragana, katakana, kanji, notes) 
VALUES (
    DEFAULT,
    $1, 
    $2,
    $3, 
    $4, 
    $5, 
    $6
);

-- name: GetCardByRomaji :one
SELECT * FROM cards WHERE romaji = $1;