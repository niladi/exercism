class QueenAttack {
    white: [number, number];
    black: [number, number];

    constructor(positioning: {white: [number, number], black: [number, number]}) {
        this.white = positioning.white;
        this.black = positioning.black;
        if (!(this.isInBounce(this.white) && this.isInBounce(this.black))) {
            throw "Out of Bounce";
        }
        if (this.pointsAreEqual(this.black, this.white)) {
            throw "Queens cannot share the same space";
        }
    }

    /**
     * canAttak
     */
    public canAttack() : boolean {
        return this.isHorizontal() || this.isVertical() || this.isDiagonal() 
    }

    private isHorizontal() : boolean {
        return this.white[0] === this.black[0] 
    }

    private isVertical(): boolean {
        return this.white[1] === this.black[1]
    }

    private isDiagonal(): boolean {
        return Math.abs(this.white[1] - this.black[1]) === Math.abs(this.white[0] - this.black[0])
    }

    private isInBounce(value: [number,number]): boolean {
        return Math.abs(value[0]) < 8 && Math.abs(value[1]) < 8
    }


    /**
     * toString
     */
    public toString() : string {
        let s: string = ""
        for (let x = 0; x < 8; x++) {
            for (let y = 0; y < 8; y++) {
                s += this.pointsAreEqual(this.white, [x, y]) 
                    ? "W" 
                    : (this.pointsAreEqual(this.black, [x, y])  ? "B" : "_")
                s += " "
            }
            s = s.trim()
            s += "\n"
        }
        return s
    }

    private pointsAreEqual(num1: [number, number], num2: [number, number]) : boolean {
        return num1[0] === num2[0] && num1[1] === num2[1]
    }



}

export default QueenAttack