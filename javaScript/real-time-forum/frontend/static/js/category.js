/* export const categoryTemplate = "asdasd"
 */

import { sendEvent } from "./index.js";
import AbstractView from "./AbstractView.js";
import { Router } from "./route.js";
export let currentThreadId
export let currentCategory;
export class GetCategoriesEvent {
    constructor(type, payload) {
        this.type = type;
        this.payload = payload;
    }
}

export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Category");
    }
    
    async getHtml() {
        return `
                <h1>Choose a category</h1>
                
            `;
    }

    async init(){   
        
        setTimeout(() => {
            sendEvent("get_categories");
            sendEvent("get_users");
        }, 100)

    }  
}


export function appendCategories(categoriesEvent) {
    currentCategory = new Map();

    // Ensure categoriesEvent is an object with a payload property
    if (typeof categoriesEvent !== 'object' || !categoriesEvent.payload) {
        console.error("Invalid categoriesEvent. It should be an object with a payload property.");
        return;
    }

    // Access the payload directly
    const categories = categoriesEvent.payload;

    // Check if the result is an array and it's not empty
    if (!Array.isArray(categories) || categories.length === 0) {
        console.error("Payload does not contain valid categories.");
        return;
    }

    //console.log("categories:", categories);

    // Create a new div for the categories container
    const categoriesDiv = document.querySelector('main');
    categoriesDiv.innerHTML = '';
    categoriesDiv.id = 'categories';
    // Iterate through the categories array and create an <a> element for each category
    categories.forEach(category => {
        categoriesDiv.appendChild(document.createElement('br'));    
        currentCategory.set(category.Id, category.Name);
        const test = document.createElement('div');
        const categoryLink = document.createElement('a');
        const categoryName = document.createElement('span');
        categoryName.textContent = category.Name;
        categoryName.classList.add('category-name');
        const categoryDescription = document.createElement('span');
        categoryDescription.textContent = category.Description;
        categoryDescription.classList.add('category-description');
        categoryLink.appendChild(categoryName);
        test.appendChild(categoryLink);
        test.appendChild(categoryDescription);
    
        categoryLink.href = category.Id;
        test.classList.add('category-row');
    
        categoriesDiv.appendChild(test);
    });
    
    
    
  /*   console.log("categoryEvent:", categoriesEvent);
    console.log("categories.payload:", categories);
    console.log("currentCategory:", currentCategory); */
   
    // Retrieve all the <a> elements that were created
    const categoryLinks = document.querySelectorAll('#categories a');

    // Attach a click event listener to each <a> element
    categoryLinks.forEach(link => {
        link.addEventListener('click', function(event) {
            event.preventDefault(); 
            currentThreadId = link.getAttribute("href");
           //  history.pushState(null, null, "/thread");
            Router("/thread" );
        });
    });
}

