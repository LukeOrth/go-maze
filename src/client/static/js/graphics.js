function setup() {
    // SCENE
    const scene = new THREE.Scene();
    
    // CAMERA
    const camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
    camera.position.set(0, 0, 100);
    camera.lookAt(0, 0, 0);

    // RENDER
    const renderer = new THREE.WebGLRenderer({
        canvas: document.getElementById("maze-graphic")
    });
    renderer.setPixelRatio(window.devicePixelRatio);
    renderer.setSize(window.innerWidth, window.innerHeight);

    return {scene, camera, renderer};
} 

const getBorders = {
    1: (s, x1, y1, x2, y2) => {s.add(borders([[x1, y1], [x1, y2]]))},
    2: (s, x1, y1, x2, y2) => {s.add(borders([[x1, y2], [x2, y2]]))},
    3: (s, x1, y1, x2, y2) => {s.add(borders([[x1, y1], [x1, y2], [x2, y2]]))},
    4: (s, x1, y1, x2, y2) => {s.add(borders([[x2, y1], [x2, y2]]))},
    5: (s, x1, y1, x2, y2) => {
        s.add(borders([[x1, y1], [x1, y2]]));
        s.add(borders([[x2, y1], [x2, y2]]));
    },
    6: (s, x1, y1, x2, y2) => {s.add(borders([[x1, y2], [x2, y2], [x2, y1]]))},
    7: (s, x1, y1, x2, y2) => {s.add(borders([[x1, y1], [x1, y2], [x2, y2], [x2, y1]]))},
    8: (s, x1, y1, x2, y2) => {s.add(borders([[x1, y1], [x2, y1]]))},
    9: (s, x1, y1, x2, y2) => {s.add(borders([[x1, y2], [x1, y1], [x2, y1]]))},
    10: (s, x1, y1, x2, y2) => {
        s.add(borders([[x1, y1], [x2, y1]]));
        s.add(borders([[x1, y2], [x2, y2]]));
    },
    11: (s, x1, y1, x2, y2) => {s.add(borders([[x2, y1], [x1, y1], [x1, y2], [x2, y2]]))},
    12: (s, x1, y1, x2, y2) => {s.add(borders([[x1, y1], [x2, y1], [x2, y2]]))},
    13: (s, x1, y1, x2, y2) => {s.add(borders([[x1, y2], [x1, y1], [x2, y1], [x2, y2]]))},
    14: (s, x1, y1, x2, y2) => {s.add(borders([[x1, y1], [x2, y1], [x2, y2], [x1, y2]]))},
    15: (s, x1, y1, x2, y2) => {s.add(borders([[x1, y1], [x2, y1], [x2, y2], [x1, y2], [x1, y1]]))},
};

function draw(maze) {
    let {scene, camera, renderer} = setup();
    
    // LOGIC
    const zeroX = -visibleWidthAtZDepth(0, camera) / 2;
    const zeroY = visibleHeightAtZDepth(0, camera) / 2;

    for (let i = 0; i < maze.cells.length; i++) {
        console.log(maze.cells[i]);
        drawCell = getBorders[maze.cells[i]];
        //console.log(drawCell(zeroX, zeroY, zeroX + maze.scale, zeroY + maze.scale));
    }

    //let coordinates = [[zeroX, zeroY],[zeroX + 10, zeroY],[zeroX + 10, zeroY - 10]];
    //scene.add(borders(coordinates));
    
    myFunc = getBorders[15];
    myFunc(scene, 0, 0, 10, 10);

    renderer.render(scene, camera);
}


function borders(coords) {
    let points = [];
    for (let i = 0; i < coords.length; i++) {
        let x = coords[i][0];
        let y = coords[i][1];
        points.push(new THREE.Vector2(x, y));
    }
    const color = new THREE.Color("lime");
    const material = new THREE.LineBasicMaterial({color: color});
    const geometry = new THREE.BufferGeometry().setFromPoints(points);
    const line = new THREE.Line(geometry, material);
    return line;
};

function visibleHeightAtZDepth(depth, camera ) {
    // compensate for cameras not positioned at z=0
    const cameraOffset = camera.position.z;
    if ( depth < cameraOffset ) depth -= cameraOffset;
    else depth += cameraOffset;

    // vertical fov in radians
    const vFOV = camera.fov * Math.PI / 180; 

    // Math.abs to ensure the result is always positive
    return 2 * Math.tan( vFOV / 2 ) * Math.abs( depth );
};


function visibleWidthAtZDepth(depth, camera) {
    const height = visibleHeightAtZDepth( depth, camera );
    return height * camera.aspect;
};
