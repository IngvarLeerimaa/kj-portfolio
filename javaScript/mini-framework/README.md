# MINI-FRAMEWORK


## Description

Objective of this project is to create a mini-framework that can be used to simplify the creation of web applications.

This **simple framework** will be used to create a simple Todo application, based on [todoMVC](https://todomvc.com/) example.

[Audit question can me found here](https://github.com/01-edu/public/blob/master/subjects/mini-framework/audit/README.md)

## Technologies used

* HTML
* CSS
* JavaScript

## Running the project 

[Just visit the project by clicking here ](https://todo4kood.netlify.app/)

### To run the project locally:

This method requires you to have live-server installed on your machine. If you don't have it installed, you can install it by running `npm install -g live-server` or installing the extentsion on vscode.

1. Run `npm install` to install all the dependencies.
2. Run `npm start` to start the server.
3. Open your browser.

## Usage

To use this framework in your project, you need to follow the steps below:

1. In your index.html file add an element with the id `app` where you want to render your application.
2. Create a new script tag pointing to your main javascript file.

Example:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title> Title </title>
    <link rel="stylesheet" href="/styles.css">
    <link rel="icon" href="/favicon.ico">
</head>
<body>
    <div id="app"></div>
    <script type="module" src="app/index.js"></script>
</body>
</html>
```

3. In your main javascript file, import the framework and create a new instance of the framework using the keyword **simple**.

4. With the **simple** functions you get access to creating online elements, adding event listeners, and adding quick routes with just one line of code. 

Example:
```javascript

import { simple } from './framework/framework.js';

// Example NewElement(possible options: tag, id, class, innerHTML, placeholder)
// Use .create() to initialize the element
const yourNewElement = new simple.NewElement('tag', 'id', 'class', 'innerHTML', 'placeholder').create();

// Example Listen(possible options: element, event, callback)
const newListner = new simple.Listen(element, "keypress", (e) => {

    console.log("You just pressed a key!");
});

//use listen() to initialize the newListener
newListener.listen()

// Example NewLink(possible options: path, text)
const newRoute = new simple.NewLink('https://www.google.com', 'I want to google!').newLink();

// Example of creating a new Router and adding a routes to it
const router = new simple.Router();

router.addRoute('/YOURPATH', () => {
});

router.navigate('/YOURPATH');
```

Created by: [IngvarLeerimaa](https://01.kood.tech/git/IngvarLeerimaa) and [MargusT][def]

[def]: https://01.kood.tech/git/MargusT