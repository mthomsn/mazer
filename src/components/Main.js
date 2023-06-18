import React from "react";  
// import GenerateGrid from "./GenerateGrid";
import GenerateMaze from "./GenerateMaze";

export default function Main() {
    
    return (
        <div>
            <h1>Mazer</h1>
            <GenerateMaze rows={20} cols={20} />
        </div>
    );
}

