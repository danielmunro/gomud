package gomud

import (
	"math/rand"
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
	Wanders             float64
	target              *Mob
	hasWandered         bool
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
			Head: &Equipment{
				"A wooden helmet",
				"Who the hell makes a helmet out of wood?",
				&Attributes{},
				Head,
			},
		},
		Delay:       0,
		Disposition: Standing,
		Attributes: &Attributes{
			Vitals: &Vitals{
				Hp:   20,
				Mana: 100,
				Mv:   100,
			},
		},
		CurrentAttr: &Attributes{
			Vitals: &Vitals{
				Hp:   20,
				Mana: 100,
				Mv:   100,
			},
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
		if m.CurrentAttr.Vitals.Mv >= m.Room.MovementCost {
			m.CurrentAttr.Vitals.Mv -= m.Room.MovementCost
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
		m.client.write(message)
	}
}

func (m *Mob) Tick() {
	m.DecrementDelay()
	m.Regen()
	m.hasWandered = false
}

func (m *Mob) Pulse() {
	if m.Wanders > 0 && !m.hasWandered && m.target == nil {
		m.Wander()
	}
	if m.target != nil {
		m.Attack()
	}
}

func (m *Mob) Attack() {
	m.target.CurrentAttr.Vitals.Hp -= 5
	m.Notify("You attack " + strings.ToLower(m.target.ShortName) + ".\n")
	m.target.Notify(m.ShortName + " attacks you.\n")
	if m.target.CurrentAttr.Vitals.Hp < 0 {
		m.Notify("You killed " + strings.ToLower(m.target.ShortName) + "!\n")
		m.target.Die()
		m.target = nil
		if m.client != nil {
			m.client.prompt()
		}
	} else if m.target.target == nil {
		m.target.target = m
	}
}

func (m *Mob) Die() {
	m.Notify("You died!\n")
	m.target = nil
	corpse := Container{}
	corpse.ShortName = "the corpse of " + strings.ToLower(m.ShortName)
	corpse.LongName = "The corpse of " + strings.ToLower(m.ShortName) + " lies here."
	m.Room.Items = append(m.Room.Items, corpse)
}

func (m *Mob) Wander() {
	if rand.Float64() < m.Wanders {
		d := m.Room.AllDirections()
		l := len(d)
		i := rand.Intn(l)
		m.Move(d[i])
		m.hasWandered = true
	}
}

func (m *Mob) Regen() {
	m.CurrentAttr.Vitals.Hp += m.Attributes.Vitals.Hp * 0.1
	m.CurrentAttr.Vitals.Mana += m.Attributes.Vitals.Mana * 0.1
	m.CurrentAttr.Vitals.Mv += m.Attributes.Vitals.Mv * 0.1
	m.normalizeAttr()
}

func (m *Mob) Status() (status string) {
	p := m.CurrentAttr.Vitals.Hp / m.Attributes.Vitals.Hp
	switch {
	case p <= .1:
		return "is in awful condition"
	case p <= .15:
		return "looks pretty hurt"
	case p <= .30:
		return "has some big nasty wounds and scratches"
	case p <= .50:
		return "has quite a few wounds"
	case p <= .75:
		return "has some small wounds and bruises"
	case p <= .99:
		return "has a few scratches"
	default:
		return "is in excellent condition"
	}
}

func (m *Mob) normalizeAttr() {
	if m.CurrentAttr.Vitals.Hp > m.Attributes.Vitals.Hp {
		m.CurrentAttr.Vitals.Hp = m.Attributes.Vitals.Hp
	}
	if m.CurrentAttr.Vitals.Mana > m.Attributes.Vitals.Mana {
		m.CurrentAttr.Vitals.Mana = m.Attributes.Vitals.Mana
	}
	if m.CurrentAttr.Vitals.Mv > m.Attributes.Vitals.Mv {
		m.CurrentAttr.Vitals.Mv = m.Attributes.Vitals.Mv
	}
}
