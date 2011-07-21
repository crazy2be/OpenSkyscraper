package rooms

import (
	"os"
	"opensky/location"
	"opensky/people"
)

type Room interface {
	Location() location.Loc2D
	Inhabitants() []people.Person
	AddInhabatent(people.Person) os.Error
}