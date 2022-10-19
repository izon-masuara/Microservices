import json
from pydoc import describe
from flask import request
from flask_restful import Resource, abort
from model import Infromation
from helper.dataToObject import dataToObject
from helper.responsemsg import get_response_msg
from init import db
from sqlalchemy import null, exc

# reqBody = reqparse.RequestParser(bundle_errors=True,)
# reqBody.add_argument("user_id",type=int, help="user id required", required=True)
# reqBody.add_argument("category",type=str, help="category required", required=True)
# reqBody.add_argument("tags",type=str, help="tags required", required=True)
# reqBody.add_argument("duration",type=int, help="duration required", required=True)
# reqBody.add_argument("total_duration",type=int, help="total_duration required", required=True)

class InputData(Resource):
    def get(self):
        try:
            resp = Infromation.query.all()
            data = dataToObject(resp)
            return data
        except:
            return "Internal server error", 500
    def post(self):
        try:
            args = json.loads(request.data)
            db.session.add(Infromation(
                args['user_id'],
                args['category'],
                args['tags'],
                args['duration'],
                args['total_duration'],
                null
            ))
            db.session.commit()
            return "Success Added", 201
        except exc.SQLAlchemyError as err:
            return "Internal server error", 500

class Analysis(Resource):
    def get(self,id):
        try:
            resp = Infromation.query.filter_by(user_id=id).all()
            data = dataToObject(resp)
            if len(data) < 20:
                error = '''{
                    "message" : "Data less than 20 rows",
                    "code" : 400
                }'''
                raise Exception(error)

            categorys = {}
            tags = {}

            '''Find the persentation of watch video'''
            for item in data:
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
                    result['tag'] = tag 

            max = 0
            for category in categorys :
                if categorys[category] > max:
                    max = categorys[category]
                    result['category'] = category

            responseResult = get_response_msg(result,200)
            return responseResult
        except Exception as err:
            messageErr = json.loads(str(err))
            if messageErr['message'] == "Data less than 20 rows":
                abort(http_status_code=400, data=str(messageErr['message']), status=str(messageErr['code']))
            else:
                return "Internal server error", 500
