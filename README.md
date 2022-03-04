# Library
Library is a simple application which can be used to add and view the books.

**APIs exposed:**

**/books - GET**

Output:

`[
    {
        "id": 1,
        "name": "book1",
        "price": 0,
        "edition": ""
    }
    .
    .
    .
]`

 - Status OK: list of books present
 - Status Internal Server Error : Server error
 
 **/books/id - GET**

Output:

`{
        "id": 1,
        "name": "book1",
        "price": 0,
        "edition": ""
    }
 `

 - Status OK: list of books present
 - Status Internal Server Error : Server error
 
**/books - POST**

Payload:

`{
        "id": 1,
        "name": "book1",
        "price": 0,
        "edition": ""
    }`
    
    
Output:

 - Status OK: book with id from database
 - Status Conflict: record with same id already present
- Status Internal Server Error : Server error

**Steps:**
 - pull library .
 - run `sudo docker run --name mongodb -d -p 27017:27017 mongo`
 - run - `go run main.go`
 - API server running on localhost:8000
 
