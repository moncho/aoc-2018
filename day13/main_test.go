package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

const circularGrid = `/----\
|    |
|    |
\----/`

const testGrid = `/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `

const testGridNoCart = `/---\        
|   |  /----\
| /-+--+-\  |
| | |  | |  |
\-+-/  \-+--/
  \------/   `

const simpleTestGrid = `|
v
|
|
|
^
|
`

const multiCrashGrid = `/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`

const intersectingGrid = `/-----\
|     |
|  /--+--\
|  |  |  |
\--+--/  |
   |     |
   \-----/`

func Test_cart_forward(t *testing.T) {
	type fields struct {
		x         int
		y         int
		nextCross int
		face      rune
	}
	tests := []struct {
		name     string
		oldState *cart
		newState *cart
	}{
		{
			name: "facing down - move forward",
			oldState: &cart{
				x:         1,
				y:         1,
				direction: faceDown,
			},
			newState: &cart{
				x:         1,
				y:         2,
				direction: faceDown,
			},
		},
		{
			name: "facing left - move forward",
			oldState: &cart{
				x:         1,
				y:         1,
				direction: faceLeft,
			},
			newState: &cart{
				x:         0,
				y:         1,
				direction: faceLeft,
			},
		},
		{
			name: "facing up - move forward",
			oldState: &cart{
				x:         1,
				y:         1,
				direction: faceUp,
			},
			newState: &cart{
				x:         1,
				y:         0,
				direction: faceUp,
			},
		},
		{
			name: "facing right - move forward",
			oldState: &cart{
				x:         1,
				y:         1,
				direction: faceRight,
			},
			newState: &cart{
				x:         2,
				y:         1,
				direction: faceRight,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.oldState.move()
			if tt.oldState.x != tt.newState.x || tt.oldState.y != tt.newState.y || tt.oldState.direction != tt.newState.direction {
				t.Errorf("Unexpect cart state after forward(), want: %v, got: %v", tt.newState, tt.oldState)
			}
		})
	}
}

func Test_cart_turnLeft(t *testing.T) {
	type fields struct {
		x         int
		y         int
		nextCross int
		face      rune
	}
	tests := []struct {
		name     string
		oldState *cart
		newState *cart
	}{
		{
			name: "facing down - move turnLeft",
			oldState: &cart{
				x:         1,
				y:         1,
				direction: faceDown,
			},
			newState: &cart{
				x:         1,
				y:         1,
				direction: faceRight,
			},
		},
		{
			name: "facing left - move turnLeft",
			oldState: &cart{
				x:         1,
				y:         1,
				direction: faceLeft,
			},
			newState: &cart{
				x:         1,
				y:         1,
				direction: faceDown,
			},
		},
		{
			name: "facing up - move turnLeft",
			oldState: &cart{
				x:         1,
				y:         1,
				direction: faceUp,
			},
			newState: &cart{
				x:         1,
				y:         1,
				direction: faceLeft,
			},
		},
		{
			name: "facing right - move turnLeft",
			oldState: &cart{
				x:         1,
				y:         1,
				direction: faceRight,
			},
			newState: &cart{
				x:         1,
				y:         1,
				direction: faceUp,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.oldState.leftTurn()
			if tt.oldState.x != tt.newState.x || tt.oldState.y != tt.newState.y || tt.oldState.direction != tt.newState.direction {
				t.Errorf("Unexpect cart state after turnLeft(), want: %v, got: %v", tt.newState, tt.oldState)
			}
		})
	}
}

func Test_cart_turnRight(t *testing.T) {
	type fields struct {
		x         int
		y         int
		nextCross int
		face      rune
	}
	tests := []struct {
		name     string
		oldState *cart
		newState *cart
	}{
		{
			name: "facing down - move turnRight",
			oldState: &cart{
				x:         1,
				y:         1,
				direction: faceDown,
			},
			newState: &cart{
				x:         1,
				y:         1,
				direction: faceLeft,
			},
		},
		{
			name: "facing left - move turnRight",
			oldState: &cart{
				x:         1,
				y:         1,
				direction: faceLeft,
			},
			newState: &cart{
				x:         1,
				y:         1,
				direction: faceUp,
			},
		},
		{
			name: "facing up - move turnRight",
			oldState: &cart{
				x:         1,
				y:         1,
				direction: faceUp,
			},
			newState: &cart{
				x:         1,
				y:         1,
				direction: faceRight,
			},
		},
		{
			name: "facing right - move turnRight",
			oldState: &cart{
				x:         1,
				y:         1,
				direction: faceRight,
			},
			newState: &cart{
				x:         1,
				y:         1,
				direction: faceDown,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.oldState.rightTurn()
			if tt.oldState.x != tt.newState.x || tt.oldState.y != tt.newState.y || tt.oldState.direction != tt.newState.direction {
				t.Errorf("Unexpect cart state after turnRight(), want: %v, got: %v", tt.newState, tt.oldState)
			}
		})
	}
}

func Test_newGrid(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name          string
		args          args
		expectedGrid  string
		expectedCarts []*cart
	}{
		{
			"Generate a grid representation",
			args{
				r: strings.NewReader(testGrid),
			},
			testGridNoCart,
			[]*cart{
				&cart{
					x:         2,
					y:         0,
					direction: faceRight,
					working:   true,
				},
				&cart{
					x:         9,
					y:         3,
					direction: faceDown,
					working:   true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid, carts := newGrid(tt.args.r)
			var s string
			for _, line := range grid {
				s += string(line)
				s += "\n"
			}
			//Remove the last newline char
			s = s[:len(s)-1]
			if s != tt.expectedGrid {
				t.Errorf("newGrid() got grid = \n%v, want \n%v", s, tt.expectedGrid)
			}
			if !reflect.DeepEqual(carts, tt.expectedCarts) {
				t.Errorf("newGrid() got carts = %v, want %v", carts, tt.expectedCarts)
			}
		})
	}
}

func Test_tick(t *testing.T) {
	type args struct {
		ticks int
		c     *cart
		grid  string
	}
	tests := []struct {
		name string
		args args
		want *cart
	}{
		{
			"cart facing right on horizontal line - moves straight",
			args{
				ticks: 1,
				c: &cart{
					x:         2,
					y:         0,
					direction: faceRight,
				},
				grid: testGrid,
			},
			&cart{
				x:         3,
				y:         0,
				direction: faceRight,
			},
		},
		{
			"cart facing right on horizontal line - moves straight",
			args{
				ticks: 1,
				c: &cart{
					x:         3,
					y:         0,
					direction: faceRight,
				},
				grid: testGrid,
			},
			&cart{
				x:         4,
				y:         0,
				direction: faceRight,
			},
		},
		{
			"cart facing right on right turn line - moves down",
			args{
				ticks: 1,
				c: &cart{
					x:         4,
					y:         0,
					direction: faceRight,
				},
				grid: testGrid,
			},
			&cart{
				x:         4,
				y:         1,
				direction: faceDown,
			},
		},
		{
			"cart facing left on right turn line - moves up",
			args{
				ticks: 1,
				c: &cart{
					x:         0,
					y:         4,
					direction: faceLeft,
				},
				grid: testGrid,
			},
			&cart{
				x:         0,
				y:         3,
				direction: faceUp,
			},
		},
		{
			"cart facing down at first crossing - turns left",
			args{
				ticks: 1,
				c: &cart{
					x:         2,
					y:         4,
					nextCross: 0,
					direction: faceDown,
				},
				grid: testGrid,
			},
			&cart{
				x:         3,
				y:         4,
				nextCross: 1,
				direction: faceRight,
			},
		},
		{
			"cart facing down at second crossing - goes straight",
			args{
				ticks: 1,
				c: &cart{
					x:         2,
					y:         4,
					nextCross: 1,
					direction: faceDown,
				},
				grid: testGrid,
			},
			&cart{
				x:         2,
				y:         5,
				nextCross: 2,
				direction: faceDown,
			},
		},
		{
			"cart facing down at third crossing - turns right",
			args{
				ticks: 1,
				c: &cart{
					x:         2,
					y:         4,
					nextCross: 2,
					direction: faceDown,
				},
				grid: testGrid,
			},
			&cart{
				x:         1,
				y:         4,
				nextCross: 0,
				direction: faceLeft,
			},
		},
		{
			"cart facing up at first crossing - turns left",
			args{
				ticks: 1,
				c: &cart{
					x:         2,
					y:         4,
					nextCross: 0,
					direction: faceUp,
				},
				grid: testGrid,
			},
			&cart{
				x:         1,
				y:         4,
				nextCross: 1,
				direction: faceLeft,
			},
		},
		{
			"cart facing up at second crossing - goes straight",
			args{
				ticks: 1,
				c: &cart{
					x:         2,
					y:         4,
					nextCross: 1,
					direction: faceUp,
				},
				grid: testGrid,
			},
			&cart{
				x:         2,
				y:         3,
				nextCross: 2,
				direction: faceUp,
			},
		},
		{
			"cart facing up at third crossing - turns right",
			args{
				ticks: 1,
				c: &cart{
					x:         2,
					y:         4,
					nextCross: 2,
					direction: faceUp,
				},
				grid: testGrid,
			},
			&cart{
				x:         3,
				y:         4,
				nextCross: 0,
				direction: faceRight,
			},
		},
		{
			"cart moves clockwise on a circular grid",
			args{
				ticks: 16,
				c: &cart{
					x:         1,
					y:         0,
					nextCross: 0,
					direction: faceRight,
				},
				grid: circularGrid,
			},
			&cart{
				x:         1,
				y:         0,
				nextCross: 0,
				direction: faceRight,
			},
		},
		{
			"cart moves counterclockwise on a circular grid",
			args{
				ticks: 16,
				c: &cart{
					x:         1,
					y:         0,
					nextCross: 0,
					direction: faceLeft,
				},
				grid: circularGrid,
			},
			&cart{
				x:         1,
				y:         0,
				nextCross: 0,
				direction: faceLeft,
			},
		},
		{
			"cart in a grid with two intersects",
			args{
				ticks: 60,
				c: &cart{
					x:         1,
					y:         0,
					nextCross: 0,
					direction: faceRight,
				},
				grid: intersectingGrid,
			},
			&cart{
				x:         1,
				y:         0,
				nextCross: 0,
				direction: faceRight,
			},
		},
		{
			"cart in a grid with two intersects",
			args{
				ticks: 60,
				c: &cart{
					x:         1,
					y:         0,
					nextCross: 0,
					direction: faceLeft,
				},
				grid: intersectingGrid,
			},
			&cart{
				x:         1,
				y:         0,
				nextCross: 0,
				direction: faceLeft,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid, _ := newGrid(strings.NewReader(tt.args.grid))
			for i := 0; i < tt.args.ticks; i++ {
				tick(tt.args.c, grid)
			}
			if !reflect.DeepEqual(tt.args.c, tt.want) {
				t.Errorf("after a tick() got cart = %v, want %v", tt.args.c, tt.want)
			}
		})
	}
}

func Test_runSimulation(t *testing.T) {

	tests := []struct {
		name  string
		grid  string
		want  int
		want1 int
	}{
		{
			"test grid simulation",
			testGrid,
			7,
			3,
		},
		{
			"simple grid simulation",
			simpleTestGrid,
			0,
			3,
		},
		{
			"multicrash grid simulation",
			multiCrashGrid,
			2,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid, carts := newGrid(strings.NewReader(tt.grid))
			broken, _ := runSimulation(grid, carts)
			if broken[0].x != tt.want {
				t.Errorf("runSimulation() got = %v, want %v", broken[0].x, tt.want)
			}
			if broken[0].y != tt.want1 {
				t.Errorf("runSimulation() got1 = %v, want %v", broken[0].y, tt.want1)
			}
		})
	}
}
