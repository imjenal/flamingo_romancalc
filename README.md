# Deeptrace


A web service implementing FizzBuzz with a Pink Flamingo and Roman Calculator 

# How To Run Locally

docker build -t deeptrace .

docker run -d 8081:8081 deeptrace

# How To Test
###FizzBuzz with a Pink Flamingo
```
curl -X GET 'http://localhost:8081/pinkflamingo?from=0&to=10'

Expected Response: 
["Pink Flamingo","Flamingo","Flamingo","Flamingo","4","Flamingo","Fizz","7","Flamingo","Fizz"]
```

###Roman Calculator
```
curl -X POST 'http://localhost:8081/romancalc' -d "{  \"expr\": \"II * VIII / II\"}"

Expected Response: VIII
```