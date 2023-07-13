import { GameObject } from "./GameObject";
import { Cells } from "./Cells";
export class Snake extends GameObject {
    constructor(info, gamemap) {
        super();

        this.id = info.id;
        this.color = info.color;
        this.gamemap = gamemap;
        this.cells = [new Cells(info.r, info.c)];
        this.speed = 5; //每帧走5个格子
        this.direction = -1; //下一步指令 0 1 2 3表示上右下左
        this.status = 0; //0不动 1运动 -1 死了
        this.nextcell = null; //下一步的目的地
        this.dx = [-1, 0, 1, 0];
        this.dy = [0, 1, 0, -1];
        this.step = 0;//回合数
        this.eps = 1e-2;
    }
    start() {

    }
    set_direction(d) {
        this.direction = d;
    }
    check_tail_increasing() {
        if (this.step <= 10) return true;
        else if ((this.step - 10) % 3 === 0) return true;
        return false;
    }
    next_step() { //将蛇的状态变为走下一步
        const d = this.direction;
        this.nextcell = new Cells(this.cells[0].r + this.dx[d], this.cells[0].c + this.dy[d]);
        this.direction = -1;
        this.status = 1;
        this.step++;

        const k = this.cells.length;
        for (let i = k; i > 0; i--) {
            this.cells[i] = JSON.parse(JSON.stringify(this.cells[i - 1]));
        }
        if (!this.gamemap.check_valid(this.nextcell)) {
            this.status = -1;
        }
    }
    update_move() {
        const dx = this.nextcell.x - this.cells[0].x;
        const dy = this.nextcell.y - this.cells[0].y;
        const dist = Math.sqrt(dx * dx + dy * dy);
        if (dist < this.eps) {
            this.cells[0] = this.nextcell;
            this.nextcell = null;
            this.status = 0;
            if (!this.check_tail_increasing()) {
                this.cells.pop();
            }
        } else {
            const move_distance = this.speed * this.timedelta / 1000;
            this.cells[0].x += move_distance * dx / dist;
            this.cells[0].y += move_distance * dy / dist;

            if (!this.check_tail_increasing()) {
                const k = this.cells.length;
                const tail = this.cells[k - 1], target = this.cells[k - 2];
                const tail_dx = target.x - tail.x;
                const tail_dy = target.y - tail.y;
                tail.x += move_distance * tail_dx / dist;
                tail.y += move_distance * tail_dy / dist;
            }
        }

    }
    update() {

        if (this.status === 1) {
            this.update_move();
        }
        this.render();
    }

    render() {
        const L = this.gamemap.L;
        const ctx = this.gamemap.ctx;
        ctx.fillStyle = this.color;
        if (this.status === -1) {
            ctx.fillStyle = 'white';
        }
        for (const cell of this.cells) {
            ctx.beginPath();
            ctx.arc(cell.x * L, cell.y * L, L / 2 * 0.8, 0, Math.PI * 2);
            ctx.fill();
        }
        for (let i = 1; i < this.cells.length; i++) {
            const a = this.cells[i - 1], b = this.cells[i];
            if (Math.abs(a.x - b.x) < this.eps && Math.abs(a.y - b.y) < this.eps)
                continue;
            if (Math.abs(a.x - b.x) < this.eps) {
                ctx.fillRect((a.x - 0.4) * L, Math.min(a.y, b.y) * L, L * 0.8, Math.abs(a.y - b.y) * L);
            } else {
                ctx.fillRect(Math.min(a.x, b.x) * L, (a.y - 0.4) * L, Math.abs(a.x - b.x) * L, L * 0.8);
            }
        }


    }
}