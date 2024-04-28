
## API Usage

This API is designed for library systems. There are User,Book and Borrow endpoints. 

Fiber framework, Goose, POSTGRESQL DB, SQLC and Docker was used by developing this API.

### Contents
* [Installation & Deployment](#installation--deployment)
* [Log into accounts](#log-into-accounts)
* [Users](#users)
	* [Get Users](#get-users)
	* [Get User](#get-user)
	* [Create User](#create-user)
    * [Delete User](#delete-user)
    * [Update User](#update-user)
* [Books](books)
	* [Get Books](#get-books)
    * [Get Book](#get-book)
    * [Create Books](#create-book)
    * [Delete Books](#delete-book)
    * [Update Books](#update-book)
* [Borrows](#borrows)
    * [Get Borrows History](#get-borrows-history)
    * [Return Borrowed Book](#return-borrowed-book)
    * [Borrow Book](#borrow-book)
    * [Get Return History](#get-return-history)
 
* ### Installation & Deployment

    1. Clone the repo
       
            git clone https://github.com/melikesenacakir/GoLibraryAPI.git
    3. Go to this file
       
            cd GoLibraryAPI
    4. Run the system in your terminal
       
           docker-compose up

### **Log Into Accounts**

```http
  POST /api/v1/login
```
**Required:**

| Parameter| Type     | Explanation                |
| :-------- | :------- | :------------------------- |
| `role` | `string` | Enter role(admin/user) |
| `username` | `string` |  Enter username |
| `password` | `string` | Enter password |


#### Response STATUS 200:OK

        {
            "role":"admin",
            "username":"admin121",
            "password":"123456"
        }

## **Users**

### **Get Users**

```http
  GET /api/v1/user
```

**Required:**

     Logged into an existing Admin account


#### Response STATUS 200:OK

        [
            {
                "role": "admin",
                "username": "admin121"
            }
        ]

### **Get User**

```http
  GET /api/v1/user/id
```

**Required:**

    - Logged into an existing Admin account

| Parameter| Type     | Explanation                |
| :-------- | :------- | :------------------------- |
| `id` | `int` |  Enter id |



#### Response STATUS 200:OK

        [
            {
                "role": "admin",
                "username": "admin121"
            }
        ]

### **Create User**

```http
  POST /api/v1/user/create
```

**Required:**

    - Logged into an existing Admin account
| Parameter| Type     | Explanation                |
| :-------- | :------- | :------------------------- |
| `role` | `string` |  Enter role(admin/user) |
| `username` | `string` |  Enter username |
| `password` | `string` |  Enter password |



#### Response STATUS 200:OK

        {
           "message": "user successfully created"
        }

### **Delete User**

```http
  DELETE /api/v1/user/delete/id
```

**Required:**

    - Logged into an existing Admin account
| Parameter| Type     | Explanation                |
| :-------- | :------- | :------------------------- |
| `id` | `int` |  Enter id to url |



#### Response STATUS 200:OK

        {
           "message": "user successfully deleted"
        }

### **Update User**

```http
  PUT /api/v1/user/update/id
```

**Required:**

    - Logged into an existing Admin account
| Parameter| Type     | Explanation                |
| :-------- | :------- | :------------------------- |
| `id` | `int` |  Enter id to url |
| `role` | `string` | Enter role(admin/user) |
| `username` | `string` |  Enter username |
| `password` | `string` | Enter password |



#### Response STATUS 200:OK

        {
           "message": "user successfully updated"
        }

## **Books**

### **Get Books**

```http
  GET /api/v1/book
```


#### Response STATUS 200:OK

        [
            {
                "id": 1,
                "name": "tanios kayası",
                "author": "amin maalouf",
                "stock": 5
            }
        ]

### **Get Book**

```http
  GET /api/v1/book/id
```
**Required:**

| Parameter| Type     | Explanation                |
| :-------- | :------- | :------------------------- |
| `id` | `int` |  Enter id to url |

#### Response STATUS 200:OK

        [
            {
                "id": 1,
                "name": "tanios kayası",
                "author": "amin maalouf",
                "stock": 5
            }
        ]

### **Create Book**

```http
  POST /api/v1/book/create
```
**Required:**

    - Logged into an existing Admin or User account

| Parameter| Type     | Explanation                |
| :-------- | :------- | :------------------------- |
| `id` | `int` |  Enter id to url |

#### Response STATUS 200:OK

        [
            {
                "id": 1,
                "name": "tanios kayası",
                "author": "amin maalouf",
                "stock": 5
            }
        ]

### **Delete Book**

```http
  DELETE /api/v1/book/delete/id
```
**Required:**

    - Logged into an existing Admin or User account

| Parameter| Type     | Explanation                |
| :-------- | :------- | :------------------------- |
| `id` | `int` |  Enter id to url |

#### Response STATUS 200:OK

        {
           "message": "book successfully deleted"
        }

### **Update Book**

```http
  DELETE /api/v1/book/update/id
```
**Required:**

    - Logged into an existing Admin or User account

| Parameter| Type     | Explanation                |
| :-------- | :------- | :------------------------- |
| `id` | `int` |  Enter id to url |
| `name` | `string` |  Enter name |
| `author` | `string` |  Enter author |
| `stock` | `int` |  Enter stock |

#### Response STATUS 200:OK

        {
           "message": "book successfully updated"
        }

## **Borrows**

### **Get Borrows History**

```http
  GET /api/v1/borrows/history
```
**Required:**

    - Logged into an existing Admin or User account


#### Response STATUS 200:OK

        [
            {
                "borrow_id": 1,
                "borrowdate": "2024-04-28T23:24:35.212264Z",
                "returndate": "2024-05-08T23:24:35.212265Z",
                "status": "borrowed",
                "username": "admin121",
                "book_name": "tanios kayası"
            },
            {
                "borrow_id": 2,
                "borrowdate": "2024-04-28T23:24:35.212264Z",
                "returndate": "2024-05-08T23:24:35.212265Z",
                "status": "returned",
                "username": "user",
                "book_name": "tanios kayası"
            }
        ]

### **Return Borrowed Book**

```http
  PUT /api/v1/borrows/return/id
```
**Required:**

    - Logged into an existing Admin or User account

| Parameter| Type     | Explanation                |
| :-------- | :------- | :------------------------- |
| `id` | `int` |  Enter id to url |

#### Response STATUS 200:OK

        {
          "message": "book returned"
        }

### **Borrow Book**

```http
  POST /api/v1/borrows/new
```
**Required:**

    - Logged into an existing Admin or User account

| Parameter| Type     | Explanation                |
| :-------- | :------- | :------------------------- |
| `user_id` | `int` |  Enter user id |
| `book_id` | `int` |  Enter book id |

#### Response STATUS 200:OK

        {
           "message": "books successfully borrowed. Please return on time!"
        }

### **Get Return History**

```http
  GET /api/v1/borrows/oldborrows
```
**Required:**

    - Logged into an existing Admin or User account


#### Response STATUS 200:OK

        [
            {
                "borrow_id": 1,
                "borrowdate": "2024-04-28T23:24:35.212264Z",
                "returndate": "2024-05-08T23:24:35.212265Z",
                "status": "returned",
                "username": "admin121",
                "book_name": "tanios kayası"
            }
        ]
