CREATE DATABASE IF NOT EXISTS bank;
USE bank;

CREATE TABLE IF NOT EXISTS
users(
    id          INTEGER AUTO_INCREMENT PRIMARY KEY,
    login        TEXT NOT NULL,
    money_amount        INTEGER NOT NULL,
    card_number       TEXT NOT NULL,
    status       BOOL NOT NULL
);

CREATE TABLE IF NOT EXISTS
users_passwords(
    id          INTEGER UNIQUE NOT NULL,
    password            TEXT NOT NULL
);

INSERT users(login, money_amount, card_number, status)
VALUES
('admin', 1000, '5173113236846375', 1),
('wendy', 100000000, '5341218336284489', 1),
('cool_girl', 345550, '5363270717384225', 0),
('cool_guy', 123544818, '5523961242234078', 1),
('meowmeow', 26534, '5271643493585747', 0);

INSERT users_passwords(id, password)
VALUES
(1, 'adminpswd'),
(2, 'ydnew'),
(3, 'cool_password'),
(4, 'greatPassword288'),
(5, 'barkquack');