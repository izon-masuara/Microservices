from random import randint
import requests
import json
categoryArr = ["Programming","Memasak","Permainan","Pekerjaan","Pelajaran","Peralatan"]
tagsArr = ["Python","Golang","Nodejs","Kantor","Mobile","PC","Desktop","Mongodb"]
titleArr = ["Belajar coding","Memasak Ikan","Bermain mobile legend","Bermain minecraft","Laptop gaming 2022"]
descriptionArr = [
    "Bermain moile legend menggunakan wifi yang begitu cepat",
    "Belajar microservice menggunakan bahasa pemrograman nodejs, golang dan python",
    "Memasak nasi menggunakan kompor listrik merupakan hal yang menyenangkan",
    "Laptop yang sangat bagus untuk digunakan dalam belajar pemnrogramman"
]
filesNameImgArr = [
    "MapsBali.png",
    "pleaseWait.png",
    "ss.png",
    "test.png",
    "waireframeBapontarAdmin.png"
]
filesNameVideosArr = [
    "coba.mp4",
    "gitar.mp4"
]

username = [
    "Budi",
    "Andreas",
    "OkkyBoy"
]

def storeUser():
    url = "http://localhost:3000/api/v1/user/register"

    payload = json.dumps({
        "username": username[0],
        "password": "admin"
    })
    payload2 = json.dumps({
        "username": username[1],
        "password": "admin"
    })
    payload3 = json.dumps({
        "username": username[2],
        "password": "admin"
    })
    headers = {
    'Content-Type': 'application/json'
    }

    response = requests.request("POST", url, headers=headers, data=payload)
    response = requests.request("POST", url, headers=headers, data=payload2)
    response = requests.request("POST", url, headers=headers, data=payload3)

    print(response.text)

def storeFile():
    url = "http://localhost:3001/api/v1/files/upload"

    generateData = 0
    stopGenerate = 30
    while generateData < stopGenerate :
        generateData += 1
        titleInt = randint(0,len(titleArr)-1)
        descInt = randint(0,len(descriptionArr)-1)
        categorycInt = randint(0,len(categoryArr)-1)
        tagsInt = randint(0,len(tagsArr)-1)
        imgInt = randint(0,len(filesNameImgArr)-1)
        videoInt = randint(0,len(filesNameVideosArr)-1)
        payload={'title': titleArr[titleInt],
        'description': descriptionArr[descInt],
        'category': categoryArr[categorycInt],
        'tags[0]': tagsArr[tagsInt],
        'tags[1]': "Programming",
        'uploadedUserId': '1'}
        files=[
        ('files',(filesNameImgArr[imgInt],open('images/'+filesNameImgArr[imgInt],'rb'),'image/png')),
        ('files',(filesNameVideosArr[videoInt],open('videos/'+filesNameVideosArr[videoInt],'rb'),'video/mp4'))
        ]
        headers = {}
        response = requests.request("POST", url, headers=headers, data=payload, files=files)

        print(response.text)

def getDataFromApiGateawayWithoutLoginBefore():
    url = "http://localhost:3002?query=movies"

    response = requests.request("GET",url=url)
    print(response.json())

def getDataFromApiGateawayLoginBefore():
    urlUser = "http://localhost:3000/api/v1/user/login"
    urlGateaway = "http://localhost:3002?query=movies"
    payload = json.dumps({
        "username": username[2],
        "password": "admin"
    })
    headers = {
    'Content-Type': 'application/json'
    }

    response = requests.request("POST", urlUser, headers=headers, data=payload)


    token = response.json()['accessToken']
    tokenResp = token
    token = json.dumps({
        "accessToken": token
    })
    
    resp = requests.request("GET",url=urlGateaway,data=token)
    print(resp.json(),"<<<<")
    return resp.json(),tokenResp


def postDataToAnalysis():

    data,token = getDataFromApiGateawayLoginBefore()

    url = "http://localhost:3002?mutation=movies"

    print(data)

    count = 0
    while count < 22 :
        payload = json.dumps({
            "accessToken": token,
            "category": categoryArr[randint(0,len(categoryArr)-1)],
            "tags" : [tagsArr[randint(0,len(tagsArr)-1)],"Programming"],
            "duration" : 2000,
            "total_duration" : data[randint(0,len(data)-1)]['Files']['size']
        })

        response = requests.request("GET",url=url,data=payload)
        count += 1
        print("success")

storeUser()
storeFile()
getDataFromApiGateawayWithoutLoginBefore()
getDataFromApiGateawayLoginBefore()
postDataToAnalysis()
