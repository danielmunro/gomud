package gomud

import "github.com/danielmunro/gomud/model"

func CreateFixtures(gs *GameService) {
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

	i5 := model.NewEquipment("a wooden shield", "A wooden practice shield is here", []string{"wooden", "shield", "practice"}, model.ShieldPosition)
	i5.Value = 5
	i5.IsStoreItem = true
	merch := model.NewMob("a merchant", "A human merchant")
	merch.MakeMerchant()
	merch.Items = append(merch.Items, i5)
	merch.Gold = 1000

	gs.AddRoom(r1)
	gs.AddRoom(r2)
	gs.AddRoom(r3)
	gs.AddMobReset(model.NewMobReset(m, r1, 1, 1))
	gs.AddMobReset(model.NewMobReset(merch, r3, 1, 1))
}
