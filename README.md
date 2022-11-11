# distributed-systems_Golang
Just a crud in Golang using Rest and gRPC to communicate with database and process with Architecture based on Microservices.

## REST API

![API REST Golang](images/golangImg.png)


[Application Link](http://localhost:5000) - Localhost

1. Create 

    Route - **/createUser** \
    HTTP Method - POST

    **body**

    ~~~json
    {
        "name": "exemple",
        "password": "passwordEncryptExample",
        "email": "email@exemple.com",
        "address": "addressExemple"
    }
    ~~~

2. Read 

    Route - **/getUserId/{id}**\
    HTTP Method - GET