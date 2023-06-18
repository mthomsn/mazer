export function findGrid() {
    const grid = document.querySelector('.grid');
    // console.log(grid);

    // iterateThroughGrid(grid);
    return grid;
}

export function startAndTarget(grid) {
    // console.log(grid.childNodes[0].children[0]);
    let num1 = randomNum(grid.childNodes.length);
    let num2 = randomNum(grid.childNodes[0].children.length);

    let start = grid.childNodes[num1].children[num2];
    let target = grid.childNodes[randomNum(grid.childNodes.length)].children[randomNum(grid.childNodes[0].children.length)];

    // console.log(grid.childNodes[num1].children[num2]);
    return [start, target];

}



export function randomNum(size) {
    return Math.floor(Math.random() * size);
}

// takes two 2d arrays and compares them
// export function compareSquares(array1, array2) {
//     // if array1 || array2 typeof != array
//     // return false
//     // if array1.length != array2.length
//     // return false
//     // iterate through array1 and compare array1[i][j] to array2[i][j]
//     // if array1[i][j] != array2[i][j]

// }