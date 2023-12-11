### Todo
Simple API for Todo List

#### API
- `GET /todos` to get all todos
- `POST /todos` to add a new todo with the following fields `title`, `description`, `done` as boolean value.
- `PUT /todos/:id` to update a todo by id, accept a json with fields as in `POST /todos`
- `DELETE /todos/:id` to delete a todo by id
