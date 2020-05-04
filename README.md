# golang with hex architech

This project go lang I try to make it flexible more more by seperate by Hexagonal architecture. It so useful when we used this. the domain logic can turn to Rest API, Command script or anything that we want to used this domain logic. Because domain logic have interface who want to integrate should be follow the interface for every layer.

reference: https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749

### Rest API 
we used gin server and call domain logic interface
```
implem/gin.server
```

### Command script 
cmd used domain logic
```
cmd
```

### Astilectron
we would like to tranform to desktop application also