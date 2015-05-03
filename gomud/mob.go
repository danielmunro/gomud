package gomud

import (
	"math/rand"
	"strings"
)

/*
	Skill defines the Name, Description, and Affect of a Mob's skill, as well
	as the Delay from using it and the Costs of its use.
	Name - the name of the Skill.
	Description - text description of what the skill does.
	Affect - how the use of the skill appears to others.
	Delay - how many ticks the skill takes to perform.
	Costs - how the attributes of a Mob are affected by using the skill.
*/
type Skill struct {
	Name, Description, Affect string
	Delay                     int
	Costs                     *Attributes
}

/*
	Mob defines the attributes of a mobile entity within the game. This can be
	a monster, player, NPC, or anything else that moves.

	ShortName - a short string naming the Mob.
	LongName - a long string naming the Mob.
	Room - the current Room location of the Mob.
	Items - an array of the Items currently within the Mob's possession.
	Equipped - the items that the Mob is currently using, represented as an
				Equipped struct.
	Attributes - the (maximum) Attributes of the Mob.
	CurrentAttr - the current values of the Mob's attributes.
	Race - the fantasy Race to which the Mob belongs.
	Delay - the tick delay between actions for the Mob.
	Skills - an array of the Mob's current Skills.
	Disposition - A Disposition describing how the Mob appears to others.
	Wanders - a float64 representing how likely it is for the Mob to move
				on its own.
*/
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

	//target represents the Mob that this Mob is attacking.
	target *Mob

	//hasWandered indicates whether this Mob is moving on its own.
	hasWandered bool
	client      *Client
}

/*
	Disposition is a string describing how a Mob appears at a given moment.
	i.e. This mob is "standing" or "crouching".
*/
type Disposition string

/*
	Four default Dispositions.
*/
const (
	Standing Disposition = "standing"
	Sitting  Disposition = "sitting"
	Laying   Disposition = "laying"
	Sleeping Disposition = "sleeping"
)

/*
	A global array of all Mobs.
*/
var mobs []*Mob

/*
	NewMob creates and returns a new Mob with default values.
*/
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

/*
	Act performs whatever action is specified within the input string. If the
	string is not within the recognized set of actions for that Mob, it will
	print "What was that?"
*/
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

/*
	Move attempts to move the Mob in Direction d. If the Mob has the movement
	necessary to travel to that room, it will move them there and automatically
	invoke the "look" action. If they lack the movement attribute, it will
	print "You are too tired to move." If there is no room in the specified direction,
	it will print "Alas, you cannot go that way." Move also notifies all Mobs remaining
	in the room of the passage of the Mob that moved and notifies all Mobs in the new
	room of the Mob's arrival.
*/
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

/*
	LeftRoom notifies a Mob that another Mob left a room and in what direction they were
	going.
*/
func (m *Mob) LeftRoom(mob *Mob, d Direction) {
	if m != mob {
		m.Notify(mob.ShortName + " left heading " + string(d) + ".\n")
	}
}

/*
	EnteredRoom notifies a Mob that another Mob entered the same room from a given direction.
*/
func (m *Mob) EnteredRoom(mob *Mob, d Direction) {
	if m != mob {
		m.Notify(mob.ShortName + " entered from the " + string(d) + ".\n")
	}
}

/*
	DecrementDelay decreases the delay before a Mob can move again by 1
	if it is above zero.
*/
func (m *Mob) DecrementDelay() {
	if m.Delay > 0 {
		m.Delay--
	}
}

/*
	Notify writes a message string to a Mob's Client.
*/
func (m *Mob) Notify(message string) {
	if m.client != nil {
		m.client.Write(message)
	}
}

/*
	Tick decrements the delay before a Mob can act again and partially regenerates depleted
	attributes (like vitals).
*/
func (m *Mob) Tick() {
	m.DecrementDelay()
	m.Regen()
	//set flag to false
	m.hasWandered = false
}

/*
	Pulse causes a Mob to act, either by attacking their current target or potentially wandering
	if they are not currently targeting any other Mobs.
*/
func (m *Mob) Pulse() {
	if m.Wanders > 0 && !m.hasWandered && m.target == nil {
		m.Wander()
	}
	if m.target != nil {
		m.Attack()
	}
}

/*
	Attack causes the Mob to attack whatever its current target is. This decrements the target's
	current Hp Attribute by 5 and notifies both Mobs of the attack. If the target's Hp is reduced to
	0 or lower, it dies. If the attacked Mob survives, but does not have a target, it is set to
	target the attacking Mob.
*/
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

/*
	Die informs a Mob of its death and creates a container to represent their corpse.
	This container is automatically added to the contents of a room.
*/
func (m *Mob) Die() {
	m.Notify("You died!\n")
	m.target = nil
	corpse := Container{}
	corpse.ShortName = "the corpse of " + strings.ToLower(m.ShortName)
	corpse.LongName = "The corpse of " + strings.ToLower(m.ShortName) + " lies here."
	m.Room.Items = append(m.Room.Items, corpse)
}

/*
	Wander makes a Mob wander one room in a given direction. The percent chance of this
	behavior occurring is based on the Mob.Wanders value.
*/
func (m *Mob) Wander() {
	if rand.Float64() < m.Wanders {
		d := m.Room.AllDirections()
		l := len(d)
		i := rand.Intn(l)
		m.Move(d[i])
		m.hasWandered = true
	}
}

/*
	Regen causes the Mob's current Vitals and other Attributes to increase
	by a value.
*/
func (m *Mob) Regen() {
	m.CurrentAttr.Vitals.Hp += m.Attributes.Vitals.Hp * 0.1
	m.CurrentAttr.Vitals.Mana += m.Attributes.Vitals.Mana * 0.1
	m.CurrentAttr.Vitals.Mv += m.Attributes.Vitals.Mv * 0.1
	m.normalizeAttr()
}

/*
	Status returns a string representing the current status of the Mob.
*/
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

/*
	normalizeAttr checks whether the current Attributes of a Mob are greater than
	the maximum attributes and sets them to be their maximum values if they are.
*/
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
