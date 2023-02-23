package main

type Artist struct {
	Id          int    //PK id for lookup
	Name        string //artist name
	Image       string //link to photo
	Works       []Work //list of assoc. works
	Description string //short blurb
}
