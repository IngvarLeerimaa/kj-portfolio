
// Login page logic
import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Profile");
    }

    async getHtml() {
        return `        <main>
            <h1> Kuidas ta siia sai? </h1>
     </main>
    `
    }

    async init() {
    }
}