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
var frontier string = "o"

func main() {
    // 1. A Grid consists of a 2 dimensional array of cells.
    var height = 10 
    var width = 10 
    var grid Grid
    var cells = make([][]Cell, width+1, height+1) 

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
    fmt.Println(grid.cells[rando_coord[0]][rando_coord[1]])
   
    // 5. Compute its frontier cells. 
    // A frontier cell of a Cell is a cell with distance 2 in state Blocked and within the grid.
    frontier := [][]int{}
    init_frontier := find_frontier(random_cell.coord, grid)
    frontier = append_frontier(frontier, init_frontier) 

    // 6. While the list of frontier cells is not empty:
    counter := 0
    for {
        if len(frontier) < 1 {
            fmt.Printf("break from frontier%v \n", frontier)
            break
        }
        fmt.Printf("counter: %v \n", counter)
        rand_index := rand.Intn(len(frontier))
        random_frontier_cell := frontier[rand_index]

        //Pick a random frontier cell from the list of frontier cells.
        //Let neighbors(frontierCell) = All cells in distance 2 in state Passage. 
        neighbors := find_neighbors(random_frontier_cell, grid)
        //Pick a random neighbor and connect the frontier cell with the neighbor by setting the cell in-between to state Passage. 
        if len(neighbors) == 0 {
            fmt.Printf("break from neighbors %v \n", neighbors)
            break 
        } 
        fmt.Printf("random frontier cell: %v \n", frontier[rand_index])
        fmt.Printf("neighbors of rand frontier: %v \n", neighbors)
        rand_neighbor := rand.Intn(len(neighbors))
        inbetween := find_inbetween(frontier[rand_index], neighbors[rand_neighbor])
        grid.cells[inbetween[0]][inbetween[1]].state = passage
        grid.cells[random_frontier_cell[0]][random_frontier_cell[1]].state = passage
        neighbors = nil
        //Compute the frontier cells of the chosen frontier cell and add them to the frontier list.
        add_cells := find_frontier(frontier[rand_index], grid)
        //fmt.Printf("cells to add: %v \n", add_cells)
        frontier = append_frontier(frontier, add_cells)
        //Remove the chosen frontier cell from the list of frontier cells.
        frontier = rm_itm(frontier, rand_index)
        //fmt.Printf("frontier @ end of loop %v \n", frontier)
        fmt.Println(frontier)
        display_grid(grid)
        counter++
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
        if grid.cells[m[0]][m[1]].state == frontier {
            continue
        }
        grid.cells[m[0]][m[1]].state = frontier 
        fcells = append(fcells, m)
    }
    fmt.Printf("cells to add to frontier: %v \n", fcells)
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
        if grid.cells[m[0]][m[1]].state == frontier {
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
    var display = [][]string{}
    var row = []string{}
    for i:=0; i<=grid.height; i++ {
        for j:=0; j<=grid.width; j++ {
            row = append(row, grid.cells[i][j].state)
        }
        display = append(display, row)
        row = nil
    }
    for i:=0; i<=grid.height; i++ {
        fmt.Println(display[i])
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

func find_inbetween(current, neighbor []int) []int {
    //fmt.Printf("current: %v \n", current)
    //fmt.Printf("neighbor: %v \n", neighbor)

    type Directions struct {
        right []int
        left []int
        up []int
        down []int
    }
    dir := Directions{
        right: []int{0, 1},
        left: []int{0, -1},
        up: []int{-1, 0},
        down: []int{1, 0},
    }
    var target = []int{}

    if current[0] - neighbor[0] == 0 {
        if current[1] - neighbor[1] > 0 {
            target = translate_slice(current, dir.left)
        } else {
            target = translate_slice(current, dir.right)
        }
    } else if current[1] - neighbor[1] == 0 {
        if current[0] - neighbor[0] > 0 {
            target = translate_slice(current, dir.up)
        } else {
            target = translate_slice(current, dir.down)
        }
    }
    return target     
}














