package app

import (
	"log"
	"time"
)

type Url struct {
	Id        int64
	Link      string
	LinkId    string
	CreatedAt time.Time
}

func GetUrl(linkId string) (*Url, error) {
	query := `
		SELECT Id, Link, LinkId, CreatedAt FROM Urls
		WHERE LinkId = ?
	`

	row := db.QueryRow(query, linkId)

	var u Url
	err := row.Scan(&u.Id, &u.Link, &u.LinkId, &u.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func CreateUrl(link string) (*Url, error) {
	query := `
		INSERT INTO Urls (Link, CreatedAt) VALUES(?, ?, ?)
	`

	dateTime := time.Now()

	result, err := db.Exec(query, link, dateTime)

	if err != nil {
		log.Println("Error while creating URL", link)
		return nil, err
	}

	Id, err := result.LastInsertId()

	linkId := GenerateHash(Id)

	query = `
		UPDATE Urls
		SET LinkId = ?
		WHERE Id = ?
	`

	result, err = db.Exec(query, linkId, Id)

	if err != nil {
		log.Println("Error while setting Link Id on URL", link)
		return nil, err
	}

	return &Url{Id, link, linkId, dateTime}, nil
}

type Visitor struct {
	Id               int
	UrlId            int
	Device           string
	Country          string
	ScreenResolution string
	Browser          string
	OsName           string
}
