curl -XGET -H "Content-type: application/json" 'http://localhost:9001/api/v1/items'

curl -XPOST -H "Content-type: application/json" -d '{"email":"test@test.com","password":"123123"}' 'http://localhost:9001/api/v1/signup'

curl -XPOST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDExODk2NzR9.8GF1krzctzQebAEIF3rkCd0vGk2lVdCT36M-d1QzoEw' -H "Content-type: application/json" -d '{"name":"coffee","price":10,"quantity":10}' 'http://localhost:9001/api/v1/items'

curl -XPOST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDExODk2NzR9.8GF1krzctzQebAEIF3rkCd0vGk2lVdCT36M-d1QzoEw' -H "Content-type: application/json" -d '{"name":"milk","price":299,"quantity":100}' 'http://localhost:9001/api/v1/items'