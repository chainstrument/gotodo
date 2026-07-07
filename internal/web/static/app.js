const form = document.getElementById("new-todo-form");
const titleInput = document.getElementById("title");
const descriptionInput = document.getElementById("description");
const list = document.getElementById("todo-list");

async function loadTodos() {
  const res = await fetch("/todos");
  const todos = await res.json();
  renderTodos(todos);
}

function renderTodos(todos) {
  list.innerHTML = "";
  for (const todo of todos) {
    list.appendChild(renderTodo(todo));
  }
}

function renderTodo(todo) {
  const li = document.createElement("li");
  li.className = todo.done ? "done" : "";

  const checkbox = document.createElement("input");
  checkbox.type = "checkbox";
  checkbox.checked = todo.done;
  checkbox.addEventListener("change", () => toggleDone(todo, checkbox.checked));

  const titleWrap = document.createElement("div");
  titleWrap.className = "title";
  titleWrap.textContent = todo.title;

  if (todo.description) {
    const desc = document.createElement("span");
    desc.className = "description";
    desc.textContent = todo.description;
    titleWrap.appendChild(desc);
  }

  const deleteBtn = document.createElement("button");
  deleteBtn.className = "delete";
  deleteBtn.textContent = "✕";
  deleteBtn.addEventListener("click", () => deleteTodo(todo.id));

  li.append(checkbox, titleWrap, deleteBtn);
  return li;
}

async function createTodo(title, description) {
  await fetch("/todos", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ title, description }),
  });
  await loadTodos();
}

async function toggleDone(todo, done) {
  await fetch(`/todos/${todo.id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ title: todo.title, description: todo.description, done }),
  });
  await loadTodos();
}

async function deleteTodo(id) {
  await fetch(`/todos/${id}`, { method: "DELETE" });
  await loadTodos();
}

form.addEventListener("submit", async (e) => {
  e.preventDefault();
  const title = titleInput.value.trim();
  if (!title) return;
  await createTodo(title, descriptionInput.value.trim());
  titleInput.value = "";
  descriptionInput.value = "";
  titleInput.focus();
});

loadTodos();
