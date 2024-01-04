package archetypes

import (
	"github.com/tnaucoin/go-ecs-ebiten/components"
	"github.com/tnaucoin/go-ecs-ebiten/layers"
	"github.com/tnaucoin/go-ecs-ebiten/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

var (
	Space = newArchetype(
		components.Space,
	)
	Wall = newArchetype(tags.Wall, components.Object)
)

type archetype struct {
	components []donburi.IComponentType
}

func newArchetype(cs ...donburi.IComponentType) *archetype {
	return &archetype{
		components: cs,
	}
}

func (a *archetype) Spawn(ecs *ecs.ECS, cs ...donburi.IComponentType) *donburi.Entry {
	e := ecs.World.Entry(ecs.Create(layers.Default, append(a.components, cs...)...,
	))
	return e
}
