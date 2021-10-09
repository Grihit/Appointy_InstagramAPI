# Appointy_InstagramAPI
An Instagram API made using Go and MongoDB

## Tasks that can be performed by this API
1. Create an User (POST request)
URL used: http://localhost:27017/users

2. Get user using id (GET request)
URL used: http://localhost:27017/users/6161701c6c8378c960d53965

3. Create a post (POST request)
URL used: http://localhost:27017/posts

4. Get post using post id (GET request)
URL used: http://localhost:27017/posts/61617ca36c8378aef4d0a3a9

5. Get all posts of a specific user (GET request)
URL used: http://localhost:27017/posts/users/61617076dc25d3de5f044084

6. Get all users (GET request)
URL used: http://localhost:27017/users

7. Get all posts (GET request)
URL used: http://localhost:27017/posts

## How to run
1. Run the command mongod in your cmd terminal at the directory where mongodb is installed
2. Make sure a database is created with the name Instagram and collections Users and Posts with their respective attributes.
3. Use mongodb compass for easy creation of database and also for viewing the data in the database
4. Use the URLs in a service like postman for POST and GET requests
