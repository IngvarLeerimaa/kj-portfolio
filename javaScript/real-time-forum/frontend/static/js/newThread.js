//export const newThreadTemplate = "We are in New Thread"

   /*  const subBtn = document.querySelector("input[type=submit]");
    
    subBtn.addEventListener ("click", () => {
        let checked = document.querySelectorAll("input[type=checkbox]:checked");
        if (checked.length === 0) {
        alert("Please select a category");
        event.preventDefault();
        } else {
        return true;
        }
    });
     */

import AbstractView from "./AbstractView.js";

export default class extends AbstractView{
    constructor(params) {
        super(params);
        this.setTitle("hiasdm");
    }

    async getHtml() {
        return "asduiqrfs"
    }

    async init() {
}
}


