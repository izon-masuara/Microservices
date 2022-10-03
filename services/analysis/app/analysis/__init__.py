import json
import pymysql
from flask import Blueprint,abort,request
from http import HTTPStatus
from db.connect import Database
from helper.responseMsg import get_response_msg

analyzedData = Blueprint("Analysis", __name__,url_prefix='/api/v1/analysis')

db = Database()

@analyzedData.route('/',methods=['GET','POST'])
def manageData():
    '''
        id pada table information digunakan sebagai indexing

        kelopokan data berdasarkan id yang login

        cari tanggal dengan 10 video terkahir terakhir  (sebenarnya ini tidak terlalu di butuhkan kecuali untuk analisis total pengunjung dalam bulan / hari / tahun)

        cek waktu tayang terlama sampai tependek

        dan simpulkan berdasarkan catergory dan tags

        analisis berdasarkan waktu tayang dan category serta tags apa yang paling baik untuk di rekomendasikan

        harusnya di api files harus memberikan total durasi video untuk di cek dengan duration yang ada di analitic apakah sudah 100% ?

        tambahakn filed duration video pada files dan anaysis service
    '''
    if request.method == 'GET':
        try :
            data = db.run_query('SELECT * FROM information')
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
            response = get_response_msg(payload,200)
            return response
        except pymysql.MySQLError as sqle:
            abort(HTTPStatus.INTERNAL_SERVER_ERROR, describe=str(sqle))
    elif request.method == 'POST':
        try:
            payload = json.loads(request.data)
            query = (
                f'''
                INSERT into information
                values
                (
                    "{payload['id']}",
                    "{payload['user_id']}",
                    "{payload['category']}",
                    "{payload['tags']}",
                    "{payload['date']}",
                    "{payload['duration']}"
                    # total duration
                );
                '''
            )
            db.run_query(query)
            return get_response_msg({"message":"success add payload"},201)
        except pymysql.MySQLError as sqle:
            abort(HTTPStatus.INTERNAL_SERVER_ERROR, describe=str(sqle)) 
        except Exception as e:
            abort(HTTPStatus.BAD_REQUEST, description=str(e))