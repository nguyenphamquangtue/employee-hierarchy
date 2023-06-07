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


1. API **POST**: `http://localhost:3000/fam/users` (Create/Register user)
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

2. API **GET**: `http://localhost:3000/fam/users/login` (Login user)
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

3. API **POST**: `http://localhost:3000/fam/employees` (Create Employee) - Need AccessToken
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
4. API **PUT**: `http://localhost:3000/fam/employees/1` (Update Employee/Add supervisor) - Need AccessToken
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

5. API **GET**: `http://localhost:3000/fam/employees?name=Nick` (Get Employee Info) - Need AccessToken
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

Các test case quan trọng cần test:
- API 3:
  - Tạo mới employee nhưng trùng tên với employee trước đó => throw message : "Employee Existed"
- API 4:
  - http://localhost:3000/fam/employees/<:id> (id là employee_id)
  - Nếu supervisor_id trùng với employee_id => throw message:  "Supervisor Cannot be the same as the employee"
  - Subordinate (Cấp dưới) của employee không thể là supervisor(Cấp trên) của employee đó
    - Ví dụ: 
        - Nick là cấp dưới của Jonas, Pete là supervisor của Jonas => Nick không thể là cấp trên của Pete
        => throw message: "Cycle Detected In Supervisor Hierarchy"