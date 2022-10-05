import json
import pymysql
from flask import Blueprint,abort,request
from http import HTTPStatus
from db.connect import Database
from helper.responseMsg import get_response_msg
from helper.dataToObject import dataToObject

analyzedData = Blueprint("Analysis", __name__,url_prefix='/api/v1/analysis')

db = Database()

@analyzedData.route('/',methods=['GET','POST'])
def manageData():
    if request.method == 'GET':
        try :
            data = db.run_query('SELECT * FROM information')
            payload = dataToObject(data)
            response = get_response_msg(payload,HTTPStatus.OK)
            return response
        except pymysql.MySQLError as sqle:
            abort(HTTPStatus.INTERNAL_SERVER_ERROR, describe=str(sqle))
    elif request.method == 'POST':
        try:
            payload = json.loads(request.data)
            query = (
                f'''
                INSERT into information (user_id,category,tags,created_at,duration,total_duration)
                values
                (
                    "{payload['user_id']}",
                    "{payload['category']}",
                    "{payload['tags']}",
                    "{payload['date']}",
                    "{payload['duration']}",
                    "{payload['total_duration']}"
                );
                '''
            )
            db.run_query(query)
            return get_response_msg({"message":"success add payload"},HTTPStatus.CREATED)
        except pymysql.MySQLError as sqle:
            abort(HTTPStatus.INTERNAL_SERVER_ERROR, describe=str(sqle)) 
        except Exception as e:
            abort(HTTPStatus.BAD_REQUEST, description=str(e))

@analyzedData.route('/<id>')
def getAbalizedData(id):
    try : 
        query = f'''SELECT * FROM information WHERE user_id={id} ORDER BY id DESC LIMIT 20'''
        data = db.run_query(query)
        if len(data) < 20 :
            raise Exception("data must be more than 10 rows")
        payload = dataToObject(data)

        categorys = {}
        tags = {}

        '''Find the persentation of watch video'''
        for item in payload:
            persentation = round((int(item['duration']) / int(item['total_duration'])) * 100,4)
            item['persentation'] = persentation

            if item['category'] in categorys :
                categorys[item['category']] = round(categorys[item['category']] + item['persentation'],3)
            else :
                categorys[item['category']] = round(item['persentation'],4)

            splitTags = item['tags'].split(",")
            
            for targetTag in splitTags:
                if targetTag in tags :
                    tags[targetTag] += 1
                else :
                    tags[targetTag] = 1

        '''Result of category and tags that recomended'''
        result = {}

        max = 0
        for tag in tags :
            if tags[tag] > max:
                max = tags[tag]
                result['tag'] = { tag : max }

        max = 0
        for category in categorys :
            if categorys[category] > max:
                max = categorys[category]
                result['category'] = { category : max }

        response = get_response_msg(result,HTTPStatus.OK)
        return response
    except pymysql.MySQLError as sqle:
        abort(HTTPStatus.INTERNAL_SERVER_ERROR, describe=str(sqle))
    except Exception as e :
        if e.args[0] == 'data must be more than 10 rows':
            abort(HTTPStatus.LENGTH_REQUIRED,description=e.args[0])
        abort(HTTPStatus.BAD_REQUEST, description=str(e))
