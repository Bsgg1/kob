const Game_Object = [];

export class GameObject {
    constructor() {
        Game_Object.push(this);
        this.has_start = false;
        this.timedelta = 0;

    }
    start() {

    }
    update() {

    }
    on_destory() {

    }
    destory() {
        this.on_destory();
        for (let i in Game_Object) {
            const obj = Game_Object[i];
            if (obj === this) {
                Game_Object.splice(i);
                break;
            }
        }
    }
}
let last_timestamp;
const step = timestamp => {
    for (let obj of Game_Object) {
        if (!obj.has_start) {
            obj.has_start = true;
            obj.start();
        } else {
            obj.timedelta = timestamp - last_timestamp;
            obj.update();
        }
    }
    last_timestamp = timestamp;
    requestAnimationFrame(step)
}

requestAnimationFrame(step)