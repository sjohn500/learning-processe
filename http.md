## what is HTTP?

the full meaning of HTTP is HYPERTEXT TRANSFER PROTOCOL

EXPLAINATION :
let me build you a vilsual story:
you(BROWSER):Sitting at the table hungry for a webpage
(HTTP): the waiter that carries your order & food 
server(KICHEN): stores & prepares the webpage for you

## how it actually works - step by step

step 1: ----REQUEST:
when you type google.com--> your browser send a REQUEST:"Hey server, please give me the google homepage!"

step 2: ----SERVER GETS IT
The server hears your oder, goes to find the page files(HTML,images, etc) and prepares them.

step 3: ----RESPONSE
server sends back a RESPONSE: a statue code (like 200 Ok = "here's your food!") + the actual webpage.

step 4: ----YOU SEE THE PAGE
your browser reads the HMTL & CSS files and draws the webpage on your screen,Done!

## HTTP status codes -- the waiters reply

# 1
200 ok : Here's your food (DONE)
# 2
301 Redirect: we moved! Go here(reloading)
# 3
404 Not found: that's dish is'nt here(not found)
# 4
500 Server Error: kitchen's on fire($)

# HTTP VS HTTPS -- what's the S for

## HTTP:
LIKE SENDING A POSTCARD - anyone can read it (<!>)

## HTTS:
LIKE A SEALED ENVELOP - only you & server can read it(<!!>)

## MORE EXPLAINATION ON HTTP:

HTTP stands for HYPERTEXT TRANSFER PROTOCOL:= it is the language that browsers and server use to communicate.
EXAMPLE:
1* YOU OPEN A BROWSER
2* YOU TYPE google.com.
3* YOUR BROWSER SENDS AN (HTTP REQUEST)
4* GOOGLE'S SERVER SENDS AND (HTTP RESPONSE)
5* THE BROWSER DISPLAYS THE PAGE.

THINK OF IT LIKE ORDERING FOOD:
you = client(browser)
Restuarant = (server)
order = (request)
food = (respons)

## the 4 magic word (HTTP METHODS)
method ----what it means----------like ordering food
GET      "give me something"    "can i see the menu?
POST     "make something new"   "i want a new pizza"
PUT      "change everything"   "change my whole order
DELETE   "Remove something"     "cancel my pizza"


## important Go package for HTTP

PACKAGE -----------------PURPOSE
net/http                   web servers and client
encoding/json              JSON handing
io                          read request bodies
html/template                HTML templates
context                  Request lifecycle management

1* net/http ==> this is the most important package .
it provides everything needed to:
* create servers
* receive requests
* send responses
* make HTTp requests to other servers

2* io ===> the io package is used for reading and writing data streams.
import "io"
in HTTP, it is commonly used to read the request body .
io.ReadAll() collects everything from thst stream.

3* encoding/json ==> (JSON stands for JavaScript Object Notation) it is a text format used to store and exchange data betweeen applications. JSON is a way of organizing information using key-value pairs:
HERE:
"firstName" is a key
"john" is its values

## why is JSON important?
when a browser, mobile app, or frontend talks to a backend server,they often exchange data as JSON.

so JSON is simply the language used to exchange data.
JSON is the language most API use to exchange data.

## what is an API?

APL stands for Application Programming Interface.

in web development, an API is set of HTTP endpoints that allow applications to communicate.

think of an API as a waiter in a restaurant:
Customer
   |
   order
    |
    waiter(API)
    |
    v
    Kitchen(server)
the waitr takes your request to the kitchen and brings back the result.

* think of an API as a doorway to a server.
* the API receives requests and returns responses.
* the API is acting as a middleman.
*  API Endpoints: an endpoint is a URL that performs a task

HTTP = the road that dat travels on (the communication rules)
JSON = the packageing format of what inside the delivery(how data is structured)
API = the post office that decides what can be sent and received

"HTTP is how browsers and servers talk, and when APIs exchange data, they usually pack that data in JSON format."

## what is javascript?
before we start expalining what is javascript let break everything down 
* 1. HTML = the structure of a webpage (the skeleton)
* 2. CSS = the look of a webpage (the clothe)
* 3. JavaScritpt = the behaviour of a webpage(the brain)

javaScript is the programming language of the web. it runs inside your brower --- no installation needed. But it can  also run on a server (that's called Node.js, which we'll get to soon)

```javascript
let name = "Serah";
let age  = 25;

console.log("Hello, my name is " + name);
console.log("I am " + age + " years old");
```

## what's happening here?

* let  = create a box to store a value (called a variable)
* name and age = the names of thise boxes
* console.log() = print something out so you can see it


### in my own word explaination
* 1. 
"When a user loads a page, the browser sends a GET request to the server using HTTP, and the server sends back a 200 OK with the page content."

```javascript
fetch("https://jsonplaceholder.typicode.com/users/1")   // 1. Go get this
  .then(response => response.json())                     // 2. Then unpack the JSON
  .then(data => console.log(data))                       // 3. Then show me the data
```


Reading it line by line like English:

* Line 1 → "Go to this URL and make a GET request"
* Line 2 → "When the response arrives, open it and read the JSON inside"
* Line 3 → "When the JSON is ready, print it out"

# what does response.json()do?



