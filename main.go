package a_module

import (
	"fmt"
)

var Songs []Song
var Singers []Singer

type Song struct {
	Title  string
	Singer string
	Year   int
}

type Singer struct {
	Name  string
	Songs int
}

func (s *Singer) Sing() {
	fmt.Println("singer sing a song")

}

func (c *Song) Year_of_song(s *Singer) {

	fmt.Println("in ", c.Year, " ", s.Name, " sing this song ", c.Title)

}

func Add_song(title string, singer string, year int, s *Singer) {

	new_song := Song{title, singer, year}

	Songs = append(Songs, new_song)

	if new_song.Singer == s.Name {
		s.Songs++
	}
}

func Add_singer(name string, songs int) {

	new_singer := Singer{name, songs}

	Singers = append(Singers, new_singer)

}
