import React, { useEffect, useRef } from 'react';
import generate from 'generate-maze';
import { randomNum } from '../helpers/startAndTarget';

function GenerateMaze({ rows, cols }) {
    const gridContainerRef = useRef(null);
    // let grid = generate(rows, cols);

    useEffect(() => {
        const gridContainer = gridContainerRef.current;

        if (gridContainer.childNodes.length === 0) {
            const grid = generate(rows, cols);
            console.log(grid);
            const start = { x: randomNum(cols), y: randomNum(rows) };
            const target = { x: randomNum(cols), y: randomNum(rows) };
            grid[start.y][start.x].start = true;
            grid[target.y][target.x].target = true;    
            // console.log(grid[target.y][target.x]);
            // console.log("hit");

            const gridElement = createGridElement(grid);
            gridContainer.appendChild(gridElement);
        }
    }, [rows, cols]);

    const createGridElement = (grid) => {
        const gridElement = document.createElement('div');
        gridElement.setAttribute('class', 'grid');

        for (let row = 0; row < grid.length; row++) {
            const rowElement = document.createElement('div');
            rowElement.setAttribute('class', 'row');

            for (let col = 0; col < grid[row].length; col++) {
                const squareElement = document.createElement('div');
                squareElement.setAttribute('class', 'square');

                if (grid[row][col].top) {
                squareElement.classList.add('top');
                }
                if (grid[row][col].bottom) {
                squareElement.classList.add('bottom');
                }
                if (grid[row][col].left) {
                squareElement.classList.add('left');
                }
                if (grid[row][col].right) {
                squareElement.classList.add('right');
                }
                if (grid[row][col].start) {
                    squareElement.classList.add('start');
                    squareElement.appendChild(document.createElement('div')).setAttribute('class', 'start-icon');
                }
                if (grid[row][col].target) {
                    squareElement.classList.add('target');
                    squareElement.appendChild(document.createElement('div')).setAttribute('class', 'target-icon');
                }
                // rowElement.appendChild(squareElement);
                rowElement.append(squareElement);
            }
            // gridElement.appendChild(rowElement);
            gridElement.append(rowElement);
        }
        return gridElement;
    };

    return <div ref={gridContainerRef} className='container'></div>;
};

export default GenerateMaze;