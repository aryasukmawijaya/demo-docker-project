# demo-docker-project
Demo docker with go and mysql

How to run:
```
docker-compose up
```

Endpoints:
- GET /user
  return all user

- POST /user
  insert new user
  body:
  ```json
  {
    "name": "Alice",
    "email": "alice@gmail.com",
    "hobby": "Coding
  }
  ```
