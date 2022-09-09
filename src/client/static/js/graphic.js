class Graphic {
    constructor(element) {
        this.scene = new THREE.Scene();
        this.camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 1, 1000);
        this.renderer = new THREE.WebGLRenderer({ canvas : element });
        this.color = new THREE.Color("lime");
    }

    setup() {
        this.camera.position.set(0, 0, 100);
        this.camera.lookAt(0, 0, 0);
        this.renderer.setPixelRatio(window.devicePixelRatio);
        this.renderer.setSize(window.innerWidth, window.innerHeight);
        this.width = -(this.visibleWidth() / 2);
        this.height = this.visibleHeight() / 2;
    }

    drawGrid(cols, rows) {
        let scale = Math.min(Math.abs(this.width / cols), this.height / rows) * 2;
        for (let i = 0; i < cols; i++) {
            let x1 = this.width + i * scale;
            let x2 = x1 + scale;
            for (let j = 0; j < rows; j++) {
                let y1 = this.height - j * scale;
                let y2 = y1 - scale;

                let cell = this.cellType[15];
                this.scene.add(cell(x1, y1, x2, y2)[0]);
            }
        }
    }

    genMaze(maze) {
        let scale = Math.min(Math.abs(this.width / maze.columns), this.height / maze.rows) * 2;
        let x1 = this.width;
        let y1 = this.height;
        let x2 = x1 + scale;
        let y2 = y1 - scale;

        const timer = ms => new Promise(res => setTimeout(res, ms))
        
        const load = async() => {
            for (let i = 0; i < maze.moves.length; i++) {
                const move = maze.moves[i];
                let cell = this.cellType[move];
                ({x1, y1, x2, y2} = this.updateXY[move](x1, y1, x2, y2, scale));
                this.scene.add(cell(x1, y1, x2, y2)[0]);
                this.renderer.render(this.scene, this.camera);
                await timer(1000);
            }
        }
        load();
    }

    updateXY = {
        8: (x1, y1, x2, y2, scale) => {return {x1: x1, y1: y1 + scale, x2: x2, y2: y2 + scale}},
        4: (x1, y1, x2, y2, scale) => {return {x1: x1 + scale, y1: y1, x2: x2 + scale, y2: y2}},
        2: (x1, y1, x2, y2, scale) => {return {x1: x1, y1: y1 - scale, x2: x2, y2: y2 - scale}},
        1: (x1, y1, x2, y2, scale) => {return {x1: x1 - scale, y1: y1, x2: x2 - scale, y2: y2}},
    }
    
    drawMaze(maze) {
        let scale = Math.min(Math.abs(this.width / maze.columns), this.height / maze.rows) * 2;

        for (let i = 0; i < maze.cells.length; i++) {
            if (maze.cells[i] == 0) {
                continue;
            }
            let x1 = this.width + (i % maze.columns) * scale
            let y1 = this.height - Math.floor((i / maze.columns % maze.rows)) * scale;
            let x2 = x1 + scale;
            let y2 = y1 - scale;

            let lines = this.cellType[maze.cells[i]];
            let results = lines(x1, y1, x2, y2);
            for (let j = 0; j < results.length; j++) {
                this.scene.add(results[j]);
            }
        }
        this.renderer.render(this.scene, this.camera);
    }

    cellType = {
        1: (x1, y1, x2, y2) => {return [this.borders([[x1, y1], [x1, y2]])]},
        2: (x1, y1, x2, y2) => {return [this.borders([[x1, y2], [x2, y2]])]},
        3: (x1, y1, x2, y2) => {return [this.borders([[x1, y1], [x1, y2], [x2, y2]])]},
        4: (x1, y1, x2, y2) => {return [this.borders([[x2, y1], [x2, y2]])]},
        5: (x1, y1, x2, y2) => {return [
            this.borders([[x1, y1], [x1, y2]]),
            this.borders([[x2, y1], [x2, y2]])
        ]},
        6: (x1, y1, x2, y2) => {return [this.borders([[x1, y2], [x2, y2], [x2, y1]])]},
        7: (x1, y1, x2, y2) => {return [this.borders([[x1, y1], [x1, y2], [x2, y2], [x2, y1]])]},
        8: (x1, y1, x2, y2) => {return [this.borders([[x1, y1], [x2, y1]])]},
        9: (x1, y1, x2, y2) => {return [this.borders([[x1, y2], [x1, y1], [x2, y1]])]},
        10: (x1, y1, x2, y2) => {return [
            this.borders([[x1, y1], [x2, y1]]),
            this.borders([[x1, y2], [x2, y2]])
        ]},
        11: (x1, y1, x2, y2) => {return [this.borders([[x2, y1], [x1, y1], [x1, y2], [x2, y2]])]},
        12: (x1, y1, x2, y2) => {return [this.borders([[x1, y1], [x2, y1], [x2, y2]])]},
        13: (x1, y1, x2, y2) => {return [this.borders([[x1, y2], [x1, y1], [x2, y1], [x2, y2]])]},
        14: (x1, y1, x2, y2) => {return [this.borders([[x1, y1], [x2, y1], [x2, y2], [x1, y2]])]},
        15: (x1, y1, x2, y2) => {return [this.borders([[x1, y1], [x2, y1], [x2, y2], [x1, y2], [x1, y1]])]},
    }

    borders(coords) {
        let points = [];
        for (let i = 0; i < coords.length; i++) {
            let x = coords[i][0];
            let y = coords[i][1];
            points.push(new THREE.Vector2(x, y));
        }
        const material = new THREE.LineBasicMaterial({color: this.color});
        const geometry = new THREE.BufferGeometry().setFromPoints(points);
        const line = new THREE.Line(geometry, material);
        return line;
    }

    visibleHeight() {
        // compensate for cameras not positioned at z=0
        const depth = this.camera.position.z;

        // vertical fov in radians
        const vFOV = this.camera.fov * Math.PI / 180; 

        // Math.abs to ensure the result is always positive
        return 2 * Math.tan( vFOV / 2 ) * Math.abs( depth );
    }

    visibleWidth() {
        return this.visibleHeight() * this.camera.aspect;
    }
}
