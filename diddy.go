package main

import rl "github.com/gen2brain/raylib-go/raylib"

import (
	"math/rand"
	"fmt"
	"math"
)
type DiddyBlxd struct {
	position rl.Vector2
}

func hoveringBlxd(blxd DiddyBlxd) bool {
	mouse := rl.GetMousePosition()
	width := ((float32)(diddy.Width) * scale)
	height := ((float32)(diddy.Height) * scale)

	left := blxd.position.X - width/2
	right := blxd.position.X + width/2
	top := blxd.position.Y - height/2
	bottom := blxd.position.Y + height/2

	return mouse.X >= left && mouse.X <= right &&
		mouse.Y >= top  && mouse.Y <= bottom
}

func gameLoop() {
	if len(blxds) < 2 || rand.Int31n(80) == 19 {
		theta := rand.Float64() * 2 * math.Pi 
		const length = 500;
		blxds = append(blxds, DiddyBlxd{ 
			position: rl.Vector2{
				X: (float32)(math.Cos(theta) * length + 400),
				Y: (float32)(math.Sin(theta) * length) + 400,
			},
		});
	}
	
	for i := len(blxds) - 1; i >= 0; i-- {
		if(hoveringBlxd(blxds[i]) && rl.IsMouseButtonPressed(rl.MouseButtonLeft)) {
			blxds = append(blxds[:i], blxds[i+1:]...)
			score++
			continue
		}

		const maxDistance = 2
		toX := 400 - blxds[i].position.X
		toY := 400 - blxds[i].position.Y

		sqDist := toX * toX + toY * toY
		if sqDist == 0 || maxDistance * maxDistance >= sqDist {
			playing = false;
			return; 
		}

		dist := (float32)(math.Sqrt((float64)(sqDist)))
		
		blxds[i].position.X += toX / dist * maxDistance;
		blxds[i].position.Y += toY / dist * maxDistance;
		
		
	}
}

func draw() {
	rl.ClearBackground(rl.NewColor(100, 149, 237, 255))

	centrePositionB := rl.Vector2{ X: 400 - ((float32)(baksa.Width) * scale / 2), Y: 400 - ((float32)(baksa.Height) * scale / 2) };
	rl.DrawTextureEx(baksa, centrePositionB, 0, scale, rl.White);

	for idx, blxd := range blxds {
		_ = idx
		centrePosition := rl.Vector2{ X: blxd.position.X - ((float32)(diddy.Width) * scale / 2), Y: blxd.position.Y - ((float32)(diddy.Height) * scale / 2) };
		rl.DrawTextureEx(diddy, centrePosition, 0, scale, rl.White);
	}

	words := fmt.Sprintf("Blxds killed :skull:: %d", score)	
	rl.DrawText(words, 290, 10, 20, rl.Black)

	if !playing {
		rl.DrawRectangle(0, 0, 800, 800, rl.NewColor(0,0,0,127));
		rl.DrawText("bro got diddyd :sob:", 240, 400-32, 32, rl.Red);
	}

}

const scale = 0.3
var diddy rl.Texture2D
var baksa rl.Texture2D
var blxds = []DiddyBlxd{} 
var playing = true
var score = 0
func main() {
	rl.InitWindow(800, 800, "freaking diddy simulator bro💀")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	diddyImg := rl.LoadImage("diddy.png")
	defer rl.UnloadImage(diddyImg);
	diddy = rl.LoadTextureFromImage(diddyImg);

	baksaImg := rl.LoadImage("baksa.png")
	defer rl.UnloadImage(baksaImg);
	baksa = rl.LoadTextureFromImage(baksaImg);


	for !rl.WindowShouldClose() {
		
		if playing {
			gameLoop();
		}

		rl.BeginDrawing()
		draw();
		rl.EndDrawing()
	}
}
