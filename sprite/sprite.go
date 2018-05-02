package main

import (
	"fmt"
	"github.com/tubelz/macaw"
	"github.com/tubelz/macaw/entity"
	"github.com/tubelz/macaw/input"
	"github.com/tubelz/macaw/system"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	var err error
	// Initialize SDL_IMG. We don't need to initialize TTF and MIX
	err = macaw.Initialize(true, false, false)
	if err != nil {
		fmt.Println("Macaw could not initialize")
	}
	defer macaw.Quit()

  input := &input.Manager{}
	em := &entity.Manager{}
  // render is the only system we need here
	render := &system.RenderSystem{Name: "render system", Window: macaw.Window, EntityManager: em}
	render.Init()
	initializeEntities(em, render)

	render.SetCamera(em.Get(1))
	gameLoop := initializeGameLoop(render, input)
	gameLoop.Run()
}

// initializeEntities will create the sprite and camera
func initializeEntities(em *entity.Manager, render *system.RenderSystem) {
	objSpritesheet := &entity.Spritesheet{}
	objSpritesheet.Init(render.Renderer, "assets/macaw.png")
	// Load sprites from spritesheet
	macawCrop := &sdl.Rect{0, 0, 200, 200}
	sprite := objSpritesheet.LoadSprite(macawCrop)
	obj := em.Create()
	obj.AddComponent("render", &sprite)
	obj.AddComponent("position", &entity.PositionComponent{&sdl.Point{100, 100}})

	camera := em.Create()
	camera.AddComponent("position", &entity.PositionComponent{&sdl.Point{0, 0}})
	camera.AddComponent("camera", &entity.CameraComponent{
		ViewportSize: sdl.Point{800, 600},
		WorldSize:    sdl.Point{1145, 600},
	})
}

func initializeGameLoop(render *system.RenderSystem, im *input.Manager) *macaw.GameLoop {
	gameLoop := &macaw.GameLoop{InputManager: im}
	sceneGame := &macaw.Scene{Name: "example"}
	sceneGame.AddRenderSystem(render)
	gameLoop.AddScene(sceneGame)

	return gameLoop
}
