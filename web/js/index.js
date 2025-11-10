const roomId = window.location.pathname.slice(1);

async function getTodos() {
  const res = await fetch(`/api/${roomId}/todos`);
  const todos = await res.json();
  return todos;
}

async function addTodo(todo) {
  await fetch(`/api/${roomId}/todos`, {
    method: "POST",
    body: todo,
  });
}

const form = document.getElementById("todoForm");

form.addEventListener("submit", async (e) => {
  e.preventDefault();

  const data = new FormData(form);

  await addTodo(data);

  const todos = await getTodos();

  renderTodos(todos);
});

window.onload = async () => {
  const todos = await getTodos();

  renderTodos(todos);
};

function renderTodos(todos) {
  const list = document.getElementById("todoList");
  list.innerHTML = "";
  todos.forEach((todo) => {
    const item = createTodoCardItem(todo);

    list.appendChild(item);
  });
}

function createTodoCardItem(todo) {
  console.log(todo);
  const li = document.createElement("li");

  const cardBase = document.createElement("div");
  cardBase.className = "todo-card";

  const title = document.createElement("h3");
  const completedCheckbox = document.createElement("input");
  const timestampFooter = document.createElement("span");

  title.className = "todo-title";
  title.innerText = todo.title;

  completedCheckbox.type = "checkbox";
  completedCheckbox.className = "todo-completed";
  completedCheckbox.checked = todo.completed;

  timestampFooter.className = "todo-timestamp";

  timestampFooter.textContent = `Created ${todo.CreatedAt} ${
    todo.UpdatedAt ? `Updated: ${todo.UpdatedAt}` : ""
  }`;

  cardBase.append(title, completedCheckbox, timestampFooter);

  li.appendChild(cardBase);
  return li;
}
