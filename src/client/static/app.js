// Generate maze event listener
const generateMaze = document.getElementById('generateMaze');

generateMaze.addEventListener('click', () => {
    // Maze inputs
    let cols = document.getElementById('mazeCol').valueAsNumber;
    let rows = document.getElementById('mazeRow').valueAsNumber;
    let scale = 10;

    if(cols > 0 && cols < 1001 && rows > 0 && rows < 1001) {
        getMaze(cols, rows, scale)
        .then(data => {
            console.log(data);
        })
    } else {
        console.log(cols, rows)
        console.log("Bad inputs");
    }
})

async function getMaze(cols, rows, scale) {
    const response = await fetch(`http://localhost:8000/maze?columns=${cols}&rows=${rows}&scale=${scale}`);
    const maze = await response.json();

    return maze
}

function drawMaze(cols, rows, cells) {
    let maze = document.getElementById("maze");
    let context = maze.getContext("2d")
    for (let y = 0; y < rows; y++) {
        for (let x = 0; x < cols; x++) {
            context.moveTo(
        }
    }
}
