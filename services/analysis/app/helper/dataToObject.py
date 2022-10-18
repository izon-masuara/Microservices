def dataToObject(data):
    payload = []
    for item in data:
        payload.append({
            "id" : item.id,
            "user_id": item.user_id,
            "category": item.category,
            "tags": item.tags,
            "duration": item.duration,
            "total_duration": item.total_duration,
            "date": str(item.date),
        })
    return payload