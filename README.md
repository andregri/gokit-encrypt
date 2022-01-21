# Go Microservice with go-kit

```
curl -X POST http://localhost:8080/encrypt \
    -d '{
        "key":"111023043350789514532147",
        "text":"I am a message"
        }'

{"message":"8/+JCfT7+gbIjzQtmCo=","error":""}
```

```
curl -X POST http://localhost:8080/decrypt \
    -d '{
        "key":"111023043350789514532147",
        "message":"8/+JCfT7+gbIjzQtmCo="
        }'

{"text":"I am a message","error":""}
```

With logging the microservice prints:
```
method=encrypt key=111023043350789514532147 text="I am a message" output="8/+JCfT7+gbIjzQtmCo=" err=null took=8.43µs
method=decrypt key=111023043350789514532147 message="8/+JCfT7+gbIjzQtmCo=" output="I am a message" err=null took=11.786µs
```