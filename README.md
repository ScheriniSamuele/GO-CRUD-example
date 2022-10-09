# Go CRUD application

## Goal

I tried to build a simple CRUD API as a fist easy project with Go language and the GIN framework.

## Folders structure

| Folder          | Description                                                      |
| --------------- | ---------------------------------------------------------------- |
| **routes**      | Contains all the routes and their methods.                       |
| **data**        | Stores the array of products and their structure.                |
| **controllers** | Contains the functions invoked by the routes.                    |
| **helpers**     | Contiens generic functions shared by multiple files in the code. |

## More info

The applications doesn't use a DB for storing data and does not use the SESSION storage.
It is just a trivial and simple application I have done has my FIRST Go project.

The makeFile is used to run the project with nodemon which restart the server on code save.
After installing nodemon globally run the command : `make run`
