# Go-Todo

Go-Todo is a Todo app with the backend written in Go using GraphQL and the frontend written with React (using TypeScript)

## Reason for another todo app

Well I want to increase my skills in writing Go and to do that writing a todo app is probably the best way

## Features of the app

| Features                                             | Done | Version |
| ---------------------------------------------------- | ---- | ------- |
| User Register/Login                                  |      |         |
| Create different types of todos                      |      |         |
| Share a unique link of the todos                     |      |         |
| Ability to order remaining todos                     |      |         |
| Ability to set a reminder for the todos              |      |         |
| Ability to set due-date for the todos                |      |         |
| Standard todo app abilities (like edit, remove etc.) |      |         |


## Run the Application

Run the example, it will spawn a GraphQL HTTP endpoint

```
go run main.go
```

Execute queries and mutations via shell.

```
// To get single ToDo item by ID
curl -g 'http://localhost:8081/graphql?query={todo(id:"1"){id,text,done}}'

// To create a ToDo item
curl -g 'http://localhost:8081/graphql?query=mutation+_{createTodo(text:"My+new+todo"){id,text,done}}'

// To get a list of ToDo items
curl -g 'http://localhost:8081/graphql?query={todoList{id,text,done}}'

// To update a ToDo
curl -g 'http://localhost:8081/graphql?query=mutation+_{updateTodo(id:"1",text:"My+new+todo+updated",done:true){id,text,done}}'
```

# go-gql-todo
