package main 
import (
    "fmt"
    "math/rand"
)

const blocked = 0
const passage = 1

type cell_coords struct {
    x, y int
}

type Cell struct {
    coord cell_coords
    state int
}

type Grid struct {
    cells [][]Cell
    height int
    width int
    frontier []cell_coords
    connected []cell_coords
}


func main() {
    // 1. A Grid consists of a 2 dimensional array of cells.
    // 2. A Cell has 2 states: Blocked or Passage.

    // 3. Start with a Grid full of Cells in state Blocked.
	var height int = 10
	var width int = 10
	var blocked string = "#"
	//passage := " "
	//var frontier_directions = [][]int{ {-2, 0}, {2, 0}, {0, -2}, {0, 2} }

    var grid = [10][10]string{}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			grid[i][j] = blocked
		}
	}
    // 4. Pick a random Cell, set it to state Passage and 
    // Compute its frontier cells. 
    // A frontier cell of a Cell is a cell with distance 2 in state Blocked and within the grid.
    var rando_coord = []int{rand.Intn(height), rand.Intn(width)}
	grid[rando_coord[0]][rando_coord[1]] = passage
	for i := 0; i < height; i++ {
		fmt.Println(grid[i])
	}
    // 3. While the list of frontier cells is not empty:
        //Pick a random frontier cell from the list of frontier cells.
        //Let neighbors(frontierCell) = All cells in distance 2 in state Passage. 
        //Pick a random neighbor and connect the frontier cell with the neighbor by setting the cell in-between to state Passage. 
        //Compute the frontier cells of the chosen frontier cell and add them to the frontier list. 
        //Remove the chosen frontier cell from the list of frontier cells.
}

// create new grid using Grid struct
func new_grid(w, h int) *Grid {
    grid := &Grid{
        width := w
        height := h
        // create new instances of cell structs and store them in cell array to create grid 
    }

}
