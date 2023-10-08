import { PriorityQueue } from 'js-sdsl';
import generate from 'generate-maze';

// data to play with
const rows = 10;
const cols = 10;
const grid = generate(rows, cols);
const start = { x: 1, y: 0 };
const target = { x: 5, y: 5 };
grid[start.y][start.x].start = true;
grid[target.y][target.x].target = true;    
// console.log(grid);

const options = ["top", "right", "bottom", "left"];

function dijkstra(grid, start, target) {
    // create priority queue
    let queue = new PriorityQueue();

    // set all nodes distance to infinity
    for (let i = 0; i < grid.length; i++){
        for (let j = 0; j < grid.length; j++){
            grid[i][j].distance = 9999;
        }
    }

    // set start node distance to 0
    grid[start.y][start.x].distance = 0;
    // console.log(grid[start.y][start.x].distance);


    // add start to queue with priority 0
    // queue.push({y: [start.y],x: [start.x]}, 0);
    queue.push(grid[start.y][start.x], 0);
    // console.log(queue);

    // while queue is not empty
    while (queue.length > 0){
        // set current node to node with lowest priority
        // currently pops the last item in the queue - or the highest priority - need to figure out this mechanic
        let current = queue.top();

        // if current node is target, return path
        if (current.target){
            // console.log(current.start);
            return current;
        }

        // for each neighbor of current node
        for (let neighbor in options){
            console.log(current[options[neighbor]]);
            // if neighbor is not visited
            if (current[options[neighbor]] == false){
                // calculate new distance
                console.log("boop");
                let newDistance = current.distance + 1;
                

                // if (newDistance < neighbor.distance){
                //     console.log(`new distance: ${newDistance}`);
                //     console.log(`neighbor distance: ${neighbor.distance}`)
                //     neighbor.distance = newDistance;
                // }
                // add neighbor to queue with priority new distance
                //    queue.push(grid[current.y], newDistance);
                // set neighbor's previous to current node
                // console.log(`neighbor: ${neighbor}`);
            }
        }
        queue.pop();
    }
    // console.log(grid);

    // return distance
    return 0;
}

function checkPaths(cell){
    let paths = [];
    // const options = ["top", "right", "bottom", "left"];

    for (let i = 0; i < options.length; i++){
        if (cell[options[i]] == false){
            paths.push(options[i]);
        }
    }

    return paths;
    // will return an array of available paths
}

function randomNum(size){
    return Math.floor(Math.random() * size);
}

function getNeighborCoord(direction){
    switch (direction) {
        case "top":

    }
}

// console.log(checkPaths({top: true, right: false, bottom: true, left: false}));

dijkstra(grid, start, target)