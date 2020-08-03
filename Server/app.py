from flask import Flask, request
from flask_restful import Resource, Api
from os import environ
import jwt


app = Flask(__name__)
api = Api(app, prefix="/api/v1")

app.config['SECRET_KEY'] = environ['JWT_SECRET']

def auth_required(f):
    def decorated(*args):
        token = None

        if 'x-access-token' in request.headers:
            token = request.headers['x-access-token']

        if not token:
            return {'message':'Token is missing'}, 401

        try:
            ## VALIDATION NEEDED!!! login and validate token.
            return {'message':'youre all good'}, 200
        except:
            return {'message':'Invalid Token'}, 401
        return f( *args)
    
    return decorated

class root(Resource):
    @auth_required
    def get(self):
        return {"test":"object"}

api.add_resource(root, '/')

if __name__ == '__main__':
    app.run(debug=environ['DEBUG'])
