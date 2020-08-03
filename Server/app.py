from flask import Flask, request
from flask_restful import Resource, Api
import jwt


app = Flask(__name__)
api = Api(app, prefix="/api/v1")


class root(Resource):

    def get(self):
        return {"test":"object"}

api.add_resource(root, '/')

if __name__ == '__main__':
    app.run(debug=True)
