package gomud

import "strings"

// ActionName is the string name of a valid command in the game.
type ActionName string

// The default Actions available to players
const (
	NorthAction    ActionName = "north"
	SouthAction    ActionName = "south"
	EastAction     ActionName = "east"
	WestAction     ActionName = "west"
	UpAction       ActionName = "up"
	DownAction     ActionName = "down"
	LookAction     ActionName = "look"
	SayAction      ActionName = "say"
	EquippedAction ActionName = "equipped"
	QuitAction     ActionName = "quit"
	ScoreAction    ActionName = "score"
	KillAction     ActionName = "kill"
)

// Action defines a single action available to Mobs.
// Name - an ActionName for this action.
// Func - the function to call when this action is taken.
type Action struct {
	Name ActionName
	Func func(m *Mob, args []string) string
}

// actions is a global array of available actions.
var actions []*Action

// init populates the global actions array.
func init() {
	actions = []*Action{
		&Action{
			Name: NorthAction,
			Func: func(m *Mob, args []string) string {
				return m.Move(North)
			},
		},
		&Action{
			Name: SouthAction,
			Func: func(m *Mob, args []string) string {
				return m.Move(South)
			},
		},
		&Action{
			Name: EastAction,
			Func: func(m *Mob, args []string) string {
				return m.Move(East)
			},
		},
		&Action{
			Name: WestAction,
			Func: func(m *Mob, args []string) string {
				return m.Move(West)
			},
		},
		&Action{
			Name: UpAction,
			Func: func(m *Mob, args []string) string {
				return m.Move(Up)
			},
		},
		&Action{
			Name: DownAction,
			Func: func(m *Mob, args []string) string {
				return m.Move(Down)
			},
		},
		&Action{
			Name: LookAction,
			Func: func(m *Mob, args []string) (output string) {
				if len(args) > 1 {
					mob := m.Room.FindMob(args[1])
					if mob != nil {
						output = mob.LongName + "\n"
						output += mob.ShortName + " " + m.Status() + ".\n"
					}
				} else {

					output = m.Room.Title + "\n" + m.Room.Description + "\n\n[Exits "
					for _, e := range m.Room.Exits {
						output += string(e.Direction)[:1]
					}
					output += "]\n"
					for _, i := range m.Room.Items {
						output += strings.ToUpper(i.ShortName) + " is here.\n"
					}
					for _, mob := range m.Room.Mobs {
						if mob != m {
							output += mob.ShortName + " is " + string(mob.Disposition) + " here.\n"
						}
					}
				}
				if output == "" {
					output = "You don't see that here.\n"
				}
				return output
			},
		},
		&Action{
			Name: ScoreAction,
			Func: func(m *Mob, args []string) string {
				output := "You are " + m.ShortName + ", a " + string(m.Race) + "\n"

				return output
			},
		},
		&Action{
			Name: SayAction,
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
		&Action{
			Name: KillAction,
			Func: func(m *Mob, args []string) string {

				if m.target != nil {
					return "You are already fighting!\n"
				}

				target := m.Room.FindMob(args[1])
				if target != nil {
					m.target = target
					m.Room.Announce(m, m.ShortName+" screams and attacks "+target.ShortName+"\n")
					return "You scream and attack!\n"
				}
				return "You don't see them here.\n"
			},
		},
		&Action{
			Name: EquippedAction,
			Func: func(m *Mob, args []string) string {
				equipped := ""
				for key, value := range m.Equipped.getAll() {
					if value != nil {
						equipped += string(key) + ": " + value.ShortName
					} else {
						equipped += string(key) + ": <none>"
					}
					equipped += "\n"
				}
				return equipped
			},
		},
		&Action{
			Name: QuitAction,
			Func: func(m *Mob, args []string) string {
				m.client.write("Goodbye!\n")
				m.client.server.removeClient(m.client)
				return ""
			},
		},
	}
}
