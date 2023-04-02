# courses

## STEP By STEP
- Create DB `courses` OR run `docker-compose up -d` in CLI
- RUN `migrate -database 'mysql://root:mauFJcuf5dhRMQrjj@tcp(localhost:3306)/courses?parseTime=true' -path ./db/migrations up`
- RUN `go run .`

## cURL
- Login Admin
```
curl --location --request POST 'http://localhost:8880/api/v1/admin/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"admin",
    "password":"admin"
}'
```

- Create Course
```
curl --location --request POST 'http://localhost:8880/api/v1/admin/course' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4iLCJ0eXBlIjoiYWRtaW4iLCJjcmVhdGVkX2F0IjoiMjAyMy0wNC0wMVQxMjo0NTowMFoifSwiZXhwIjoxNzEyMDM2NzQ1LCJpYXQiOjE2ODA0MTQzNDV9.7JvphkdJStXJfKY9YZYwh9ssvbdg2WQfYir-_HXPcYw' \
--form 'name="GO 5"' \
--form 'category_id="1"' \
--form 'price="130000"' \
--form 'image=@"/Users/ihsanmaulana/Desktop/Screen Shot 2023-02-21 at 13.48.24.png"'
```

- Update Course
```
curl --location --request PUT 'http://localhost:8880/api/v1/admin/course/7' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"PHP 2",
    "category_id":1,
    "price":120000
}'
```

- Get List Courses
```
curl --location --request GET 'http://localhost:8880/api/v1/course?page=1&limit=10&name=GO 4&price=130000&sort_by=price DESC' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4iLCJ0eXBlIjoiYWRtaW4iLCJjcmVhdGVkX2F0IjoiMjAyMy0wNC0wMVQxMjo0NTowMFoifSwiZXhwIjoxNzEyMDM2NzQ1LCJpYXQiOjE2ODA0MTQzNDV9.7JvphkdJStXJfKY9YZYwh9ssvbdg2WQfYir-_HXPcYw'
```

- Get By ID
```
curl --location --request GET 'http://localhost:8880/api/v1/admin/course/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4iLCJ0eXBlIjoiYWRtaW4iLCJjcmVhdGVkX2F0IjoiMjAyMy0wNC0wMlQwODowMTo0NFoifSwiZXhwIjoxNzEyMDQ2MTE0LCJpYXQiOjE2ODA0MjM3MTR9.r8CEnPY7cWYc63qpwIIVOaOygYrqulxAfq0tr7gNEHg'
```

- DELETE User
```
curl --location --request DELETE 'http://localhost:8880/api/v1/admin/user/4' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImlkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwidHlwZSI6ImFkbWluIiwiY3JlYXRlZF9hdCI6IjIwMjMtMDQtMDJUMDg6MDE6NDRaIn0sImV4cCI6MTcxMjA1OTI5MiwiaWF0IjoxNjgwNDM2ODkyfQ.xzdiHFuK8Uh-o2WUeLc9q-7sDgpjm_Ab0_XEiVsF_3s'
```

- Statistic
```
curl --location --request GET 'http://localhost:8880/api/v1/admin/statistic' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImlkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwidHlwZSI6ImFkbWluIiwiY3JlYXRlZF9hdCI6IjIwMjMtMDQtMDJUMDg6MDE6NDRaIn0sImV4cCI6MTcxMjA1OTI5MiwiaWF0IjoxNjgwNDM2ODkyfQ.xzdiHFuK8Uh-o2WUeLc9q-7sDgpjm_Ab0_XEiVsF_3s'
```


## Curl USER
- Register User
```
curl --location --request POST 'http://localhost:8880/api/v1/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"usertiga",
    "password":"usertiga"
}'
```

- Get Course OR Search Course
```
curl --location --request GET 'http://localhost:8880/api/v1/user/course?page=1&limit=10&name=GO 4&price=130000&sort_by=price DESC'
```

- Get Course Detail
```
curl --location --request GET 'http://localhost:8880/api/v1/user/course/1'
```

- Get Category Course
```
curl --location --request GET 'http://localhost:8880/api/v1/user/category/course'
```

