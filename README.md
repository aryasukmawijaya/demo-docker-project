# demo-docker-project
Demo docker with go and mysql

How to run:
```
docker-compose up
```

Endpoints:
- GET /user <br>
  return all user

- POST /user <br>
  insert new user <br>
  body:
  ```json
  {
    "name": "Alice",
    "email": "alice@gmail.com",
    "hobby": "Coding"
  }
  ```
