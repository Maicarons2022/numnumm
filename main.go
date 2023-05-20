package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/audio"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
    screenWidth  = 640
    screenHeight = 480
)

type Game struct {
    scene *Scene
}

type Scene struct {
    background *ebiten.Image
    music      *audio.Player
}

type Object struct {
    image  *ebiten.Image
    x      float64
    y      float64
}

func main() {
    ebiten.SetWindowSize(screenWidth, screenHeight)
    ebiten.SetWindowTitle("My Game")

    game := &Game{
        scene: &Scene{},
    }

    game.scene.background, _, _ = ebitenutil.NewImageFromFile("background.png")
    //game.scene.music, _, _ = audio.NewPlayerFromFile("music.mp3")

    objectImage, _, _ := ebitenutil.NewImageFromFile("object.png")
    object := &Object{
        image: objectImage,
        x:     0,
        y:     0,
    }

    game.scene.AddObject(object)

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 640, 480
}

func (g *Game) Update() error {
    if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
        g.scene.music.Play()
    }

    g.scene.Update()

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    g.scene.Draw(screen)
}

func (s *Scene) AddObject(object *Object) {
    // Add object to scene
}

func (s *Scene) Update() {
    // Update scene objects
}

func (s *Scene) Draw(screen *ebiten.Image) {
    // Draw scene objects and background
}