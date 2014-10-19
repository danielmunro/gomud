package gomud

import (
    "strings"
)

type Skill struct {
    Name, Description, Affect string
    Delay int
    Costs *Attributes
}

type Race struct {
    Name, Description string
    Attributes *Attributes
    Skills []*Skill
}

type Mob struct {
    ShortName, LongName string
    Room *Room
    Items []*Item
    Equipped *Equipped
    Attributes *Attributes
    Race *Race
    Delay int
    Skills []*Skill
    client *Client
}

func NewMob() *Mob {
    mob := &Mob{
        ShortName: "mob",
        LongName: "A generic mob stands here, fresh out of the factory.",
        Equipped: &Equipped {
            Head: &Item {
                ShortName: "A wooden helmet",
                LongName: "Who the hell makes a helmet out of wood?",
            },
        },
        Delay: 0,
        Room: rooms[1],
    }
    rooms[1].AddMob(mob)
    return mob
}

func (m *Mob) Act(input string) string {
    if len(input) > 0 {
        args := strings.Split(input, " ")
        for _, c := range commands {
            if strings.Index(c.Name, args[0]) == 0 {
                return c.Func(m, args)
            }
        }
        return "What was that?\n"
    }
    return "\n"
}

func (m *Mob) Move(d Direction) string {
    if room, ok := m.Room.Rooms[d]; ok {
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
        return "Alas, you cannot go that way.\n"
    }
}

func (m *Mob) LeftRoom(mob *Mob, d Direction) {
    if m != mob {
        m.Notify(mob.ShortName + " left heading "+string(d)+".\n")
    }
}

func (m *Mob) EnteredRoom(mob *Mob, d Direction) {
    if m != mob {
        m.Notify(mob.ShortName + " entered from the "+string(d)+".\n")
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
