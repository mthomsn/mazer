package main 
import (
    "fmt"
    "math/rand"
)

type Cell struct {
    coord []int
    state string
    frontier []int
}

type Grid struct {
    cells [][]Cell
    height int
    width int
}

var directions = [][]int{ {-2, 0}, {2, 0}, {0, -2}, {0, 2} }

func main() {
    var height = 5 
    var width = 5
    var blocked string = "#"
    var passage string = " "
    var grid Grid
    var cells = make([][]Cell, width+1) 

    grid.height = height
    grid.width = width
    grid.cells = cells

    // 1. A Grid consists of a 2 dimensional array of cells.
    // 2. A Cell has 2 states: Blocked or Passage.
    // 3. Start with a Grid full of Cells in state Blocked.
    for i := 0; i <= grid.height; i++ {
        for j := 0; j <= grid.width; j++ {
            var new_cell = Cell{ coord: []int{i, j} , state: blocked }
            //fmt.Println(new_cell)
            cells[i] = append(cells[i], new_cell)
        }
        //fmt.Println(cells[i])
    }

    // 4. Pick a random Cell, set it to state Passage and 
    var rando_coord = []int{rand.Intn(height), rand.Intn(width)}
    var random_cell = grid.cells[rando_coord[0]][rando_coord[1]]
    random_cell.state = passage
   
    // Compute its frontier cells. 
    // A frontier cell of a Cell is a cell with distance 2 in state Blocked and within the grid.\
    test_cell := grid.cells[3][3].coord
    var move = []int{-2, 0}
    var test_trans = translate_slice(test_cell, move)
    fmt.Println(test_trans)

    //fmt.Println(find_frontier(test_cell))

    // 3. While the list of frontier cells is not empty:
        //Pick a random frontier cell from the list of frontier cells.
        //Let neighbors(frontierCell) = All cells in distance 2 in state Passage. 
        //Pick a random neighbor and connect the frontier cell with the neighbor by setting the cell in-between to state Passage. 
        //Compute the frontier cells of the chosen frontier cell and add them to the frontier list. 
        //Remove the chosen frontier cell from the list of frontier cells.
}

func find_frontier(current []int) int {
    var fcells = []int{}
    
    for _, value := range directions {
        fmt.Println(value)
    }

    fmt.Println(fcells)
    return 0
    // initial checks
    // frontier cell is within grid
   
    // frontier cell state = blocked
}

func translate_slice(input, qty []int) []int {
    var new_slice = []int{}
    for i, _ := range qty  {
        new_slice = append(new_slice, input[i]+qty[i])
    }
    return new_slice
}
















