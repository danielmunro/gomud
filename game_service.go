package gomud

import (
	"errors"
	"fmt"
	"github.com/danielmunro/gomud/io"
	"github.com/danielmunro/gomud/message"
	"github.com/danielmunro/gomud/model"
	"log"
	"time"
)

type GameService struct {
	mobService      *MobService
	locationService *LocationService
	roomService     *RoomService
	actionService   *ActionService
	server          *io.Server
	buffers         []*io.Buffer
	eventService    *EventService
	logins          []*Login
}

func NewGameService(server *io.Server) *GameService {
	ms := NewMobService()
	ls := newLocationService()
	es := NewEventService(ls, ms)
	as := newActionService(ls, es)
	gs := &GameService{
		mobService:      ms,
		locationService: ls,
		roomService:     newRoomService(),
		server:          server,
		eventService:    es,
		actionService:   as,
	}
	return gs
}

func (gs *GameService) StartServer() {
	bufferWriter := make(chan *io.Buffer)
	go gs.server.Listen(bufferWriter)
	go gs.StartPulses()
	gs.ListenForNewBuffers(bufferWriter)
}

func (gs *GameService) StartPulses() {
	for {
		gs.eventService.Publish(NewSystemEvent(PulseEventType))
		time.Sleep(1 * time.Second)
	}
}

func (gs *GameService) ListenForNewBuffers(bufferWriter chan *io.Buffer) {
	for {
		select {
		case b := <-bufferWriter:
			gs.HandleBuffer(b)
			break
		}
	}
}

func (gs *GameService) HandleBuffer(b *io.Buffer) *io.Output {
	mob := gs.findMobOrLogin(b.Client)
	return gs.processBuffer(b, mob)
}

func (gs *GameService) AddMobReset(mobReset *model.MobReset) {
	gs.mobService.addMobReset(mobReset)
}

func (gs *GameService) AddRoom(room *model.Room) {
	gs.roomService.addRoom(room)
}

func (gs *GameService) RespawnResets() {
	for _, mr := range gs.mobService.mobResets {
		mobsInRoom := gs.locationService.countMobsInRoom(mr.Mob, mr.Room)
		mobsInGame := gs.locationService.countMobsInGame(mr.Mob)
		if mr.MaxInRoom > mobsInRoom && mr.MaxInGame > mobsInGame {
			mob := mr.Mob
			gs.locationService.spawnMobToRoom(mob, mr.Room)
		}
	}
}

func (gs *GameService) CreateFixtures() {
	r1 := model.NewRoom("Room 1", "You are in the first Room")
	r2 := model.NewRoom("Room 2", "You are in the second Room")
	r3 := model.NewRoom("Room 3", "You are in the third Room")

	r1.Exits = append(r1.Exits, model.NewExit(r2, model.SouthDirection))
	r1.Exits = append(r1.Exits, model.NewExit(r3, model.WestDirection))

	m := model.NewMob("a test Mob", "A test Mob")

	r2.Exits = append(r2.Exits, model.NewExit(r1, model.NorthDirection))
	r3.Exits = append(r3.Exits, model.NewExit(r1, model.EastDirection))

	i1 := model.NewItem("an item", "An item is here", []string{"item"})
	i2 := model.NewItem("an item", "An item is here", []string{"item"})
	i3 := model.NewEquipment("a cowboy hat", "A sturdy cowboy hat.", []string{"cowboy", "hat"}, model.HeadPosition)
	i4 := model.NewEquipment("a baseball cap", "A worn baseball cap.", []string{"baseball", "cap"}, model.HeadPosition)

	i1.Position = model.HeadPosition
	i2.Position = model.HeldPosition

	r1.Items = append(r1.Items, i1)
	r1.Items = append(r1.Items, i2)
	r1.Items = append(r1.Items, i3)
	r1.Items = append(r1.Items, i4)

	gs.AddRoom(r1)
	gs.AddRoom(r2)
	gs.AddRoom(r3)
	gs.AddMobReset(model.NewMobReset(m, r1, 1, 1))
	gs.RespawnResets()
}

func (gs *GameService) ChangeMobRoom(mob *model.Mob, room *model.Room) {
	gs.locationService.changeMobRoom(mob, room)
}

func (gs *GameService) processBuffer(b *io.Buffer, mob *model.Mob) *io.Output {
	log.Printf("handling Buffer: %s", b.ToString())
	action := findActionByCommand(b.GetCommand())
	context := gs.buildActionContext(mob, action, b)
	ctx := context.findErrorContext()
	if ctx != nil {
		log.Printf("context error: %s, %s", ctx.syntax, ctx.error.Error())
		output := io.NewOutputToRequestCreator(b, io.ErrorStatus, ctx.error.Error())
		b.Client.WritePrompt(output.MessageToRequestCreator)
		return output
	}
	output := action.mutator(context, gs.actionService)
	b.Client.WritePrompt(output.MessageToRequestCreator)
	if action.chainToCommand != "" {
		log.Printf("action %s chained to %s", action.command, action.chainToCommand)
		return gs.processBuffer(io.NewBuffer(b.Client, string(action.chainToCommand)), mob)
	}
	return output
}

func (gs *GameService) findMobOrLogin(client *io.Client) *model.Mob {
	mob, err := gs.findMobForClient(client)
	if err != nil {
		gs.dummyLogin(client)
		mob, _ = gs.findMobForClient(client)
	}
	return mob
}

func (gs *GameService) findMobForClient(client *io.Client) (*model.Mob, error) {
	for _, l := range gs.logins {
		if l.client == client {
			return l.mob, nil
		}
	}
	return nil, errors.New("no mob found")
}

func (gs *GameService) dummyLogin(client *io.Client) {
	login := NewLogin(client, model.NewMob("tester mctesterson", "A test Mob."))
	gs.logins = append(gs.logins, login)
	gs.locationService.spawnMobToRoom(login.mob, gs.roomService.rooms[0])
}

func (gs *GameService) getThingFromSyntax(syntax syntax, ac *ActionContext) (interface{}, error) {
	buffer := ac.buffer
	switch syntax {
	case commandSyntax:
		return string(syntax), nil
	case mobInRoomSyntax:
		return gs.locationService.findMobInRoom(buffer, ac.room)
	case merchantInRoomSyntax:
		for _, m := range gs.locationService.getMobsInRoom(ac.room) {
			if m.IsMerchant() {
				return m, nil
			}
		}
		return nil, errors.New("no merchant is here")
	case itemInTargetInventorySyntax:
		mob := ac.getFirstMob()
		return mob.FindItem(buffer)
	case itemInInventorySyntax:
		return ac.mob.FindItem(buffer)
	case itemInRoomSyntax:
		return ac.room.FindItem(buffer)
	case itemEquippedSyntax:
		return ac.mob.FindEquipped(buffer)
	case exitDirectionSyntax:
		return ac.room.GetExitFromDirection(string(buffer.GetCommand()))
	default:
		return nil, errors.New("syntax not implemented")
	}
}

func (gs *GameService) buildActionContext(mob *model.Mob, action *Action, buffer *io.Buffer) *ActionContext {
	actionContext := &ActionContext{}
	actionContext.mob = mob
	actionContext.room = gs.locationService.getRoomForMob(mob)
	actionContext.hasDisposition = action.mobHasDisposition(mob)
	actionContext.buffer = buffer
	if !actionContext.hasDisposition {
		actionContext.results = append(actionContext.results, newContext(
			commandSyntax,
			nil,
			errors.New(fmt.Sprintf(message.ErrorDispositionMismatch, string(action.dispositions[0])))))
		return actionContext
	}
	for _, syntax := range action.syntax {
		thing, err := gs.getThingFromSyntax(syntax, actionContext)
		actionContext.results = append(actionContext.results, newContext(syntax, thing, err))
	}
	return actionContext
}
