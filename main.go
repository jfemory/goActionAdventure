package main

import (
	//"gonum.org/v1/gonum/stat"
	//	"encoding/json"
	//"fmt"
	"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/ebitenutil"
	// "github.com/jfemory/goActionAdventure/obj"
	"image/color"
	"log"
)

/////////////////////Globals//////////////////////////////////
const windowTitle = "goAA"
const scale = 1
const windowWidth = 1024
const windowHeight = 768

var gameMode = 0

//////////////////////////////////////////////////////////////

//update is the main loop of the ebiten engine. Core loop is here.
func (state *gameState) update(screen *ebiten.Image) error {
	//gameMode 0 initializes assets
	if state.gameMode == 0 {
		if err := state.initAssets(); err != nil {
			log.Fatal(err)
		}
		state.gameMode = 1
	}
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//Do stuff goes here
	screen.Fill(color.RGBA{255, 0, 0, 255})

	globalCount(state)
	state.player.Mover()
	renderHelper(state, screen)
	state.drawSprite(screen)
	return nil
}

func main() {
	var state gameState
	err := state.initGameState()
	if err != nil {
		log.Fatal("failed to initialize game state. ")
	}
	if err := ebiten.Run(state.update, windowWidth, windowHeight, scale, windowTitle); err != nil {
		log.Fatal(err)
	}
}

func globalCount(state *gameState) {
	if state.globalCounter < 0xffffffffffffffff {
		state.globalCounter++
		//fmt.Println(state.globalCounter)
	}
}
