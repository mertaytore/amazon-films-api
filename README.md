# amazon-films-api
Necessary information of any movie in Amazon, in JSON format, written in Go!

### How to set up? ###
* **Summary of set up**  
`go get ./..`  
change directory to project's folder  
`go build -o "scraper"`  
`./scraper`  

Now you have a working scraper where you can make requests from your favorite web browser.  

* **Dependencies**  
[goquery](https://github.com/PuerkitoBio/goquery): is used to fetch the needed parts of the movie webpage, in an easy manner.  
[mux](https://github.com/gorilla/mux): mux is used to create a router and its handlers. Easy to pass arguments via the route itself.  

### How to use? ###

Depending on the Amazon's region of the movie ID - com, co.uk and de works fine for this implementation - specify the region and the movie ID to the link in your browser. Since hardcoding amazon.de for the base url region wouldn't give correct results for all, specifying the region in the link seemed to provide easier usage.
  
**Sample execution: ** http://localhost:8080/movie/amazon/de/B00K19SD8Q   
**Result: **   
{  
&nbsp;&nbsp;&nbsp;&nbsp;"title":"Um Jeden Preis",  
&nbsp;&nbsp;&nbsp;&nbsp;"release_year":2013,  
&nbsp;&nbsp;&nbsp;&nbsp;"actors":[    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"Dennis Quaid",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;" Zac Efron",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;" Kim Dickens"  
&nbsp;&nbsp;&nbsp;&nbsp;],  
&nbsp;&nbsp;&nbsp;&nbsp;"posters":"https://images-eu.ssl-images-amazon.com/images/I/51VELYHd4TL._SX200_QL80_.jpg",  
&nbsp;&nbsp;&nbsp;&nbsp;"similar_ids":[    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00HDZMP94",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00JM0JXYI",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00L9KET84",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00O4G7QQ2",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00HUZXE8S",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00ILR6K9Y",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00I4DWI6O",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00ILNAY5O",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00HXL6E8G",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00ILNV05W",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00ILIGIU4",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00JKGLKFY",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00G0NPW1I",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00IK6HLUI",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00IYSFS6Q",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00N1ECZ2S",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B01KUYZ6TU",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00JLCZN8M",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B06XP6H33J",  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"B00JKEXKCC"  
&nbsp;&nbsp;&nbsp;&nbsp;]  
}  
  
**Sample execution 2: ** http://localhost:8080/movie/amazon/co.uk/B00HDMDADW

**Sample execution 3: ** http://localhost:8080/movie/amazon/com/B0018OFN1S
