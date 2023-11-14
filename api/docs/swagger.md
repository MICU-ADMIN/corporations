```yaml

basePath: /api
definitions:
  controllers.LoginPayload:
    properties:
      email:
        type: string
      password:
        type: string
    required:
    - email
    - password
    type: object
  controllers.UserDetails:
    properties:
      email:
        type: string
      name:
        type: string
      password:
        type: string
    required:
    - email
    - name
    - password
    type: object
host: localhost:8080
info:
  contact:
    email: support@swagger.io
    name: API Support
    url: http://demo.com/support
  description: Create  Go REST API with JWT Authentication in Gin Framework
  license:
    name: MIT
    url: https://opensource.org/licenses/MIT
  termsOfService: demo.com
  title: Swagger JWT API
  version: "1.0"
paths:
  /protected/profile:
    get:
      operationId: GetUserByToken
      parameters:
      - description: Authorization header using the Bearer scheme
        in: header
        name: Authorization
        required: true
        type: string
      produces:
      - application/json
      responses:
        "200":
          description: Success
          schema:
            type: string
        "400":
          description: Error
          schema:
            type: string
      summary: Get User By Token
      tags:
      - User
  /public/login:
    post:
      consumes:
      - application/json
      description: Login
      operationId: LoginUser
      parameters:
      - description: Login
        in: body
        name: EnterDetails
        required: true
        schema:
          $ref: '#/definitions/controllers.LoginPayload'
      responses:
        "200":
          description: Success
          schema:
            type: string
        "400":
          description: Error
          schema:
            type: string
      summary: Login User
      tags:
      - User
  /public/signup:
    post:
      consumes:
      - application/json
      description: Signin
      operationId: SignupUser
      parameters:
      - description: Signin
        in: body
        name: EnterDetails
        required: true
        schema:
          $ref: '#/definitions/controllers.UserDetails'
      responses:
        "200":
          description: Success
          schema:
            type: string
        "400":
          description: Error
          schema:
            type: string
      summary: Signup User
      tags:
      - User
schemes:
- http
- https
securityDefinitions:
  ApiKeyAuth:
    in: header
    name: Authorization
    type: apiKey
  BasicAuth:
    type: basic
swagger: "2.0"


```

```mermaid

Here is a Mermaid Markdown overview for the provided Go file:
### API Description

The API is a Go RESTful API with JWT authentication implemented using the Gin Framework. It has the following endpoints:

### Login Endpoint

The `/public/login` endpoint allows users to log in to the API using their email and password. The request body must contain an object with the email and password fields. If the credentials are valid, the API will return a `200` status code and a JSON response with the user's details. If the credentials are invalid, the API will return a `400` status code and a JSON response with an error message.
### Signup Endpoint

The `/public/signup` endpoint allows users to sign up for an account on the API. The request body must contain an object with the email and password fields. If the signup is successful, the API will return a `200` status code and a JSON response with the user's details. If the signup is unsuccessful, the API will return a `400` status code and a JSON response with an error message.
###

```
