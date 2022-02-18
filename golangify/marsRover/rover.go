package main

import (
	"image"
	"log"
	"time"
)

func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
	r.Stop()
	time.Sleep(3 * time.Second)
	r.Start()
	time.Sleep(3 * time.Second)
}

// RoverDriver ведет марсоход по поверхности Марса
type RoverDriver struct {
	commandc chan command
}

// NewRoverDriver начинает новый RoverDriver и возвращает его.
func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(2)
	stop  = command(3)
)

// drive отвечает за вождение марсохода. Ожидается
// что он будет начинаться в горутине.
func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	speed := 1
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			case stop:
				speed = 0
			case start:
				speed = 1
			}
			log.Printf("new direction %v; speed %d", direction, speed)
		case <-nextMove:
			pos = pos.Add(direction.Mul(speed))
			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

// Left поворачивает марсоход налево (90° потив часовой стрелки).
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// Right поворачивает марсоход направо (90° по часовой стрелке).
func (r *RoverDriver) Right() {
	r.commandc <- right
}

// Stop останавливает марсоход.
func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

// Start начинает движение марсохода.
func (r *RoverDriver) Start() {
	r.commandc <- start
}
