# Chat Room
This is a simple chat room written in [**Go**](https://golang.org/).

Database implemented using [gorm](https://gorm.io/) (An ORM for Go).
# ![Chatroom Example App](Output)

## Some Features:
1. Sign up using unique username, password and email
2. Sign in using username and password
3. Send global messages to everyone in chatroom
4. Send Private messages to a user
5. Share files to groups, users or everyone in the chatroom
6. Create groups, add/remove users to/from group and send messages to group members  
 
## Commands
### Entrance
To sign-up: 
`username, password, email`
(email might be empty but don't forget the comma)

To sign-in: 
`username, password`
### Change info
`change>username, password, email` you can leave at most 2 of them empty but don't forget the commas
### Private Message
`username> message`
### Global Message
`all> message` 
### Share File
`file> username/all> path`
`file> gp> groupname> path`
### Get File
`get> filename`
### New Group
`new> groupname`
### Add User To Group
`add gp> groupname> username`
### Remove User From Group
`rm gp> groupname> username`
### Send Message To Group
`gp> groupname> message`

## Models

### User
|ID|Username|Password| Email | Groups |  Friends  |
|:----:|:----:|:---------:|:---------:|:---------:|:---------:|
|uint| string |  string |  string |  []string | []string     |

### Group
|  Name  | Admin | Members | 
|:----:|:---------:|:---------:|
| string |  uint |  []uint |

## Directory Structure
```
.
├── client
│   ├── cmd
│   │   └── root.go
│   ├── downloads
│   ├── handler
│   │   ├── client.go
│   │   ├── entrance.go
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
│   ├── change_info.go
│   ├── file.go
│   ├── group.go
│   ├── private_message.go
│   ├── request.go
│   ├── sign_in.go
│   └── sign_up.go
├── response
│   ├── change_info.go
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
    │   ├── change_info.go
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

## TODO
- [ ] More efficient error handling
- [ ] Improve worker pool
- [ ] Implement friends feature for users
- [ ] Add profile photo for users
- [ ] Improve gui
