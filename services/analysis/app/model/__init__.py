from datetime import datetime
from init import db

class Infromation(db.Model):
    id = db.Column('info_id', db.Integer, primary_key = True)
    user_id = db.Column(db.Integer)
    category = db.Column(db.String(50))  
    tags = db.Column(db.String(200))
    duration = db.Column(db.Integer)
    total_duration = db.Column(db.Integer)
    date = db.Column(db.DateTime)

    def __init__(self, user_id, category, tags, duration, total_duration, date):
        self.user_id = user_id
        self.category = category
        self.tags = tags
        self.duration = duration
        self.total_duration = total_duration
        self.date = datetime.now()
    


# INSERT into information (user_id,category,tags,created_at,duration,total_duration)
#                 values
#                 (
#                     "{payload['user_id']}",
#                     "{payload['category']}",
#                     "{payload['tags']}",
#                     "{payload['date']}",
#                     "{payload['duration']}",
#                     "{payload['total_duration']}"
#                 );
#                 '''