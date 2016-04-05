# GoApiDemo
This is a golang restful api demo using **mux&amp;xorm**.

## Install dependencies
```
go get github.com/gorilla/mux
go get github.com/go-xorm/xorm
go get github.com/mattn/go-sqlite3
```

Or using **godep**:
`godep go install`

## Run
`go run *.go`

## TestCase
### Create user
```
$curl -XPOST -d '{"name":"Alice"}' "http://localhost:80/users"
{
    "id": "11231244213",
    "name": "Alice" ,
    "type": "user"
}
```
### List all users
```
$curl -XGET "http://localhost:80/users"
[
  {
    "id": "21341231231",
    "name": "Bob" ,
    "type": "user"
  },
  {
    "id": "31231242322",
    "name": "Samantha" ,
    "type": "user"
  }
]
```
### Create relationship(PUT)
If two users have "liked" each other, then the state of the relationship is "matched".

**Notes:** If "matched" is changed to "disliked", the one not modified will remain "matched".
```
$curl -XPUT -d '{"state":"liked"}'
"http://localhost:80/users/11231244213/relationships/21341231231"
{
    "user_id": "21341231231",
    "state": "liked" ,
    "type": "relationship"
}
$curl -XPUT -d '{"state":"liked"}'
"http://localhost:80/users/21341231231/relationships/11231244213"
{
    "user_id": "11231244213",
    "state": "matched" ,
    "type": "relationship"
}
$curl -XPUT -d '{"state":"disliked"}'
"http://localhost:80/users/21341231231/relationships/11231244213"
{
    "user_id": "11231244213",
    "state": "disliked" ,
    "type": "relationship"
}
```
### List relationship
```
$curl -XGET "http://localhost:80/users/11231244213/relationships"
[
  {
    "user_id": "222333444",
    "state": "liked" ,
    "type": "relationship"
  },
  {
    "user_id": "333222444",
    "state": "matched" ,
    "type": "relationship"
  },
  {
    "user_id": "444333222",
    "state": "disliked" ,
    "type": "relationship"
  }
]
```

## Database structure
```
//User
type User struct {
    Id   int64  `json:"id"`
    Name string `xorm:"unique not null" json:"name"`
    Type string `json:"type" xorm:"not null default user"`
}
//Relationship
type Relationship struct {
    Id       int64  `json:"id"`
    Owner_id int64  `json:"-" xorm:"not null unique(a)"`
    User_id  int64  `json:"user_id" xorm:"not null unique(a)"`
    State    string `json:"state" xorm:"not null"`
    Type     string `json:"type" xorm:"not null default relationship"`
}
```
Database will be auto synced to goapidemo.db in the current dir.
Now only sqlite is supported.

## TODO
1. Only support sqlite for now
2. **No foreign key for now**, and no check in the api
3. No TestCases for now