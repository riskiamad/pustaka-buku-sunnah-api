<h1>pustaka-buku-sunnah-api v1</h1>

<h2>This is my first REST API using go gin framework, following the tutorial from YT video by Agung Setiawan https://www.youtube.com/watch?v=GjI0GSvmcSU</h2>

<h2>How to use the app ?</h2>

<h3>if you want to try the app you can use my host as endpoint, replace host with https://my1stapi.herokuapp.com/</h3>

<h3>make sure you are add v1 after host address</h3>

<details>
           <summary>1. Find all data()</summary>
           <p>GET host/v1/books
  
  this endpoint will return all data in json</p>
</details>

<details>
           <summary>2. Find specified data by id</summary>
           <p>GET host/v1/books/:id
   
   ex: host/v1/books/5
   the above example will return specified book with id 5</p>
</details>

<details>
           <summary>3. Add new book to database</summary>
           <p>POST host/v1/books
    
    make sure you are add body parameter like these :
    
    {
      "title":"string data",
      "Price":int,
      "Description":"string data",
      "Rating":int,
      "Discount":int
    }
    
    please beware with the data type of each field, becuse I am add validation to them.</p>
</details>

<details>
           <summary>4. Update existing book in database</summary>
           <p>PUT host/v1/books/:id
    
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
    
    please beware with the data type of each field, becuse I am add validation to them.</p>
</details>

<details>
           <summary>5. Delete any record in existing data</summary>
           <p>DELETE host/v1/books/:id
    
    ex: host/v1/books/5
    the above example will delete the record with id 5</p>
</details>

if you have some inquiries regarding this repository, don't hesitate to send me email at  [riskiamad@gmail.com](mailto:riskiamad@gmail.com "send mail to me")
