package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tnaucoin/go-ecs-ebiten/config"
	"github.com/tnaucoin/go-ecs-ebiten/scenes"
	"image"
	"log"
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Game struct {
	bounds image.Rectangle
	scene  Scene
}

func NewGame() *Game {
	g := &Game{
		bounds: image.Rectangle{},
		scene:  &scenes.PlatformerScene{},
	}
	return g
}

func (g *Game) Update() error {
	g.scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.scene.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	g.bounds = image.Rect(0, 0, width, height)
	return width, height
}

func main() {
	ebiten.SetWindowSize(config.C.Width, config.C.Height)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
