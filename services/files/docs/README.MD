# API USER DOCUMENTATION

## REQUIRMENTS
+ nodejs 
+ npm
+ mongodb

## Step by step to run the code

+ run this app on folder ```files```
```
    npm install
```
```
    npm run dev
```
<br><br>
-----------

## ENDOPOINTS

```
    Base url : http://localhost:3001/api/v1/files
```
<br>

---
+ POST /upload

Request
``` 
    Headers : None
    Params  : None
    Body    : form data { 
                title           : string,
                description     : string,
                category        : string,
                tags            : array,
                uploadedUserId  : number,
                files           : [thubmnail,video],
            }
```
Response
```
    string
```
---
<br>

+ GET /

Request
``` 
    Headers : None
    Params  : None
    Body    : None
```
Response
```
    Array of object
```
---

<br>

+ GET /image

Request
``` 
    Headers : None
    Params  : id
    Body    : None
```
Response
```
    Buffer image
```
---

<br>

+ GET /video

Request
``` 
    Headers : None
    Params  : id
    Body    : None
```
Response
```
    Buffer video
```