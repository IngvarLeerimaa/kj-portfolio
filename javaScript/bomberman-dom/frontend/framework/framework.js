class NewElement {
    constructor(tag, {attrs = {}, children = []} = {}) {
        this.tag = tag;
        this.attrs = attrs || {};
        this.children = children || [];
    }
    create() {
        const element = document.createElement(this.tag);
        
        for (const [k, v] of Object.entries(this.attrs)) {
            element.setAttribute(k, v);
        }

        this.children.forEach(child => {
            if (typeof child === 'string') {
                element.appendChild(document.createTextNode(child));
            } else {
                element.appendChild(new NewElement(...child).create());
            }
 
        })

        return element;
    }
}

class Listen {
    constructor(element, event, callback) {
        this.element = element;
        this.event = event;
        this.callback = callback;
    }
    listen() {
        this.element.addEventListener(this.event, this.callback);
    }
    remove() {
        this.element.removeEventListener(this.event, this.callback);
    }
}

class NewLink {
    constructor(path, text) {
        this.path = path;
        this.text = text;
    }

    newLink() {
        const link = document.createElement('a');
        link.classList.add('link');
        link.href = this.path;
        link.textContent = this.text;
        return link;
    }
}

class Route {
    constructor(path, callback) {
        this.path = path;
        this.callback = callback;
    }
}

class Router {
    constructor() {
        this.routes = [];
        window.addEventListener('popstate', this.resolve.bind(this));
    }

    addRoute(path, callback) {
        this.routes.push(new Route(path, callback));
    }

    resolve() {
        const path = window.location.pathname;
        const route = this.routes.find(route => route.path === path);

        if (route) {
            route.callback();
        } else {
            alert(`What are you trying to do? Path ${path} not found`);
            console.log(`Path ${path} not found`);
            // TODO: Handle 404 or redirect to a default route
        }
    }

    navigate(path) {
        window.history.pushState({}, '', path);
        this.resolve();
    }
}


const simple = {
    NewElement,
    Listen,
    NewLink,
    Route,
    Router,
};

export default simple;