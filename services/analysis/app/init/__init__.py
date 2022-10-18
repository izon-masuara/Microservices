import os
from flask import Flask
from flask_restful import Api
from flask_sqlalchemy import SQLAlchemy
from dotenv import load_dotenv
load_dotenv()

HOST = os.getenv('DB_HOST')
USERNAME = os.getenv('DB_USERNAME')
PASS = os.getenv('DB_PASS')
PORT = os.getenv('DB_PORT')
NAME = os.getenv('DB_NAME')

app = Flask(__name__)
app.config["SQLALCHEMY_DATABASE_URI"] = f"mysql+mysqlconnector://{USERNAME}:{PASS}@{HOST}:{PORT}/{NAME}"
app.config["BUNDLE_ERRORS"] = True
db = SQLAlchemy()
db.init_app(app)
api = Api(app)