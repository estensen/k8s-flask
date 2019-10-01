import logging
import requests
import time

from flask import abort, Flask, jsonify
from jaeger_client import Config

app = Flask(__name__)

books = [
    {
        'id': 0,
        'title': 'The Noma Guide to Fermentation',
        'author': 'Rene Redzepi,  David Zilber',
        'published': '2018'
    },
    {
        'id': 1,
        'title': 'The Design of Everyday Things',
        'author': ' Donald A. Norman',
        'published': '2013'
    },
    {
        'id': 2,
        'title': 'Mastering Bitcoin',
        'author': ' Andreas M. Antonopoulos',
        'published': '2014'
    }
]


def slow_func():
    time.sleep(2)


@app.route('/', methods=['GET'])
def index():
    with tracer.start_span('index') as span:
        slow_func()
        r = requests.get('http://host.docker.internal:8080/club')
        if r.status_code == 200:
            return '''
            <h1>Reading Club</h1>
            <a href='/api/v1.0/books'>Books</a>
            '''
        else:
            abort(500)


@app.route('/healthz', methods=['GET'])
def readiness():
    return ""


@app.route('/api/v1.0/books', methods=['GET'])
def get_books():
    with tracer.start_span('get-books') as span:
        return jsonify(books)


@app.route('/api/v1.0/books/<int:book_id>', methods=['GET'])
def get_book(book_id):
    with tracer.start_span('get-book') as span:
        book = [book for book in books if book['id'] == book_id]
        if len(book) == 0:
            abort(404)
        return jsonify({'book': book[0]})


def init_tracer(service):
    log_level = logging.DEBUG
    logging.getLogger('').handlers = []
    logging.basicConfig(format='%(asctime)s %(message)s', level=log_level)

    config = Config(
        config={  # usually read from some yaml config
            'sampler': {
                'type': 'const',
                'param': 1,
            },
            'logging': True,
        },
        service_name=service,
    )
    return config.initialize_tracer()


if __name__ == '__main__':
    tracer = init_tracer('books')
    app.run(host="0.0.0.0", debug=True)
