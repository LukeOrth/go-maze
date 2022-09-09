// Init Graphic
const graphic = new Graphic(document.getElementById("graphic"));
graphic.setup();

// Generate maze event listener
const generateMaze = document.getElementById('generateMaze');

generateMaze.addEventListener('click', () => {
    // Maze inputs
    let cols = document.getElementById('mazeCol').valueAsNumber;
    let rows = document.getElementById('mazeRow').valueAsNumber;
    let scale = 10;

    if(cols > 0 && cols < 1001 && rows > 0 && rows < 1001) {
        //graphic.drawGrid(cols, rows);
        getMaze(cols, rows, scale)
        .then(data => {
            graphic.drawMaze(data);
        });
    } else {
        console.log("Bad inputs");
    }
})

async function getMaze(cols, rows, scale) {
    const response = await fetch(`http://localhost:8000/maze?columns=${cols}&rows=${rows}&scale=${scale}`);
    const maze = await response.json();

    return maze
}
