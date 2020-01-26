package gomud

import (
	"errors"
	"github.com/danielmunro/gomud/io"
	"log"
	"strings"
)

type GameService struct {
	mobService *MobService
	locationService *LocationService
	roomService *RoomService
	server *Server
	buffers []*io.Buffer
	eventService *EventService
	logins []*Login
}

func NewGameService(server *Server) *GameService {
	gs := &GameService{
		mobService:      newMobService(),
		locationService: newLocationService(),
		roomService:     newRoomService(),
		server:          server,
	}
	gs.eventService = NewEventService(gs)
	return gs
}

func (gs *GameService) StartServer() {
	bufferWriter := make(chan *io.Buffer)
	go gs.server.Listen(bufferWriter)
	gs.ListenForNewBuffers(bufferWriter)
}

func (gs *GameService) ListenForNewBuffers(bufferWriter chan *io.Buffer) {
	for {
		select {
		case b := <- bufferWriter:
			gs.HandleBuffer(b)
			break
		}
	}
}

func (gs *GameService) HandleBuffer(b *io.Buffer) *output {
	mob, err := gs.findMobForClient(b.Client)
	if err != nil {
		gs.dummyLogin(b.Client)
		mob, _ = gs.findMobForClient(b.Client)
	}
	log.Printf("handling buffer: %s", b.ToString())
	input := io.NewInput(b.Client, strings.Split(b.Input, " "))
	action := findActionByCommand(input.GetCommand())
	output := action.mutator(input, gs.buildActionContext(mob, action, input), gs.eventService)
	b.Client.WritePrompt(output.messageToRequestCreator)
	if action.chainToCommand != "" {
		log.Printf("action %s chained to %s", action.command, action.chainToCommand)
		action = findActionByCommand(action.chainToCommand)
		output = action.mutator(
			io.NewInput(
				b.Client,
				[]string{string(action.command)}),
			gs.buildActionContext(mob, action, input),
			gs.eventService)
		b.Client.WritePrompt(output.messageToRequestCreator)
	}
	return output
}

func (gs *GameService) AddMobReset(mobReset *MobReset) {
	gs.mobService.addMobReset(mobReset)
}

func (gs *GameService) AddRoom(room *Room) {
	gs.roomService.addRoom(room)
}

func (gs *GameService) RespawnResets() {
	for _, mr := range gs.mobService.mobResets {
		mobsInRoom := gs.locationService.countMobsInRoom(mr.mob, mr.room)
		mobsInGame := gs.locationService.countMobsInGame(mr.mob)
		if mr.maxInRoom > mobsInRoom && mr.maxInGame > mobsInGame {
			gs.locationService.spawnMobToRoom(mr.mob, mr.room)
		}
	}
}

func (gs *GameService) CreateFixtures() {
	r1 := newRoom("Room 1", "You are in the first Room")
	r2 := newRoom("Room 2", "You are in the second Room")
	r3 := newRoom("Room 3", "You are in the third Room")

	r1.exits = append(r1.exits, newExit(r2, dSouth))
	r1.exits = append(r1.exits, newExit(r3, dWest))

	m := newMob("a test Mob", "A test Mob")

	r2.exits = append(r2.exits, newExit(r1, dNorth))
	r3.exits = append(r3.exits, newExit(r1, dEast))

	i1 := newItem("an item", "An item is here", []string{"item"})
	i2 := newItem("an item", "An item is here", []string{"item"})

	i1.position = held
	i2.position = held

	r1.items = append(r1.items, i1)
	r1.items = append(r1.items, i2)

	gs.AddRoom(r1)
	gs.AddRoom(r2)
	gs.AddRoom(r3)
	gs.AddMobReset(NewMobReset(m, r1, 1, 1))
	gs.RespawnResets()
}

func (gs *GameService) ChangeMobRoom(mob *Mob, room *Room) {
	gs.locationService.changeMobRoom(mob, room)
}

func (gs *GameService) findMobForClient(client *io.Client) (*Mob, error) {
	for _, l := range gs.logins {
		if l.client == client {
			return l.mob, nil
		}
	}
	return nil, errors.New("no mob found")
}

func (gs *GameService) dummyLogin(client *io.Client) {
	login := NewLogin(client, newMob("tester mctesterson", "A test Mob."))
	gs.logins = append(gs.logins, login)
	gs.locationService.spawnMobToRoom(login.mob, gs.roomService.rooms[0])
}

func (gs *GameService) getThingFromSyntax(syntax syntax, input *io.Input, ac *ActionContext) interface{} {
	switch syntax {
	case mobInRoomSyntax:
		return gs.locationService.findMobInRoom(input, ac.room)
	default:
		return nil
	}
}

func (gs *GameService) buildActionContext(mob *Mob, action *action, input *io.Input) *ActionContext {
	actionContext := &ActionContext{}
	actionContext.mob = mob
	actionContext.room = gs.locationService.getRoomForMob(mob)
	actionContext.hasDisposition = action.mobHasDisposition(mob)
	for _, s := range action.syntax {
		actionContext.results = append(actionContext.results, &context{
			syntax: s,
			thing: gs.getThingFromSyntax(s, input, actionContext),
		})
	}
	return actionContext
}
