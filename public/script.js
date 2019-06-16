const $ = document.querySelector.bind(document);

waitUntilDomIsReady()
  .then(registerEventHandlers)
  .then(getTodos);

function registerEventHandlers() {
  const createTodoForm = $('.todo-form');
  createTodoForm.addEventListener('submit', handleCreateTodoFormSubmit);
}

function handleCreateTodoFormSubmit(event) {
  event.preventDefault();
  const formData = [...event.target.elements].reduce((formData, formElement) => {
    if (formElement.name) {
      formData[formElement.name] = formElement.value;
    }
    return formData;
  }, {});
  addTodo(formData)
}

function getTodos() {
  fetch('/api/todos')
    .then(response => response.json())
    .then(renderTodos)
    .catch(console.error.bind(console));
}

function addTodo(todoData) {
  optimisticallyRenderTodoItem(todoData);
  const fetchParams = {
    method: 'POST',
    body: JSON.stringify(todoData)
  }
  fetch('/api/todos', fetchParams)
    .then(getTodos)
    .catch(console.error.bind(console));
}

function optimisticallyRenderTodoItem(todoData) {
  const todoList = $('.todo-list');
  const temporaryContainer = document.createElement('div');
  temporaryContainer.innerHTML = createHTML(todoData);
  todoList.insertBefore(temporaryContainer.firstChild, todoList.firstChild)
}

function renderTodos(todos) {
  if (!todos) {
    return
  }
  const html = todos.map(createHTML).join('')
  $('.todo-list').innerHTML = html;
}

function createHTML(todo) {
  return [
    '<li class=".todo-list_item">',
    '  ' + todo.Title + (todo.Done ? '   ✓' : ''),
    '</li>'
  ].join('')
}

function waitUntilDomIsReady() {
  return new Promise(resolve => {
    window.addEventListener('DOMContentLoaded', () => resolve());
  })
}