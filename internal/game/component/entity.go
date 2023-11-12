package component

var nextID int = 0

func getNextID() int {
	nextID++
	return nextID
}

type Entity int

func NewEntity() Entity {
	e := Entity(getNextID())
	return e
}
