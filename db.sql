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
