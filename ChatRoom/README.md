## Chat Room
This is a simple chat room written in [**Go**](https://golang.org/).

Database implemented using [gorm](https://gorm.io/) (An ORM for Go).
# ![Chatroom Example App](Output)

### Some Features:
1. Sign up using unique username, password and email
2. Sign in using username and password
3. Send global messages to everyone in chatroom
4. Send Private messages to a user
5. Share files globally or privately
 
 
### Directory Structure
```
├── client
│   ├── cmd
│   │   └── root.go
│   ├── download
│   ├── handler
│   │   ├── client.go
│   │   ├── file.go
│   │   ├── global_message.go
│   │   ├── message.go
│   │   ├── sign.go
│   │   └── write_file.go
│   ├── main.go
│   └── view
│       └── layout.go
├── config
│   ├── config.go
│   └── default.go
├── db
├── Dockerfile
├── go.mod
├── go.sum
├── myDB.db
├── pkg
│   └── error.go
├── request
│   ├── file.go
│   ├── private_message.go
│   ├── request.go
│   ├── sign_in.go
│   └── sign_up.go
├── response
│   ├── file.go
│   ├── global_message.go
│   ├── private_message.go
│   ├── response.go
│   └── sign.go
└── server
    ├── cmd
    │   └── root.go
    ├── db
    │   └── db.go
    ├── handler
    │   ├── client.go
    │   ├── message.go
    │   ├── read_file.go
    │   ├── sign_in.go
    │   ├── sign_up.go
    │   └── write_file.go
    ├── main.go
    ├── model
    │   ├── client.go
    │   ├── group.go
    │   └── user.go
    └── storage
        ├── 1605333429_Operating_System_Concepts_10th_Edition.pdf
        └── 1605335245_Operating_System_Concepts_10th_Edition.pdf
```

