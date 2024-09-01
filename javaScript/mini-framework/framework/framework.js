class NewElement {
    constructor(tag, id, className, innerHTML, placeholder, type) {
        this.tag = tag;
        this.id = id;
        this.className = className;
        this.innerHTML = innerHTML;
        this.placeholder = placeholder;
        this.type = type;
    }
    create() {
        const element = document.createElement(this.tag);
        if (this.id) {
            element.id = this.id;
        }
        if (this.className) {
            element.className = this.className;
        }
        if (this.innerHTML) {
            element.innerHTML = this.innerHTML;
        }
        if (this.placeholder) {
            element.placeholder = this.placeholder;
        }
        if (this.type) {
            element.type = this.type;
        }
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
}

class NewLink {
    constructor(path, text) {
        this.path = path;
        this.text = text;
    }

    newLink() {
        const link = document.createElement('a');
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