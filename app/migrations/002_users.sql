-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Users (
    Id INT AUTO_INCREMENT,
    Email VARCHAR(256) NOT NULL,
    Username VARCHAR(256),
    UPassword TEXT,
    CreatedAt DATETIME, 
    PRIMARY KEY(Id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS UsersUrl (
    Id INT AUTO_INCREMENT,
    UrlId INT NOT NULL,
    UserId INT NOT NULL,
    CreatedAt DATETIME,
    PRIMARY KEY(Id),
    FOREIGN KEY(UrlId) REFERENCES Urls(Id),
    FOREIGN KEY(UserId) REFERENCES Users(Id)
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE Visitors 
ADD CreatedAt DATETIME;
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
ALTER TABLE Visitors
DROP CreatedAt;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE UsersUrl;
DROP TABLE Users;
-- +goose StatementEnd