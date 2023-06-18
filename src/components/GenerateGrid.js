import React, { useEffect, useRef } from 'react';

function GenerateMaze({ rows, cols }) {
  const gridContainerRef = useRef(null);

  useEffect(() => {
    const gridContainer = gridContainerRef.current;
    
    // if (gridContainer.childNodes.length === 0) {
    //   const gridElement = createGridElement(rows, cols);
    //   gridContainer.appendChild(gridElement);
    // }

    const gridElement = createGridElement(rows, cols);
    gridContainer.appendChild(gridElement);

  }, [rows, cols]);

  const createGridElement = (rows, cols) => {
    const gridElement = document.createElement('div');
    gridElement.setAttribute('class', 'empty-grid');

    for (let row = 0; row < rows; row++) {
      const rowElement = document.createElement('div');
      rowElement.setAttribute('class', 'row');

      for (let col = 0; col < cols; col++) {
        const squareElement = document.createElement('div');
        squareElement.setAttribute('class', 'empty-grid-square');

        rowElement.appendChild(squareElement);
      }

      gridElement.appendChild(rowElement);
    }

    console.log(gridElement);

    return gridElement;
  };

  return <div ref={gridContainerRef} className='container'></div>;
};

export default GenerateMaze;
