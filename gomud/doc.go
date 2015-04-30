/*
	gomud contains the code necessary to run a simple client-server MUD
	(multi-user dungeon) game over telnet. The files break down as follows:
		areas.go -

		affect.go -

		attributes.go - Attributes defines the AC, Vitals, Stats, and Attributes structs
					that encapsulate the numerical characteristics of a Mob.

		client.go - Defines the Client structure and its methods.

		item.go - Defines the Item interface and Equipment & Equipped structs with
					their methods.

		message.go - Defines the Message structure and its methods.

		mob.go - Defines the Skill and Mob structs and the Disposition type;
					defines methods for Mob.

		race.go - Defines the Race struct and six default races.

		realm.go - Defines the Room struct and Direction type and defines methods
					for their usage.

		server.go - Defines the Server structure and its methods.
*/
package gomud
