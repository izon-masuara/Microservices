# API USER DOCUMENTATION

## ENDOPOINTS  

```
    Base url : http://localhost:3000/api/v1/user
```
<br>

---
+ POST /login

Request
``` 
    Headers : None
    Params  : None
    Body    : Json { username :string, password : string }
```
Response
```
    string
```
---
<br>

+ POST /register

Request
``` 
    Headers : None
    Params  : None
    Body    : Json { username: string, password:string }
```
Response
```
    string
```
---

<br>

+ POST /token

Request
``` 
    Headers : None
    Params  : None
    Body    : Json { accessToken : string }
```
Response
```
    Json { user_id : int, username : string }
```