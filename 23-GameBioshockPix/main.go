package main

import (
	//"fmt"
	"image/color"
	"log"
	//"math/rand"
	//"sync"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

)
type Shooter struct{}

type BigDaddy struct {
	ecs.BasicEntity
	ControlComponent
	common.RenderComponent
	common.SpaceComponent
}

func (pong *Shooter) Preload() {
	err := engo.Files.Load("BigDaddy.png")
	if err != nil {
		log.Println(err)
	}
}

func (pong *Shooter) Setup(w *ecs.World) {
	common.SetBackground(color.Black)

	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&ControlSystem{})

	paddleTexture, err := common.LoadedSprite("BigDaddy.png")
	if err != nil {
		log.Println(err)
	}

	engo.Input.RegisterAxis("arrows", engo.AxisKeyPair{engo.ArrowUp, engo.ArrowDown},engo.AxisKeyPair{engo.ArrowLeft,engo.ArrowRight})

	schemes := []string{"arrows"}

	for i := 0; i < 1; i++ {

		paddle := BigDaddy{BasicEntity: ecs.NewBasic()}
		paddle.RenderComponent = common.RenderComponent{
			Drawable: paddleTexture,
			Scale:    engo.Point{2, 2},
		}

		x := float32(0)

		if i != 0 {
			x = engo.GameWidth() - 16
		}

		paddle.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{x, (engo.GameHeight() - paddleTexture.Height()) / 1},
			Width:    paddle.RenderComponent.Scale.X * paddleTexture.Width(),
			Height:   paddle.RenderComponent.Scale.Y * paddleTexture.Height(),
		}

		paddle.ControlComponent = ControlComponent{Scheme: schemes[i]}

		// Add our entity to the appropriate systems

		for _, system := range w.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(&paddle.BasicEntity, &paddle.RenderComponent, &paddle.SpaceComponent)
			case *ControlSystem:
				sys.Add(&paddle.BasicEntity, &paddle.ControlComponent, &paddle.SpaceComponent)
			}
		}

	}
}

func (*Shooter) Type() string { return "Shooter" }

type ControlComponent struct {
	Scheme string


}
type controlEntity struct {
	*ecs.BasicEntity
	*ControlComponent
	*common.SpaceComponent
}

type ControlSystem struct {
	entities []controlEntity

	mouseTrackerBasic ecs.BasicEntity
	mouseTrackerMouse common.MouseComponent
}

func (c *ControlSystem) New(w *ecs.World) {

	c.mouseTrackerBasic = ecs.NewBasic()
	c.mouseTrackerMouse.Track = true

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&c.mouseTrackerBasic, &c.mouseTrackerMouse, nil, nil)
		}
	}
}
func (c *ControlSystem) Add(basic *ecs.BasicEntity, control *ControlComponent, space *common.SpaceComponent) {

	c.entities = append(c.entities, controlEntity{basic, control, space})

}
func (c *ControlSystem) Remove(basic ecs.BasicEntity) {

	delete := -1
	for index, e := range c.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		c.entities = append(c.entities[:delete], c.entities[delete+1:]...)
	}

}
func (c *ControlSystem) Update(dt float32) {

	for _, e := range c.entities {

		speed := engo.GameWidth() * dt

		//Esto es para los movimientos.
4
		vert := engo.Input.Axis(e.ControlComponent.Scheme)
		e.SpaceComponent.Position.Y += speed * vert.Value()

		var moveThisOne bool

		//Esto reinicia el movimiento cuando paso el mouse.

		//if engo.Input.Mouse.X > engo.WindowWidth()/2 && e.ControlComponent.Scheme == "arrows" {
		//	moveThisOne = true
		// if engo.Input.Mouse.X < engo.WindowWidth()/2 && e.ControlComponent.Scheme == "wasd" {
		//	moveThisOne = true
		//}

		if moveThisOne {
			e.SpaceComponent.Position.Y = c.mouseTrackerMouse.MouseY - e.SpaceComponent.Height/2
		}

		// Esto permite que el personaje no se salga de la pantalla, arriba o abajo.

		if (e.SpaceComponent.Height + e.SpaceComponent.Position.Y) > engo.GameHeight() {
			e.SpaceComponent.Position.Y = engo.GameHeight() - e.SpaceComponent.Height

		} else if e.SpaceComponent.Position.Y < 0 {
			e.SpaceComponent.Position.Y = 0
		}
	}




}

func main() {
	opts := engo.RunOptions{
		Title:         "Bioshock Demo",
		Width:         600,
		Height:        600,
		ScaleOnResize: true,
	}
	engo.Run(opts, &Shooter{})
}