# crewe

This project is a kanban-like task management system. 

This project should *not* be deployed on public docker as it doesn't have any protection against SQL injection and the whole project is meant for me to learn how to code in Go.

## endpoints

``` /api/v1 ``` is the main entrypoint.

```POST /tasks``` requires ```"Name": Name, "Status": Status, "ProjectID":id, "assignedTo": id, "description": Text(Optional) ```

GET and DELETE on ```/tasks{id}``` requires the ```id``` of the task.

```POST /projects``` requires ```"Name": Name```

GET and DELETE on ```/projects/{id}``` requires the ```id``` of the project.

```POST /users/register``` requires ```"Name": Name, "email": Email, "firstName": FirstName, "lastName": LastName, "password": Pwd```

A cookie is returned as the login token, all other commands (except for registering) requires the token
