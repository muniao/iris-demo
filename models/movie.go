package models

//实体对象
type Movie struct {

    Id     int  `json:"id"`
    Name   string `json:"name"`
    Year   int    `json:"year"`
    Poster string `json:"poster"`
}
