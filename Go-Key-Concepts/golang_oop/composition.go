package main

import "fmt"

// Composition

type Musical interface {
	Play() string
}

// Musician
type Musician struct {
	Name       string
	Instrument string
}

func (m *Musician) Play() string {
	return fmt.Sprintf("%s is playing the %s.", m.Name, m.Instrument)
}

// Vocalist
type Vocalist struct {
	Name string
}

func (v *Vocalist) Play() string {
	return fmt.Sprintf("%s is singing.", v.Name)
}

// DJ
type DJ struct {
	Name string
}

func (d *DJ) Play() string {
	return fmt.Sprintf("%s is spinning tracks.", d.Name)
}

// Music Group
type Band struct {
	Name    string
	Members []Musical
}

func (b *Band) PerformConcert() string { // PlayMusic()
	musicNotes := ""
	for _, member := range b.Members {
		musicNotes += fmt.Sprintf("%s \n", member.Play())
	}
	if musicNotes == "" {
		return fmt.Sprintf("Error in Playing Music!")
	}

	return musicNotes
}

func main() {
	john := &Musician{Name: "John", Instrument: "Guitar"}
	paul := &Musician{Name: "Paul", Instrument: "Bass"}
	george := &Musician{Name: "George", Instrument: "Guitar"}
	ringo := &Musician{Name: "Ringo", Instrument: "Drums"}

	// new additional member's that are joined with Musical Interface implementation and Play()
	lennon := &Vocalist{Name: "Lennon"}
	dj := &DJ{Name: "DJ Cool"}

	beatlesMembers := []Musical{john, paul, george, ringo, lennon, dj}

	beatles := Band{Name: "The Beatles", Members: beatlesMembers}

	fmt.Println(beatles.PerformConcert())
}
