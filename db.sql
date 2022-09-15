CREATE DATABASE Battleships;
USE Battleships;

CREATE TABLE Players (
    Id int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    Name VARCHAR(25) NOT NULL,
    Email VARCHAR(50) NOT NULL
);

CREATE TABLE Games (
    Id int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    PlayerId int NOT NULL,
    OpponentId int NOT NULL,
    TurnCount int NOT NULL DEFAULT 0,
    FOREIGN KEY (PlayerId) REFERENCES Players(Id),
    FOREIGN KEY (OpponentId) REFERENCES Players(Id)
);


CREATE TABLE Boards (
    Id int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    Fields VARCHAR(100) NOT NULL
);

ALTER TABLE Games
    ADD COLUMN PlayerBoardId int NOT NULL,
    ADD COLUMN OpponentBoardId int NOT NULL,
    ADD FOREIGN KEY PlayerBoardFK(PlayerBoardId) REFERENCES Boards(Id),
    ADD FOREIGN KEY OpponentBoardFK(OpponentBoardId) REFERENCES Boards(Id);

ALTER TABLE Games
    ADD COLUMN Status int NOT NULL;

ALTER TABLE Boards
    ADD COLUMN ShipCount int NOT NULL;

ALTER TABLE Boards
    ADD COLUMN PlayerId int NOT NULL,
    ADD FOREIGN KEY PlayerIdFK(PlayerId) REFERENCES Players(Id);