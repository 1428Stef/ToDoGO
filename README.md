<img src="img.png" width="1000px">

# ToDoGO
This is the best tool for controlling the gardens We organize daily projects, projects and people.

## Quick start 

### 1. Installation 
```bash
git clone https://github.com/1428Stef/ToDoGO.git
cd ToDoGO/src/ToDo
```
### 2. Settings
Create a `ToDoGO/src/ToDo/stotage/storage.json`:
```json
[
  {
    "title": "hello",
    "mark": false,
    "date": "2026-03-27 14:01:24",
    "id": 912864339711
  }
]
```
### 3. Start 
**CLI:** 
```bash
make run-cli ARGS="add -title hello"
```

#### Commands

| Command | Flag | Description |
|---------|------|-------------|
| `add`  | `--title` | Create a new task |
| `list` | — | Show all tasks |
| `done` | `--id` | Mark a task as completed |
| `del`  | `--id` | Delete a task |
| `edit` | `--id`, `--title` | Edit task title |
| `help` | — | Show all commands |
---
**API-server:**
```bash
make run-api
```
**API:** `http://localhost:9091`

#### Endpoints
 
| Method | Endpoint | Body | Description |
|--------|----------|------|-------------|
| `GET`    | `/list` | — | Get all tasks (JSON) |
| `POST`   | `/add`  | Task title (plain text) | Create a task |
| `PATCH`  | `/done` | Task ID (plain text) | Mark as done |
| `DELETE` | `/del`  | Task ID (plain text) | Delete a task |
---
**License:**
MIT
