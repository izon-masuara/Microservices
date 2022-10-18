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
    Json { category : String, tag : String }
```