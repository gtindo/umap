-- +goose Up
-- +goose StatementBegin
CREATE TABLE Urls (
    Id INT AUTO_INCREMENT,
    Link TEXT NOT NULL,
    LinkId VARCHAR(256) NOT NULL,
    CreatedAt DATETIME, 
    PRIMARY KEY(Id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE Visitors (
    Id INT AUTO_INCREMENT,
    UrlId INT NOT NULL,
    Device VARCHAR(100),
    Country VARCHAR(100),
    ScreenResolution VARCHAR(20),
    Browser VARCHAR(30),
    OsName VARCHAR(30),
    PRIMARY KEY(Id),
    FOREIGN KEY(UrlId) REFERENCES Urls(Id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Visitors;
DROP TABLE Urls;
-- +goose StatementEnd
