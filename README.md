# distributed-systems_Golang
Just a crud in Golang using Rest and gRPC to communicate with database and process with Architecture based on Microservices.

## REST API CRUD

[Application Link](http://localhost:5000) - Localhost

![API REST Golang](https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.mitrais.com%2Fnews-updates%2Fhow-to-dockerize-a-restful-api-with-golang-and-postgres%2F&psig=AOvVaw1jupGswdIEUOTn-ee3qWIA&ust=1668268624498000&source=images&cd=vfe&ved=0CBAQjRxqFwoTCNjkgoi_pvsCFQAAAAAdAAAAABAE)

1. Create 

    Route - **/createUser**
    HTTP Method - POST

    **BODY**

    ~~~json
    {
        "name": "exemple",
        "password": "passwordEncryptExample",
        "email": "email@exemple.com",
        "address": "addressExemple"
    }
    ~~~

2. Read 

    Route - **/getUserId/{id}**
    HTTP Method - GET