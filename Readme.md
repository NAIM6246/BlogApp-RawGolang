# WorkSpacify-BlogApp

After downloading the project run: go mod tidy 
Need a folder named "uploads" in the root directory of the project to upload file

## create user
### Method: POST
### Endpoint: localhost:8005/api/users
### Req body: 
{
    "name" : "naim",
    "email" : "naim@gmail.com",
    "password" : "pass12345"
}

## get all users
### Method : GET
### Endpoint: localhost:8005/api/users/

## Create Post
### Method: POST
### Endpoint: localhost:8005/api/posts/
### Req body: form 
#### Note: representing the field name like json
{
    "file": file.png,
    "description": "demo des",
    "like_count": 0,
    "unlike_count": 0,
    "comment_count": 0,
    "author_id": 1
}

## Get posts (paginated)
### Method: Get
### Endpoint: localhost:8005/api/posts/
### Req body: 
{
    "last_id" : 20,
    "limit" : 5
}


## Create Reaction 
### Method: POST
### Endpoint: localhost:8005/api/reactions/
### Req body : 
ex-1: for unlike
{
    "post_id": 2,
    "user_id": 6,
    "is_unlike": true
}

ex-2: for like 
{
    "post_id": 2,
    "user_id": 6,
    "is_like": true
}


## Get reacted users of post
### Method: Get
### Endpoint : localhost:8005/api/reactions/reacted-users
### Req body:
ex-1: liked users 
{
    "post_id": 1,
    "limit": 2,
    "last_id": 0,
    "liked": true
}

ex-2: unliked users
{
    "post_id": 1,
    "limit": 2,
    "last_id": 0,
    "liked": false
}