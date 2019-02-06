from flask import abort, Flask, jsonify

app = Flask(__name__)

books = [
    {
        'id': 0,
        'title': 'The Noma Guide to Fermentation',
        'author': 'Rene Redzepi,  David Zilber',
        'year_published': '2018'
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


@app.route('/', methods=['GET'])
def index():
    return '''
    <h1>Reading Club</h1>
    <a href='/api/v1.0/books'>Books</a>
    '''


@app.route('/api/v1.0/books', methods=['GET'])
def get_books():
    return jsonify(books)


@app.route('/api/v1.0/books/<int:book_id>', methods=['GET'])
def get_book(book_id):
    book = [book for book in books if book['id'] == book_id]
    if len(book) == 0:
        abort(404)
    return jsonify({'book': book[0]})


if __name__ == '__main__':
    app.run(host="0.0.0.0", debug=True)
