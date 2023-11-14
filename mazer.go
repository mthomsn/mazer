package main 
import (
    "fmt"
    "math/rand"
    "time"
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
var maze_start string = "s"
var maze_finish string = "f"

func main() {
    // 1. A Grid consists of a 2 dimensional array of cells.
    var height = 20 
    var width = 20 
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
    // need to add verification that random cell is within grid
    var rando_coord = rand_coord(grid)
    var random_cell = grid.cells[rando_coord[0]][rando_coord[1]]
    grid.cells[rando_coord[0]][rando_coord[1]].state = passage
   
    // 5. Compute its frontier cells. 
    // A frontier cell of a Cell is a cell with distance 2 in state Blocked and within the grid.
    frontier := [][]int{}
    init_frontier := find_frontier(random_cell.coord, grid)
    frontier = append_frontier(frontier, init_frontier) 

    // 6. While the list of frontier cells is not empty:
    counter := 0
    for {
        if len(frontier) < 1 {
            break
        }
        rand_index := rand.Intn(len(frontier))
        random_frontier_cell := frontier[rand_index]

        //Pick a random frontier cell from the list of frontier cells.
        //Let neighbors(frontierCell) = All cells in distance 2 in state Passage. 
        neighbors := find_neighbors(random_frontier_cell, grid)
        //Pick a random neighbor and connect the frontier cell with the neighbor by setting the cell in-between to state Passage. 
        if len(neighbors) == 0 {
            break 
        } 
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
        counter++
    }
    //fmt.Printf("counter: %v \n", counter)
    start := generate_start(grid)
    rand.Seed(time.Now().UnixNano())
    finish := generate_finish(grid, start)
    grid.cells[start[0]][start[1]].state = maze_start
    grid.cells[finish[0]][finish[1]].state = maze_finish 
    display_grid(grid)
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
    for i := range qty  {
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

func generate_start(grid Grid) []int{
    var rand_row int
    rand_row = rand.Intn(grid.height)
    var rand_col int
    if rand_row == 0 || rand_row == grid.height {
        rand_col = rand.Intn(grid.width)
        if rand_col == 0 {
            rand_col++
        } else if rand_col == grid.width {
            rand_col--
        }
    } else {
        options := []int{0, grid.width}
        rand_col = options[rand.Intn(len(options))]
    }
    start := []int{rand_row, rand_col}
    return start
}

func generate_finish(grid Grid, start []int) []int{
    var x, y int
    var finish = []int{}

    if start[0] <= grid.height / 2 {
        x = randInt((grid.height - (grid.height / 2)), grid.height)
    } else {
        x = randInt(0, (grid.height / 2))
    }
    finish = append(finish, x)

    if x == 0 || x == grid.height {
        y = randInt(0, grid.width)
    } else if start[1] < grid.width {
        y = grid.width 
    } else {
        y = 0 
    }
    finish = append(finish, y)
    return finish
}

func randInt(min, max int) int {
    return min + rand.Intn(max - min)
}

func rand_coord(grid Grid) []int {
    var random_coord = []int{}
    var row, col int

    for {
        random_coord = nil
        row = rand.Intn(grid.height)
        if row == 0 || row % 2 == 0 || row == grid.height {
            continue
        }
        random_coord = append(random_coord, row)
        col = rand.Intn(grid.width)
        if col == 0 || col % 2 == 0 || col == grid.width {
            continue
        }
        random_coord = append(random_coord, col)
        break
    } 
    return random_coord
}

func init() {
    rand.Seed(time.Now().UnixNano())
}







