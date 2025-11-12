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

async function deleteTodo(id) {
  const res = await fetch(`/api/${roomId}/todos/${id}`, {
    method: "DELETE",
  });

  if (res.status == 500) {
    throw new Error("Failed to delete todo item.");
  }
}

// async function updateTodo() {}

const form = document.getElementById("todoForm");

form.addEventListener("submit", async (e) => {
  e.preventDefault();

  const data = new FormData(form);

  await addTodo(data);

  const todos = await getTodos();

  document.getElementById("todoInput").value = "";

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
  cardBase.id = todo.ID;
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

  const buttonContainer = document.createElement("div");

  buttonContainer.className = "btn-container";

  const editButton = document.createElement("button");
  const deleteButton = document.createElement("button");

  editButton.innerText = "Edit";
  editButton.className = "edit-btn";
  editButton.id = generateRandomId();

  deleteButton.innerText = "Delete";
  deleteButton.className = "delete-btn";
  deleteButton.id = generateRandomId();

  buttonContainer.append(editButton, deleteButton);

  editButton.onclick = (event) => {};

  deleteButton.onclick = async (_) => {
    try {
      await deleteTodo(todo.ID);
      cardBase.remove();
    } catch (error) {
      alert(error.message);
    }
  };

  cardBase.append(title, completedCheckbox, timestampFooter, buttonContainer);

  li.appendChild(cardBase);
  return li;
}

function generateRandomId() {
  return Math.random().toString(36).slice(3);
}
