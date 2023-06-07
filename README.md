Usage:
- `docker-compose up -d`

API **POST**: `http://localhost:3000/fam/users` (Create/Register user)
- Body:
  ```json
  {
    "username": "quangtue",
    "password": "quangtue"
  }
  ```
- Response:
  ```json
  {
    "code": 1,
    "data": {
      "_id": 1
    },
    "message": "Success",
    "success": true
  }
  ```

API **GET**: `http://localhost:3000/fam/users/login` (Login user)
- Body:
  ```json
  {
    "username": "quangtue",
    "password": "quangtue"
  }
  ```
- Response:
  ```json
  {
    "code": 1,
    "data": {
      "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODYyMDg0MzQsImlkIjoxLCJ1c2VybmFtZSI6InF1YW5ndHVlIn0.gLhDlC1Jxlu6YitMFfGxiei6PYdoyAmFiKt8w3JD3ow"
    },
    "message": "Success",
    "success": true
  }
  ```

API **POST**: `http://localhost:3000/fam/employees` (Create Employee) - Need AccessToken
- Body:
  ```json
  {
    "name": "Tue"
  }
  ```
- Response:
  ```json
  {
    "code": 1,
    "data": {
      "_id": 8
    },
    "message": "Success",
    "success": true
  }
  ```

API **PUT**: `http://localhost:3000/fam/employees/5` (Update Employee/Add supervisor) - Need AccessToken
- Body:
  ```json
  {
    "supervisor_id": 3
  }
  ```
- Response:
  ```json
  {
    "code": 1,
    "data": {
      "_id": 5
    },
    "message": "Success",
    "success": true
  }
  ```

API **GET**: `http://localhost:3000/fam/employees?name=Nick` (Get Employee Info) - Need AccessToken
- Response:
  ```json
  {
    "code": 1,
    "data": {
      "_id": 3,
      "name": "Nick",
      "supervisor_id": 1,
      "supervisor_name": "Jonas",
      "subordinates": [
        {
          "_id": 2,
          "name": "Sophie",
          "supervisor_id": 3,
          "subordinates": null
        },
        {
          "_id": 4,
          "name": "Pete",
          "supervisor_id": 3,
          "subordinates": null
        },
        {
          "_id": 5,
          "name": "Barbara",
          "supervisor_id": 3,
          "subordinates": null
        }
      ]
    },
    "message": "Success",
    "success": true
  }
  ```
