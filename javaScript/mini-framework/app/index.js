import simple from "../framework/framework.js";

const App = document.getElementById("app");

let todoList = []; // {task: "task", taskCompleted: false}
addToLocalStorage(); // Initialize the local storage

// NewElement(tag, id, className, innerHTML, placeholder, type)
const container = new simple.NewElement("section", null, "todoapp").create();
const header = new simple.NewElement("header", null, "header").create();
const head = new simple.NewElement("h1", null, null, "Todo || !Todo").create();
const input = new simple.NewElement("input", null, "new-todo", null, "What to do?").create();
const toggleDiv = new simple.NewElement("div", null, "toggle-all-container").create();
const toggleAll = new simple.NewElement("input", null, "toggle-all", null, null, "checkbox").create();
const label = new simple.NewElement("label", null, "toggle-all-label").create();
const edit = new simple.NewElement("input", null, "edit").create();

const main = new simple.NewElement("main", null, "main").create();
const footer = new simple.NewElement("footer", null, "footer").create();
const count = new simple.NewElement("span", "todo-count", "todo-count").create();
const filters = new simple.NewElement("ul", null, "filters").create();
const all = new simple.NewElement("a", "all", "selected", "All").create();
const active = new simple.NewElement("a", "active", null, "Active").create();
const completed = new simple.NewElement("a", null, "completed", "Completed").create();
const clear = new simple.NewElement("button", "clear", "clear-completed", "Clear Completed").create();
main.style.display = "none";
footer.style.display = "none";


const about = new simple.NewElement("footer", null, "info", "Need more info?").create();
about.appendChild(new simple.NewElement("p", null, null, "Double click on a task to edit it").create());
const repoDirect = new simple.NewLink("https://01.kood.tech/git/IngvarLeerimaa/mini-framework", "Click here to see the repo").newLink();
about.appendChild(repoDirect);
about.appendChild(new simple.NewElement("br").create());
const auditDirect = new simple.NewLink("https://github.com/01-edu/public/blob/master/subjects/mini-framework/audit/README.md", "Click here to see audit questions").newLink();
about.appendChild(auditDirect);


// Listen(element, event, callback)
const newTodo = new simple.Listen(input, "keypress", (e) => {
    if (e.key !== "Enter" || !input.value.trim()) {
        return;
    }

    todoList.push({ task: input.value.trim(), taskCompleted: false });
    addToLocalStorage();
    input.value = "";
    render();
});

newTodo.listen();

//Router and path initialization
const router = new simple.Router();

router.addRoute('/active', () => {
    render();
});
router.addRoute('/completed', () => {
    render();
});
router.addRoute('/', () => {
    render();
});

// Filters
const allTasks = new simple.Listen(all, "click", () => {
    todoList = JSON.parse(localStorage.getItem("todoList"));
    all.classList.add("selected");
    active.classList.remove("selected");
    completed.classList.remove("selected");
    router.navigate('/');
});

const activeTasks = new simple.Listen(active, "click", () => {
    todoList = JSON.parse(localStorage.getItem("todoList"));
    todoList = todoList.filter(todo => !todo.taskCompleted);
    all.classList.remove("selected");
    active.classList.add("selected");
    completed.classList.remove("selected");

    router.navigate('/active')
});

const completedTasks = new simple.Listen(completed, "click", () => {
    todoList = JSON.parse(localStorage.getItem("todoList"));
    todoList = todoList.filter(todo => todo.taskCompleted);
    all.classList.remove("selected");
    active.classList.remove("selected");
    completed.classList.add("selected");
    router.navigate('/completed');
    
});

const clearCompleted = new simple.Listen(clear, "click", (e) => {
    todoList = JSON.parse(localStorage.getItem("todoList"));
    todoList = todoList.filter(todo => !todo.taskCompleted);
    localStorage.removeItem("todoList");
    addToLocalStorage();
    count.innerHTML = "";
    render();
});

const toggleAllTasks = new simple.Listen(label, "click", () => {
    todoList = JSON.parse(localStorage.getItem("todoList"));
    const isCompleted = todoList.find(todo => !todo.taskCompleted) === undefined ? false : true;
    todoList.forEach(todo => {
        todo.taskCompleted = isCompleted;
    });
    addToLocalStorage();
    render();
});

// Appending the filters to the footer
const newAll = new simple.NewElement("li").create();
newAll.appendChild(all);
filters.appendChild(newAll);
const newActive = new simple.NewElement("li").create();
newActive.appendChild(active);
filters.appendChild(newActive);
const newCompleted = new simple.NewElement("li").create();
newCompleted.appendChild(completed);
filters.appendChild(newCompleted);

allTasks.listen();
activeTasks.listen();
completedTasks.listen();
clearCompleted.listen();

// Appending to footer
footer.appendChild(count);
footer.appendChild(filters);
footer.appendChild(clear);

header.appendChild(head);
header.appendChild(input);

// Appending to container
container.appendChild(header);

// Toggle all tasks
label.setAttribute("for", "toggle-all");
toggleAllTasks.listen();
toggleDiv.appendChild(toggleAll);
toggleDiv.appendChild(label);
main.appendChild(toggleDiv);                          
container.appendChild(main);
container.appendChild(footer);

App.appendChild(container);
App.appendChild(about);


function render() { 

    removeOld();
    itemsLeft();

    
    const renderList = new simple.NewElement("ul", null, "todo-list").create();

    if (localStorage.getItem("todoList").length > 2 ){
        main.style = "";
        footer.style = "";
        clear.style.display = "none";
    } else {
        main.style.display = "none";
        footer.style.display = "none";
    }

    todoList.forEach((todo, index) => {
        const todoItem = new simple.NewElement("li", index).create();
        const todoItemDiv = new simple.NewElement("div", null, "view").create();
        const todoItemLabel = new simple.NewElement("label", null, null, todo.task).create();
        const todoItemInput = new simple.NewElement("input", null, "toggle", null, null, "checkbox").create();
        const todoItemButton = new simple.NewElement("button", null, "destroy").create();


        if (index === 0) {
            todoItem.id = 0;
        }
        if (todo.taskCompleted) {
            todoItem.classList.add("completed");
            todoItemInput.checked = true;
            clear.style = "";
            }
        

        todoItemDiv.appendChild(todoItemInput);
        todoItemDiv.appendChild(todoItemLabel);
        todoItemDiv.appendChild(todoItemButton);
        todoItem.appendChild(todoItemDiv);
        renderList.appendChild(todoItem);
        

        todoItemLabel.addEventListener("dblclick", () => {
            const editInput = new simple.NewElement("input", null, "edit", todo.task).create();
            todoItem.classList.add("editing");
            editInput.value = todo.task;
            todoItem.appendChild(editInput);
            todoItemDiv.remove()
            editInput.focus();

            editInput.addEventListener("keypress", (e) => {
                if (e.key !== "Enter" ) {
                    return;
                }
                todo.task = editInput.value.trim();
                addToLocalStorage();
                render();
            });

            editInput.addEventListener("focusout", () => {
                if (editInput.value.trim() === "") {
                    render();
                    return
                }
                todo.task = editInput.value.trim();
                addToLocalStorage();
                render();
            });
        });

        //Marks the task in ui and changes the status in local storage
        todoItemInput.addEventListener("click", () => {
            todo.taskCompleted = !todo.taskCompleted;
            changeStatus(todo, index);
            itemsLeft();
            render();
        });

        //Deletes the task from the list and local storage
        todoItemButton.addEventListener("click", () => {
            todoList.splice(index, 1);
            removeFromLocal(todo);
            addToLocalStorage();
            render();
        });
    });

    main.appendChild(renderList);
}

function addToLocalStorage() {
    localStorage.setItem("todoList", JSON.stringify(todoList));
}

function removeOld() {
    const old = container.querySelector(".todo-list");
    if (old) {
        main.removeChild(old);
    }
}

function changeStatus(obj, i) {
    let list = JSON.parse(localStorage.getItem("todoList"));
    list.forEach((todo, index) => { 
        if (todo.task === obj.task && i === index) {
            todo.taskCompleted = !todo.taskCompleted;
        }
    });

    localStorage.setItem("todoList", JSON.stringify(list));
}

function itemsLeft() {
    let buf = JSON.parse(localStorage.getItem("todoList"));
    if (buf === 0) {
        count.innerHTML = "";
        return;
    }

    let c = 0;
    buf.forEach(todo => {
        if (!todo.taskCompleted) {
            c++;
        }
    });

    count.innerHTML = `${c} ${c == 1 ? "item" : "items"} left`;
    return;

}

function removeFromLocal(obj) {
    let list = JSON.parse(localStorage.getItem("todoList"));
    list = list.filter(todo => todo.task !== obj.task);
    localStorage.setItem("todoList", JSON.stringify(list));
}