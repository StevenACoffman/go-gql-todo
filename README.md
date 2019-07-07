# go-gql-todo

Go-gql-todo is a Todo app with the backend written in Go using GraphQL and the frontend written with React (using TypeScript).
The backend began as the example [gqlgen](https://github.com/99designs/gqlgen) code. 

## Run the Application

Run the example, it will spawn a GraphQL HTTP endpoint

```
go run ./scripts/gqlgen.go
go generate ./...
go run ./cmd/todo.go
```
Once this is running, you should be able to open a browser to `http://localhost:8081/playground`
and make a graphql query like:
```
# Write your query or mutation here
query{todo(id:"1"){id,text,done}}
```

Execute queries and mutations via shell.

```
// To get single ToDo item by ID
curl -g 'http://localhost:8081/query?query={todo(id:"1"){id,text,done}}'

// To create a ToDo item
curl -g 'http://localhost:8081/query?query=mutation+_{createTodo(text:"My+new+todo"){id,text,done}}'

// To get a list of ToDo items
curl -g 'http://localhost:8081/query?query={todoList{id,text,done}}'

// To update a ToDo
curl -g 'http://localhost:8081/query?query=mutation+_{updateTodo(id:"1",changes:{text:"My+new+todo+updated",done:true}){id,text,done}}'
```
