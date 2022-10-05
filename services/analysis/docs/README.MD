# API USER DOCUMENTATION

## REQUIRMENTS
+ python3
+ mysql

## Step by step to run the code
+ create ```.env``` file you can follow what you need on ```.env-example``` file
+ Create Database ```analysis``` and table ```Information``` manually

+ run testing before. Go to folder ```app```
    ```
        python3 test.py
    ```
    if all of test pass you can use this app.
+ run this app on folder ```app```
```
    python3 app.py
```
<br><br>
-----------

## ENDOPOINTS

```
    Base url : http://localhost:5000/api/v1/analysis
```
<br>

---
+ GET /

Request
``` 
    Headers : None
    Params  : None
    Body    : None
```
Response
```
    array of object
```
---
<br>

+ POST /

Request
``` 
    Headers : None
    Params  : None
    Body    : Json  {
                        user_id: int, 
                        category:string,
                        tags : string,
                        date : "yyyy-mm-dd"
                        duration : int
                        total_duartion : int
                    }
```
Response
```
    string
```
---

<br>

+ GET /<id>

Request
``` 
    Headers : None
    Params  : int(id)
    Body    : None
```
Response
```
    Json { category : float, tag : int }
```