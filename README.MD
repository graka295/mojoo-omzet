set config application.develop.yaml
<br>

api login
http://localhost:1323/api/login
body
{
    "username": "admin1",
    "password": "admin1"
}

api omzet merchant
http://localhost:1323/api/admin/merchant
body
{
    "id": 1,
    "year": 2021,
    "month": 11,
    "page": 1
}

api omzet outlets
http://localhost:1323/api/admin/outlets
body
{
    "id": 1,
    "year": 2021,
    "month": 11,
    "page": 1
}
