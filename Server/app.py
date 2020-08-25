from flask import Flask, request
from flask_sqlalchemy import SQLAlchemy
from flask_marshmallow import Marshmallow
from flask_restful import Api, Resource
from os import environ

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///test.db'
db = SQLAlchemy(app)
ma = Marshmallow(app)
api = Api(app)


class Post(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    url = db.Column(db.String(255))
    reg = db.Column(db.String(255))

    def __repr__(self):
        return '<Post %s>' % self.url


class PostSchema(ma.Schema):
    class Meta:
        fields = ("id", "url", "reg")


post_schema = PostSchema()
posts_schema = PostSchema(many=True)


def auth_required(f):
    def decorated(*args):
        token = None

        if 'x-access-token' in request.headers:
            token = request.headers['x-access-token']

        if not token:
            return {'message': 'Token is missing'}, 401

        try:
            # VALIDATION NEEDED!!! login and validate token.
            # return {'message': 'youre all good'}, 200
            print("passed")
        except:
            return {'message': 'Invalid Token'}, 401
        return f(*args)

    return decorated


class ScrapeURL(Resource):
    @auth_required
    def get(self):
        posts = Post.query.all()
        return posts_schema.dump(posts)

    def post(self):
        new_post = Post(
            url=request.json['url'],
            reg=request.json['reg']
        )
        db.session.add(new_post)
        db.session.commit()
        return post_schema.dump(new_post)


api.add_resource(ScrapeURL, '/s')

if __name__ == '__main__':
    app.run(debug=environ['DEBUG'])
