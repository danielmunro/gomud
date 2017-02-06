package gomud

func startRoom(l *Listener) *room {
	r1 := newRoom("Room 1", "You are in the first room")
	r2 := newRoom("Room 2", "You are in the second room")
	r3 := newRoom("Room 3", "You are in the third room")

	r1.exits = append(r1.exits, newExit(r2, dSouth))
	r1.exits = append(r1.exits, newExit(r3, dWest))

	m := &mob{
		name:        "a test mob",
		description: "A test mob",
		room:        r1,
		roles:       []role{mobile, scavenger},
	}
	r1.mobs = append(r1.mobs, m)
	l.mobs = append(l.mobs, m)

	r2.exits = append(r2.exits, newExit(r1, dNorth))
	r3.exits = append(r3.exits, newExit(r1, dEast))

	i1 := newItem("an item", "An item is here", []string{"item"})
	i2 := newItem("an item", "An item is here", []string{"item"})

	i1.position = held
	i2.position = held

	r1.items = append(r1.items, i1)
	r1.items = append(r1.items, i2)

	return r1
}
