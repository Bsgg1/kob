import { GameObject } from "./GameObject";
import { Snake } from "./Snake";
import { Wall } from "./Wall";
export class GameMap extends GameObject {
    constructor(ctx, parent) {
        super();
        this.ctx = ctx;
        this.parent = parent;
        this.L = 0;
        this.rows = 13;
        this.cols = 13;
        this.walls = [];
        this.walls_count = 20;
        this.Snakes = [
            new Snake({ id: 0, color: "#4876ec", r: this.rows - 2, c: 1 }, this),
            new Snake({ id: 1, color: "#f94848", r: 1, c: this.cols - 2 }, this),
        ]
    }
    //检查两个玩家的连通性
    check(g, sx, sy, tx, ty) {
        if (sx == tx && sy == ty) return true;
        g[sx][sy] = true;
        let dx = [0, 0, 1, -1], dy = [1, -1, 0, 0];
        for (let i = 0; i < 4; i++) {
            let nx = sx + dx[i], ny = sy + dy[i];
            if (!g[nx][ny] && this.check(g, nx, ny, tx, ty)) return true;
        }
        return false;
    }
    create_walls() {
        const g = [];
        for (let i = 0; i < this.rows; i++) {
            g[i] = [];
            for (let j = 0; j < this.cols; j++) {
                g[i][j] = false;
            }
        }
        for (let i = 0; i < this.rows; i++) {
            g[i][0] = g[i][this.cols - 1] = true;
        }
        for (let i = 0; i < this.cols; i++) {
            g[0][i] = g[this.rows - 1][i] = true;
        }
        for (let i = 0; i < 20; i += 2) {
            for (let j = 0; j <= 1000; j++) {
                let r = parseInt(Math.random() * this.rows);
                let c = parseInt(Math.random() * this.cols);
                if (g[r][c] || g[c][r]) continue;
                if (r == this.rows - 2 && c == 1 || r == 1 && c == this.cols - 1) continue;
                g[r][c] = g[c][r] = true;
                break;
            }
        }
        const copy_g = JSON.parse(JSON.stringify(g));
        if (!this.check(copy_g, this.rows - 2, 1, 1, this.cols - 2)) return false;
        for (let i = 0; i < this.rows; i++) {
            for (let j = 0; j < this.cols; j++) {
                if (g[i][j])
                    this.walls.push(new Wall(i, j, this));
            }
        }
        return true;
    }
    check_valid(cell) {
        for (const wall of this.walls) {
            if (wall.r === cell.r && wall.c === cell.c) return false;
        }

        for (const snake of this.Snakes) {
            let k = snake.cells.length;
            if (!snake.check_tail_increasing()) {
                k--;
            }
            for (let i = 0; i < k; i++) {
                if (cell.r === snake.cells[i].r && cell.c === snake.cells[i].c) return false;
            }
        }
        return true;
    }
    add_listening_events() {
        this.ctx.canvas.focus();
        const [snake0, snake1] = this.Snakes;
        this.ctx.canvas.addEventListener("keydown", e => {
            console.log(e.key);
            if (e.key === 'w') snake0.set_direction(0);
            else if (e.key === 'd') snake0.set_direction(1);
            else if (e.key === 's') snake0.set_direction(2);
            else if (e.key === 'a') snake0.set_direction(3);
            else if (e.key === 'ArrowUp') snake1.set_direction(0);
            else if (e.key === 'ArrowRight') snake1.set_direction(1);
            else if (e.key === 'ArrowDown') snake1.set_direction(2);
            else if (e.key === 'ArrowLeft') snake1.set_direction(3);
        });
    }
    start() {
        for (let i = 0; i <= 1000; i++) {
            if (this.create_walls()) break;
        }
        this.add_listening_events();
    }
    check_ready() { //判断两条蛇是否准备好进入下一回合
        for (const snake of this.Snakes) {
            if (snake.status !== 0) return false;
            if (snake.direction === -1) return false;
        }
        return true;
    }
    update_size() {
        this.L = parseInt(Math.min(this.parent.clientWidth / this.cols, this.parent.clientHeight / this.rows));
        this.ctx.canvas.width = this.L * this.cols;
        this.ctx.canvas.height = this.L * this.rows;
    }
    next_step() {
        for (const snake of this.Snakes) {
            snake.next_step();
        }
    }
    update() {
        this.update_size();
        if (this.check_ready()) {
            this.next_step();
        }
        this.render();
    }

    render() {
        // 取颜色
        const color_eve = "#AAD751", color_odd = "#A2D149";
        // 染色
        for (let r = 0; r < this.rows; r++)
            for (let c = 0; c < this.cols; c++) {
                if ((r + c) % 2 == 0) {
                    this.ctx.fillStyle = color_eve;
                } else {
                    this.ctx.fillStyle = color_odd;
                }
                //左上角左边，明确canvas坐标系
                this.ctx.fillRect(c * this.L, r * this.L, this.L, this.L);
            }
    }


}