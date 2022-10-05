def dataToObject(data):
    payload = []
    for i in range(len(data)) :
        payload.append({
            "id" : data[i][0],
            "user_id" : data[i][1],
            "category" : data[i][2],
            "tags" : data[i][3],
            "date" : data[i][4],
            "duration" : data[i][5],
            "total_duration": data[i][6]
        })
    return payload