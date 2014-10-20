package gomud

import (
	"strings"
)

type Skill struct {
	Name, Description, Affect string
	Delay                     int
	Costs                     *Attributes
}

type Mob struct {
	ShortName, LongName string
	Room                *Room
	Items               []*Item
	Equipped            *Equipped
	Attributes          *Attributes
	CurrentAttr         *Attributes
	Race                RaceType
	Delay               int
	Skills              []*Skill
	Disposition         Disposition
	client              *Client
}

type Disposition string

const (
	Standing Disposition = "standing"
	Sitting  Disposition = "sitting"
	Laying   Disposition = "laying"
	Sleeping Disposition = "sleeping"
)

var mobs []*Mob

func NewMob() *Mob {
	mob := &Mob{
		ShortName: "mob",
		LongName:  "A generic mob stands here, fresh out of the factory.",
		Equipped: &Equipped{
			Head: &Item{
				ShortName: "A wooden helmet",
				LongName:  "Who the hell makes a helmet out of wood?",
			},
		},
		Delay:       0,
		Disposition: Standing,
		Attributes: &Attributes{
			Hp:   20,
			Mana: 100,
			Mv:   100,
		},
		CurrentAttr: &Attributes{
			Hp:   20,
			Mana: 100,
			Mv:   100,
		},
		Room: rooms[1],
		Race: Human,
	}
	rooms[1].AddMob(mob)
	mobs = append(mobs, mob)
	return mob
}

func (m *Mob) Act(input string) string {
	if len(input) > 0 {
		args := strings.Split(input, " ")
		for _, a := range actions {
			if strings.Index(string(a.Name), args[0]) == 0 {
				return a.Func(m, args)
			}
		}
		return "What was that?\n"
	}
	return "\n"
}

func (m *Mob) Move(d Direction) string {
	if room, ok := m.Room.Rooms[d]; ok {
		if m.CurrentAttr.Mv >= m.Room.MovementCost {
			m.CurrentAttr.Mv -= m.Room.MovementCost
			for _, mob := range m.Room.Mobs {
				mob.LeftRoom(m, d)
			}
			m.Room.RemoveMob(m)
			room.AddMob(m)
			m.Room = room
			for _, mob := range m.Room.Mobs {
				od, _ := OppositeDirection(d)
				mob.EnteredRoom(m, od)
			}
			return m.Act("look")
		} else {
			return "You are too tired to move.\n"
		}
	} else {
		return "Alas, you cannot go that way.\n"
	}
}

func (m *Mob) LeftRoom(mob *Mob, d Direction) {
	if m != mob {
		m.Notify(mob.ShortName + " left heading " + string(d) + ".\n")
	}
}

func (m *Mob) EnteredRoom(mob *Mob, d Direction) {
	if m != mob {
		m.Notify(mob.ShortName + " entered from the " + string(d) + ".\n")
	}
}

func (m *Mob) DecrementDelay() {
	if m.Delay > 0 {
		m.Delay--
	}
}

func (m *Mob) Notify(message string) {
	if m.client != nil {
		m.client.Write(message)
	}
}

func (m *Mob) Tick() {
	m.DecrementDelay()
	m.Regen()
}

func (m *Mob) Pulse() {

}

func (m *Mob) Regen() {
	m.CurrentAttr.Hp += m.Attributes.Hp * 0.1
	m.CurrentAttr.Mana += m.Attributes.Mana * 0.1
	m.CurrentAttr.Mv += m.Attributes.Mv * 0.1
	m.normalizeAttr()
}

func (m *Mob) normalizeAttr() {
	if m.CurrentAttr.Hp > m.Attributes.Hp {
		m.CurrentAttr.Hp = m.Attributes.Hp
	}
	if m.CurrentAttr.Mana > m.Attributes.Mana {
		m.CurrentAttr.Mana = m.Attributes.Mana
	}
	if m.CurrentAttr.Mv > m.Attributes.Mv {
		m.CurrentAttr.Mv = m.Attributes.Mv
	}
}
