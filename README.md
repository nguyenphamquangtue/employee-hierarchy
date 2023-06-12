Usage:
- `docker-compose up -d`

- Khi bạn start app, tôi đã seed sẵn các data:

| ID | Name    | Supervisor ID |
|----|---------|---------------|
| 1  | Jonas   | null          |
| 2  | Sophie  | 3             |
| 3  | Nick    | 1             |
| 4  | Pete    | 3             |
| 5  | Barbara | 3             |


1. API **POST**: `http://localhost:3000/fram/users` (Create/Register user)
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

2. API **GET**: `http://localhost:3000/fram/users/login` (Login user)
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

3. API **POST**: `http://localhost:3000/fram/employees` (Create Employee) - Need AccessToken
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
4. API **PUT**: `http://localhost:3000/fram/employees/1` (Update Employee/Add supervisor) - Need AccessToken
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

5. API **GET**: `http://localhost:3000/fram/employees?name=Nick` (Get Employee Info) - Need AccessToken
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

Important test cases to consider:

- API 3:
    Create a new employee with the same name as an existing employee => throw message: "Employee Existed"
- API 4:
    + http://localhost:3000/fram/employees/<:id> (where id is the employee_id)
      If supervisor_id is the same as employee_id => throw message: "Supervisor Cannot be the same as the employee"
    + A subordinate (lower level) of an employee cannot be the supervisor (higher level) of that employee.
      Example:
      Nick is a subordinate of Jonas, Pete is the supervisor of Jonas => Nick cannot be the superior of Pete
      => throw message: "Cycle Detected In Supervisor Hierarchy"