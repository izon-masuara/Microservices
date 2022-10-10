# API USER DOCUMENTATION

## REQUIRMENTS
+ golang
+ postgresql

## Step by step to run the code
+ create ```.env``` file you can follow what you need on ```.env-example``` file
+ Create Database ```User_Movie``` and ```User_Movie_Test``` manually
+ go to folder migration and run user.go
    ```
        go run user.go 
    ```
+ run testing before. Go to folder ```__test__/cmd```
    ```
        go test
    ```
    if all of test pass you can use this app.
+ run this app on folder ```user```
```
    go run main.go
```
<br><br>
-----------

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