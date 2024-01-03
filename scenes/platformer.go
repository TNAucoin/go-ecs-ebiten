package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tnaucoin/go-ecs-ebiten/factory"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"image/color"
	"sync"
)

type PlatformerScene struct {
	ecs  *ecs.ECS
	once sync.Once
}

func (ps *PlatformerScene) Update() {
	ps.once.Do(ps.configure)
	ps.ecs.Update()
}

func (ps *PlatformerScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 40, 255})
	ps.ecs.Draw(screen)
}

func (ps *PlatformerScene) configure() {
	pecs := ecs.NewECS(donburi.NewWorld())
	// Add Systems
	// Add Rendering Systems
	ps.ecs = pecs

	// Define the world space. Space here is a grid of 16x16 cells.
	// Each cell will contain 0 or more objects, collisions are found by
	// checking the cell, and if it contains objects.
	space := factory.CreateSpace(ps.ecs)
}
