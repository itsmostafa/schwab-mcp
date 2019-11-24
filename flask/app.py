from flask import Flask, jsonify, request


app = Flask(__name__)

books = [
    {"name": "Green Eggs and Ham", "price": 7.99, "isbn": 9780394848,},
    {"name": "The Cat In The Hat", "price": 6.99, "isbn": 9784823748,},
]

# GET /books
@app.route('/')
def get_books():
    return jsonify({"books": books})

# POST /books
@app.route('/books', methods=['POST'])
def add_book():
    return jsonify(request.get_json())

@app.route('/books/<int:isbn>')
def get_book_by_isbn(isbn):
    return_value = {}
    for book in books:
        if book["isbn"] == isbn:
            return_value = {
                "name": book["name"],
                "price": book["price"],
            }
    return jsonify(return_value)

app.run(port=5000)
