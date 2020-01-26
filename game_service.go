package gomud

import (
	"log"
	"strings"
)

type GameService struct {
	mobService *MobService
	locationService *LocationService
	roomService *RoomService
	server *Server
	buffers []*Buffer
}

func NewGameService(server *Server) *GameService {
	return &GameService{
		mobService:      newMobService(),
		locationService: newLocationService(),
		roomService:     newRoomService(),
		server:          server,
	}
}

func (gs *GameService) StartServer() {
	bufferWriter := make(chan *Buffer)
	go gs.server.Listen(bufferWriter)
	gs.ListenForNewBuffers(bufferWriter)
}

func (gs *GameService) ListenForNewBuffers(bufferWriter chan *Buffer) {
	for {
		select {
		case b := <- bufferWriter:
			log.Printf("handling buffer: %s", b.ToString())
			room := gs.locationService.getRoomForMob(b.client.mob)
			input := newInput(b.client, room, strings.Split(b.input, " "))
			action := findActionByCommand(input.getCommand())
			actionContext := gs.buildActionContext(b.client.mob, action, input)
			output := action.mutator(input, actionContext)
			b.client.writePrompt(output.messageToRequestCreator)
			break
		}
	}
}

func (gs *GameService) AddMobReset(mobReset *MobReset) {
	gs.mobService.addMobReset(mobReset)
}

func (gs *GameService) AddRoom(room *room) {
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
	r1 := newRoom("Room 1", "You are in the first room")
	r2 := newRoom("Room 2", "You are in the second room")
	r3 := newRoom("Room 3", "You are in the third room")

	r1.exits = append(r1.exits, newExit(r2, dSouth))
	r1.exits = append(r1.exits, newExit(r3, dWest))

	m := newMob("a test mob", "A test mob")

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

func (gs *GameService) getThingFromSyntax(syntax syntax, input *input) interface{} {
	switch syntax {
	case mobInRoomSyntax:
		return gs.locationService.findMobInRoom(input)
	default:
		return nil
	}
}

func (gs *GameService) buildActionContext(mob *mob, action *action, input *input) *ActionContext {
	actionContext := &ActionContext{}
	actionContext.hasDisposition = action.mobHasDisposition(mob)
	for _, s := range action.syntax {
		actionContext.results = append(actionContext.results, &context{
			syntax: s,
			thing: gs.getThingFromSyntax(s, input),
		})
	}
	return actionContext
}
