from http import HTTPStatus
from flask import Flask
from db.connect import Database
from analysis import analyzedData
from helper.responseMsg import get_response_msg

baseUrl = "/api/v1"

def create_app():
    app = Flask(__name__)
    app.register_blueprint(analyzedData)
    return app


app = create_app()
db = Database()

## ================================[ Error Handler Defined - Start ]
## HTTP 404 error handler
@app.errorhandler(HTTPStatus.NOT_FOUND)
def page_not_found(e):    
    return get_response_msg(data=str(e), status_code=HTTPStatus.NOT_FOUND)


## HTTP 400 error handler
@app.errorhandler(HTTPStatus.BAD_REQUEST)
def bad_request(e):
    return get_response_msg(str(e), HTTPStatus.BAD_REQUEST)


## HTTP 500 error handler
@app.errorhandler(HTTPStatus.INTERNAL_SERVER_ERROR)
def internal_server_error(e):
    return get_response_msg(str(e), HTTPStatus.INTERNAL_SERVER_ERROR)
## ==================================[ Error Handler Defined - End ]

if __name__ == '__main__':
    ## Launch the application 
    app.run(debug=True)