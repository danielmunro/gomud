package model

import "github.com/jinzhu/gorm"

type Position string

const (
	LightPosition     Position = "light"
	Finger1Position   Position = "finger"
	Finger2Position   Position = "finger"
	Neck1Position     Position = "neck"
	Neck2Position     Position = "neck"
	TorsoPosition     Position = "torso"
	HeadPosition      Position = "head"
	LegsPosition      Position = "legs"
	FeetPosition      Position = "feet"
	HandsPosition     Position = "hands"
	ArmsPosition      Position = "arms"
	ShieldPosition    Position = "shield"
	BodyPosition      Position = "body"
	WaistPosition     Position = "waist"
	Wrist1Position    Position = "wrist"
	Wrist2Position    Position = "wrist"
	WieldPosition     Position = "wield"
	HeldPosition      Position = "held"
	FloatingPosition  Position = "floating"
	SecondaryPosition Position = "secondary"
)

type Item struct {
	gorm.Model
	name        string
	description string
	identifiers []string
	attributes  *Attributes
	Position    Position
}

func NewItem(name string, description string, identifiers []string) *Item {
	return &Item{
		name:        name,
		description: description,
		identifiers: identifiers,
	}
}

func (i *Item) String() string {
	return i.name
}
