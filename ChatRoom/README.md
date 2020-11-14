# Chat Room
This is a simple chat room written in [**Go**](https://golang.org/).

Database implemented using [gorm](https://gorm.io/) (An ORM for Go).
# ![Chatroom Example App](Output)

## Some Features:
1. Sign up using unique username, password and email
2. Sign in using username and password
3. Send global messages to everyone in chatroom
4. Send Private messages to a user
5. Share files
6. Create groups, add/remove users to/from group and send messages to group members  
 
## Commands
### Private Message
`username> message`
### Global Message
`all> message` 
### Share File
`file> username/all> path`
### New Group
`new> groupname`
### Add User To Group
`add gp> groupname> username`
### Remove User From Group
`rm gp> groupname> username`
### Send Message To Group
`gp> groupname> message`
 
## Directory Structure
```
.
├── client
│   ├── cmd
│   │   └── root.go
│   ├── downloads
│   ├── handler
│   │   ├── client.go
│   │   ├── file.go
│   │   ├── global_message.go
│   │   ├── group.go
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
├── Output
├── pkg
│   └── error.go
├── README.md
├── request
│   ├── file.go
│   ├── group.go
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
    │   ├── group.go
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
```

