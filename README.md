# goent-todo
This is a simple todo api for try to something.

# Installation
You can download this project by clone with this command.
<br></br>
`git clone https://github.com/Mustafaakilll/goent-todo.git`

# Run
You have to have `postgresql` for run this project and configure `.env` file to right configuration.
Alternatively you can use `docker` to run postgresql without install it.
- First: `docker compose up` to run postgresql.
- Second: `go run .` to run `go` project.
- Finally you can go to `localhost:3000/api/v1/swagger/index.html` for swagger ui.

# Doc
You can view godoc with `godoc -http:3001` command and in browser you can search this website `http://localhost:3001/pkg/github.com/mustafaakilll/ent_todo/`.
