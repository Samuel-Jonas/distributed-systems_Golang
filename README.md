# distributed-systems_Golang
Just a crud in Golang using Rest and gRPC to communicate with database and process with Architecture based on Microservices.

## REST API

![API REST Golang](images/golangImg.png)



#### Setup application 

    1. Clone this repos
    2. Ensure Docker is running correctly on your machine
    3. Install Docker Compose if you don´t have it
    4. In the aplication´s local path run:
    


~~~shell
docker-compose up
~~~


Application listening on - <http://localhost:5000>

1. Create 

    Route - **/createUser** \
    HTTP Method - POST

    **Request body**

    ~~~json
    {
        "name": "exemple",
        "password": "passwordEncryptExample",
        "email": "email@exemple.com",
        "address": "addressExemple"
    }
    ~~~

2. Read 

    Route - **/getUserId/{id}** \
    HTTP Method - GET \
    Path Variable - User identifier passed in URL

    **Request body**

    ~~~javascript
    empty
    ~~~

    Route - **/getUsers** \
    HTTP Method - GET 
    
    **Request body**

    ~~~javascript
    empty
    ~~~

3. Update

    Route - **/updateUserId/{id}** \
    HTTP Method - PUT \
    Path Variable - User identifier passed in URL

    **Request body**

    ~~~json
    {
        "name": "exemple",
        "password": "passwordEncryptExample",
        "email": "email@exemple.com",
        "address": "addressExemple"
    }
    ~~~

4. Delete

    Route - **/deleteUserId/{id}** \
    HTTP Method - DELETE \
    Path Variable - User identifier passed in URL

    **Request body**

    ~~~javascript
    empty
    ~~~