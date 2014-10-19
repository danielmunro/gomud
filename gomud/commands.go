package gomud

import "strings"

type Command struct {
	Name string
	Func func(m *Mob, args []string) string
}

var commands []*Command

func init() {
	commands = []*Command{
		&Command{
			Name: "north",
			Func: func(m *Mob, args []string) string {
				return m.Move(North)
			},
		},
		&Command{
			Name: "south",
			Func: func(m *Mob, args []string) string {
				return m.Move(South)
			},
		},
		&Command{
			Name: "east",
			Func: func(m *Mob, args []string) string {
				return m.Move(East)
			},
		},
		&Command{
			Name: "west",
			Func: func(m *Mob, args []string) string {
				return m.Move(West)
			},
		},
		&Command{
			Name: "up",
			Func: func(m *Mob, args []string) string {
				return m.Move(Up)
			},
		},
		&Command{
			Name: "down",
			Func: func(m *Mob, args []string) string {
				return m.Move(Down)
			},
		},
		&Command{
			Name: "look",
			Func: func(m *Mob, args []string) string {
				output := m.Room.Title + "\n" + m.Room.Description + "\n\n[Exits "
				for d, _ := range m.Room.Rooms {
					output += string(d)[:1]
				}
				output += "]\n"
				for _, mob := range m.Room.Mobs {
					if mob != m {
						output += mob.ShortName + " is " + string(mob.Disposition) + " here.\n"
					}
				}
				return output
			},
		},
		&Command{
			Name: "say",
			Func: func(m *Mob, args []string) string {
				message := strings.Join(args[1:], " ")
				for _, mob := range m.Room.Mobs {
					if mob != m {
						mob.Notify(m.ShortName + " says, \"" + message + "\"\n")
					}
				}
				return "You say, \"" + message + "\"\n"
			},
		},
		&Command{
			Name: "equipped",
			Func: func(m *Mob, args []string) string {
				equipped := ""
				for key, value := range m.Equipped.getAll() {
					if value != nil {
						equipped += key.String() + ": " + value.ShortName
					} else {
						equipped += key.String() + ": <none>"
					}
					equipped += "\n"
				}
				return equipped
			},
		},
	}
}
