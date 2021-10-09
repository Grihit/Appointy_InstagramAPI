# Appointy_InstagramAPI
An Instagram API made using Go and MongoDB

## Tasks that can be performed by this API
1. Create an User (POST request) <br />
URL used: http://localhost:27017/users <br />
![Create A User 1](https://github.com/Grihit/Appointy_InstagramAPI/blob/master/Images/Create%20User%201.PNG) 
The data is entered in body in JSON format and _id is automatically created <br />
![Create A User 2](https://github.com/Grihit/Appointy_InstagramAPI/blob/master/Images/Create%20User%202.PNG)
The user is successfully inserted to the database <br />

2. Get user using id (GET request) <br />
URL used: http://localhost:27017/users/6161701c6c8378c960d53965 <br />
![Get User using id](https://github.com/Grihit/Appointy_InstagramAPI/blob/master/Images/Get%20User%20using%20id.PNG)
The User details of the user id in the url is displayed <br />

3. Create a post (POST request) <br />
URL used: http://localhost:27017/posts <br />
![Create a post 1](https://github.com/Grihit/Appointy_InstagramAPI/blob/master/Images/Create%20a%20post%201.PNG)
The data is entered in body in JSON format and _id is automatically created  <br />
![Create a post 2](https://github.com/Grihit/Appointy_InstagramAPI/blob/master/Images/Create%20a%20post%202.PNG)
The post is successfully inserted to the database <br />

4. Get post using post id (GET request) <br />
URL used: http://localhost:27017/posts/61617ca36c8378aef4d0a3a9 <br />
![Get post using post id](https://github.com/Grihit/Appointy_InstagramAPI/blob/master/Images/Get%20post%20using%20post%20id.PNG)
The post details of the post id in the url is displayed <br />

5. Get all posts of a specific user (GET request) <br />
URL used: http://localhost:27017/posts/users/61617076dc25d3de5f044084 <br />
![Get all posts of a specific user](https://github.com/Grihit/Appointy_InstagramAPI/blob/master/Images/Get%20posts%20of%20a%20specific%20user%201.PNG)
![Get all posts of a specific user 2](https://github.com/Grihit/Appointy_InstagramAPI/blob/master/Images/Get%20posts%20of%20a%20specific%20user%202.PNG)
All the posts of the user with the user id in the url are displayed <br />

6. Get all users (GET request) <br />
URL used: http://localhost:27017/users <br />
![Get all users](https://github.com/Grihit/Appointy_InstagramAPI/blob/master/Images/Get%20all%20users.PNG)
Details of all the users in the database <br />

7. Get all posts (GET request) <br />
URL used: http://localhost:27017/posts <br />
![Get all posts](https://github.com/Grihit/Appointy_InstagramAPI/blob/master/Images/Get%20all%20posts.PNG)
Details of all the posts in the database <br />
<br />
Note: The (upper)body data has no significance when the request is GET

## How to run
1. Run the command mongod in your cmd terminal at the directory where mongodb is installed.
2. Make sure a database is created with the name Instagram.
3. Use mongodb compass for easy creation of database and also for viewing the data in the database.
4. Use the URLs in a service like postman for POST and GET requests.
