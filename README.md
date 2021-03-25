API is running on port 8080. Swagger is running on port 8001.

## Usage 
Build and run:
```
make
```

Use:
```
curl --header "Content-Type: application/json" \                                                                                               
  --request POST \
  --data '{"msisdn": "7307406945"}' \     
  http://localhost:8080/msisdn
```
