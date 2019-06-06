package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
)

/////////////////////Globals//////////////////////////////////
const windowTitle = "goAA"
const scale = 1
const windowWidth = 1024
const windowHeight = 768

var gameMode = 0
var mc sprite

//////////////////////////////////////////////////////////////

type sprite struct {
	posX float64 //x position
	posY float64 //y position
	//	velocity vec  //Use this if you need velocity
	height int //height, used to determine collisions on multiheight maps
	img    *ebiten.Image
	//items
	//hold
}

//update is the main loop of the ebiten engine. Core loop is here.
func update(screen *ebiten.Image) error {
	//gameMode 0 initializes assets
	if gameMode == 0 {
		initAssets()
		gameMode = 1
	}
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//Do stuff goes here
	screen.Fill(color.RGBA{255, 0, 0, 255})
	mc.movement()
	mc.drawSprite(screen)
	return nil
}

func (s *sprite) movement() {
	//wasd controls
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		s.posX = s.posX + 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		s.posX = s.posX - 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		s.posY = s.posY + 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		s.posY = s.posY - 1
	}
}

func (s sprite) drawSprite(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(mc.posX, mc.posY)
	screen.DrawImage(mc.img, opts)
}

func main() {
	if err := ebiten.Run(update, windowWidth, windowHeight, scale, windowTitle); err != nil {
		log.Fatal(err)
	}
}

func initAssets() {
	var err error
	mc.img, _, err = ebitenutil.NewImageFromFile("gopherTop.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	mc.posX = windowWidth / 2
	mc.posY = windowHeight / 2
	world, _ := loadWorldMap("testMap.json")
	fmt.Println(world)
}
