package main

import (
	"image"
	"log"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
)

type RoverDriver struct {
	commandc chan command
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}       // current position
	direction := image.Point{X: 1, Y: 0} // current direction
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval) // make initial timer channel
	for {
		select {
		case c := <-r.commandc: // waits for the new comands
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
			}
			log.Printf("new direction %v", direction)
		case <-nextMove: // wait for the timer
			pos = pos.Add(direction)
			log.Printf("moved to %v", pos)        // outputs position
			nextMove = time.After(updateInterval) // makes the next timer channel
		}
	}
}

// turns the rover left
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// turns the rover right
func (r *RoverDriver) Right() {
	r.commandc <- right
}

func main() {
	r := NewRoverDriver()
	time.Sleep(2 * time.Second)
	r.Left()
	time.Sleep(2 * time.Second)
	r.Left()
	time.Sleep(2 * time.Second)
	r.Left()
	time.Sleep(2 * time.Second)
	r.Right()
	time.Sleep(2 * time.Second)
	r.Right()
	time.Sleep(2 * time.Second)
	r.Right()
	time.Sleep(2 * time.Second)

}
