This is a Api build using golang and net/http for routing simple get and post user data you can simply run it 
using followiing command on your terminal make sure golang is installed 
```
go run main.go
```
and test it postman by hitting 
Post request on 
```
http://localhost:3000/user
```
sample request 
```
{
    "Name" : "Md Amaan",
    "ID" : 2,
    "valid": false
}
```
GET request on 
id = Whatever integer you filled in ID field above post request which got registered 
```
http://localhost:3000/user/id
```
Response 
{Name: Md Amaan, ID : 2}
