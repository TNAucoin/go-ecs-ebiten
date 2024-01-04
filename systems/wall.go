package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	dresolv "github.com/tnaucoin/go-ecs-ebiten/resolv"
	"github.com/tnaucoin/go-ecs-ebiten/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"image/color"
)

func DrawWall(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.Wall.Each(ecs.World, func(e *donburi.Entry) {
		o := dresolv.GetObject(e)
		drawColor := color.RGBA{R: 60, G: 60, B: 60, A: 255}
		vector.DrawFilledRect(screen, float32(o.X), float32(o.Y), float32(o.W), float32(o.H), drawColor, false)
	})
}
