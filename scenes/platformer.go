package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/tnaucoin/go-ecs-ebiten/config"
	"github.com/tnaucoin/go-ecs-ebiten/factory"
	"github.com/tnaucoin/go-ecs-ebiten/layers"
	dresolv "github.com/tnaucoin/go-ecs-ebiten/resolv"
	"github.com/tnaucoin/go-ecs-ebiten/systems"
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
	pecs.AddRenderer(layers.Default, systems.DrawWall)
	ps.ecs = pecs
	gw, gh := float64(config.C.Width), float64(config.C.Height)
	// Define the world space. Space here is a grid of 16x16 cells.
	// Each cell will contain 0 or more objects, collisions are found by
	// checking the cell, and if it contains objects.
	space := factory.CreateSpace(ps.ecs)

	dresolv.Add(space,
		factory.CreateWall(ps.ecs, resolv.NewObject(0, 0, 16, gh, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(gw-16, 0, 16, gh, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(0, 0, gw, 16, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(0, gh-24, gw, 32, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(160, gh-56, 160, 32, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(320, 64, 32, 160, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(64, 128, 16, 160, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(gw-128, 64, 128, 16, "solid")),
		factory.CreateWall(ps.ecs, resolv.NewObject(gw-128, gh-88, 128, 16, "solid")),
	)

}
