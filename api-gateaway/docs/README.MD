# API USER DOCUMENTATION

## REQUIRMENTS
+ python3
+ mysql

## Step by step to run the code
```
    docker-compose up
```
<br>
-----------

## ENDOPOINTS

```
    Base url : http://localhost:3002
```
<br>

---
+ GET /?query=movies

Request
``` 
    Headers : None
    Params  : None
    Body    : {"accessToken" : String } || Optional
```
Response
```
    array of object
```
---
<br>
+ GET /?mutation=movies

Request
``` 
    Headers : None
    Params  : None
    Body    : {
        "accessToken": String,
        "category": String,
        "tags" : Array Of String,
        "duration" : Integer,
        "total_duration" : Integer
    }
```
Response
```
    array of object
```
---