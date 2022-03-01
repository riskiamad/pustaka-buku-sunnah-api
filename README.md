pustaka-buku-sunnah-api v1

This is my first REST API using go gin framework, following the tutorial from YT video by Agung Setiawan https://www.youtube.com/watch?v=GjI0GSvmcSU

How to use the app ?

if you want to try the app you can use my host as endpoint, replace host with https://my1stapi.herokuapp.com/

make sure you are add v1 after host address
1. Find all data()
  GET host/v1/books
  
  this endpoint will return all data in json

2. Find specified data by id
   GET host/v1/books/:id
   
   ex: host/v1/books/5
   the above example will return specified book with id 5
 
3. Add new book to database
    POST host/v1/books
    
    make sure you are add body parameter like these :
    
    {
      "title":"string data",
      "Price":int,
      "Description":"string data",
      "Rating":int,
      "Discount":int
    }
    
    please beware with the data type of each field, becuse I am add validation to them.

4. Update existing book in database
    PUT host/v1/books/:id
    
    ex: host/v1.books/5
    the above example will modify the data with id 5
    
    make sure you are add body parameter like these :
    
    {
      "title":"string data",
      "Price":int,
      "Description":"string data",
      "Rating":int,
      "Discount":int
    }
    
    please beware with the data type of each field, becuse I am add validation to them.

5. Delete some record in existing data
    DELETE host/v1/books/:id
    
    ex: host/v1/books/5
    the above example will delete the record with id 5

if you have some inquiries regarding this repository, don't hesitate to send me email at riskiamad@gmail.com
