from flask import Flask, request
from flask_sqlalchemy import SQLAlchemy
from flask_marshmallow import Marshmallow
from flask_restful import Api, Resource


app = Flask(__name__)

app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///proxypax.db'
db = SQLAlchemy(app)
ma = Marshmallow(app)
api = Api(app, prefix="/api/v1")


class Post(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    url = db.Column(db.String(2048))
    reg = db.Column(db.String(1024))

    def __repr__(self):
        return '<Post %s>' % self.url


class PostSchema(ma.Schema):
    class Meta:
        fields = ("url", "reg")


post_schema = PostSchema()
posts_schema = PostSchema(many=True)


class User(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    token = db.Column(db.String(255))
    email = db.Column(db.String(255))

    def __repr__(self):
        return '<User %s>' % self.token


class UserSchema(ma.Schema):
    class Meta:
        fields = ("token", "email")


user_schema = UserSchema()


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
        except Exception:
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


api.add_resource(ScrapeURL, '/list')

if __name__ == '__main__':
    app.run(debug=True)
