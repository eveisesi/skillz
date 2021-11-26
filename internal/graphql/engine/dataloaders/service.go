package dataloaders

import (
	"time"

	"github.com/eveisesi/skillz/internal/alliance"
	"github.com/eveisesi/skillz/internal/character"
	"github.com/eveisesi/skillz/internal/clone"
	"github.com/eveisesi/skillz/internal/corporation"
	"github.com/eveisesi/skillz/internal/graphql/engine/dataloaders/generated"
	"github.com/eveisesi/skillz/internal/universe"
)

type API interface {
	AllianceLoader() *generated.AllianceLoader
	CategoryLoader() *generated.CategoryLoader
	CharacterLoader() *generated.CharacterLoader
	CloneLoader() *generated.CloneLoader
	CorporationLoader() *generated.CorporationLoader
	ConstellationLoader() *generated.ConstellationLoader
	ImplantLoader() *generated.ImplantLoader
	GroupLoader() *generated.GroupLoader
	RegionLoader() *generated.RegionLoader
	SolarSystemLoader() *generated.SolarSystemLoader
	StationLoader() *generated.StationLoader
	StructureLoader() *generated.StructureLoader
	TypeLoader() *generated.TypeLoader
	TypeAttributeLoader() *generated.TypeAttributeLoader
}

type Service struct {
	wait  time.Duration
	batch int

	character   character.API
	corporation corporation.API
	alliance    alliance.API
	clone       clone.API
	universe    universe.API
}

func New(wait time.Duration, batch int, character character.API, corporation corporation.API, alliance alliance.API, clone clone.API, universe universe.API) *Service {
	return &Service{
		wait, batch, character, corporation, alliance, clone, universe,
	}
}
