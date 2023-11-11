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
var neighbors = [][]int{ {-1, 0}, {1, 0}, {0, -1}, {0, 1} }
var blocked string = "#"
var passage string = " "

func main() {
    // 1. A Grid consists of a 2 dimensional array of cells.
    var height = 5 
    var width = 5
    var grid Grid
    var cells = make([][]Cell, width+1) 

    grid.height = height
    grid.width = width
    grid.cells = cells

    // 2. A Cell has 2 states: Blocked or Passage.
    //      - made blocked/passage global variables so they can be accessed by helper functions
    // 3. Start with a Grid full of Cells in state Blocked.
    for i := 0; i <= grid.height; i++ {
        for j := 0; j <= grid.width; j++ {
            var new_cell = Cell{ coord: []int{i, j} , state: blocked }
            cells[i] = append(cells[i], new_cell)
        }
    }

    // 4. Pick a random Cell, set it to state Passage
    var rando_coord = []int{rand.Intn(height), rand.Intn(width)}
    var random_cell = grid.cells[rando_coord[0]][rando_coord[1]]
    grid.cells[rando_coord[0]][rando_coord[1]].state = passage
   
    // 5. Compute its frontier cells. 
    // A frontier cell of a Cell is a cell with distance 2 in state Blocked and within the grid.
    frontier := [][]int{}
    init_frontier := find_frontier(random_cell.coord, grid)
    frontier = append_frontier(frontier, init_frontier) 
    display_grid(grid)
    fmt.Println(len(frontier))

    // 6. While the list of frontier cells is not empty:
    for {
        if len(frontier) < 1 {
            break
        }
        rand_index := rand.Intn(len(frontier))
        //Pick a random frontier cell from the list of frontier cells.
        //Let neighbors(frontierCell) = All cells in distance 2 in state Passage. 
        neighbors := find_neighbors(frontier[rand_index], grid)
        //Pick a random neighbor and connect the frontier cell with the neighbor by setting the cell in-between to state Passage. 
        rand_neighbor := rand.Intn(len(neighbors))
        fmt.Println(rand_neighbor)
        //Compute the frontier cells of the chosen frontier cell and add them to the frontier list. 
        //Remove the chosen frontier cell from the list of frontier cells.
        frontier = rm_itm(frontier, rand_index)
    }
    fmt.Println("done")
}

func build_maze() {

}

func find_frontier(current []int, grid Grid) [][]int {
    var fcells = [][]int{}
    for _, qty := range directions {
        var m = translate_slice(current, qty)
        // frontier cell is within grid
        if m[0] <= 0 || m[1] >= grid.width || m[0] >= grid.height || m[1] <= 0 {
            continue
        } 
        // frontier cell state = blocked
        if grid.cells[m[0]][m[1]].state == passage {
            continue
        } 
        fcells = append(fcells, m)
    }
    return fcells
}

func find_neighbors(current []int, grid Grid) [][]int{
    // all cells that are distance 2 away from current cell with state of passage
    neighbors := [][]int{}
    for _, qty := range directions {
        var m = translate_slice(current, qty)
        if m[0] <= 0 || m[1] >= grid.width || m[0] >= grid.height || m[1] <= 0 {
            continue
        }
        if grid.cells[m[0]][m[1]].state == blocked {
            continue
        } 
        neighbors = append(neighbors, m)
    }
    return neighbors
}

func translate_slice(input, qty []int) []int {
    var new_slice = []int{}
    for i, _ := range qty  {
        new_slice = append(new_slice, input[i]+qty[i])
    }
    return new_slice
}

func display_grid(grid Grid) {
    for i:=0; i<grid.height; i++ {
        fmt.Println(grid.cells[i])
    }
}

func append_frontier(current, new [][]int) [][]int {
    for _, value := range new {
        current = append(current, value)
    }
    return current
}

func rm_itm(s [][]int, index int) [][]int {
    s = append(s[:index], s[index+1:]...)
    return s
}

func find_inbetween(current, neighbor [][]int) [][]int {
    // subtract frontier from neighbor => difference 
    // if current[0] % neighbor[0] == 0 {
        // if current[0] - neighbor[0] > 0 {
            // translate current [1, 0] & set cell state = passage 
        //} else {
            // translate current [-1, 0] & set cell state = passage
        //}
    //} else if current[1] % neighbor[1] == 0 {
        // if current[1] - neighbor[1] > 0 {
            // translate current [0, 1] & set cell state = passage 
        //} else {
            // translate current [0, -1] & set cell state = passage
        //}
    //}
    if current[0] % neighbor[0] == 0 {
        if current[0] - neighbor[0] > 0 {
            // translate current [1, 0] & set cell state = passage 
        } else {
            // translate current [-1, 0] & set cell state = passage
        }
    } else if current[1] % neighbor[1] == 0 {
        if current[1] - neighbor[1] > 0 {
            // translate current [0, 1] & set cell state = passage 
        } else {
            // translate current [0, -1] & set cell state = passage
        }
    }
    } 
}














