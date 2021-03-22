### Core: PostgreSQL

* Task completed by: Bruno, Bohdan
* Reviewer: Artem Grukhal

* Was the stated problem solved in an acceptable manner?

- Yes, everything works as expected.

* What has been done well and why?

- Error handling returns meaningful messages.


* What is not well implemented and why?

- Probably, there could be more ORM-like approach, instead of writing plain SQL.


* How easy is it to understand?

- The code is self-explanatory.

* Recommendations to simplify

- Use ORM.


* Did the implementers use anything special?

- SQL scripts sometimes were rather complicated, but work well.


* Anything else funny?

- Its difficult to say, that there is anything funny in databases.

### Core: MongoDB

* Task completed by: Artem, Bohdan
* Reviewer: 

* Was the stated problem solved in an acceptable manner?

-

* What has been done well and why?

-


* What is not well implemented and why?

-


* How easy is it to understand?

-

* Recommendations to simplify

-


* Did the implementers use anything special?

-


* Anything else funny?

-

### Core: Redis

* Task completed by: Bohdan
* Reviewer: 

* Was the stated problem solved in an acceptable manner?

-

* What has been done well and why?

-


* What is not well implemented and why?

-


* How easy is it to understand?

-

* Recommendations to simplify

-


* Did the implementers use anything special?

-


* Anything else funny?

-

### Core: HTTP

* Task completed by: Bohdan
* Reviewer: 

* Was the stated problem solved in an acceptable manner?

-

* What has been done well and why?

-


* What is not well implemented and why?

-


* How easy is it to understand?

-

* Recommendations to simplify

-


* Did the implementers use anything special?

-


* Anything else funny?

-

### Core: WebSocket

* Task completed by: Artem, Bohdan
* Reviewer: 

* Was the stated problem solved in an acceptable manner?

-

* What has been done well and why?

-


* What is not well implemented and why?

-


* How easy is it to understand?

-

* Recommendations to simplify

-


* Did the implementers use anything special?

-


* Anything else funny?

-

### Core: gRPC

* Task completed by: Bruno
* Reviewer: 

* Was the stated problem solved in an acceptable manner?

-

* What has been done well and why?

-


* What is not well implemented and why?

-


* How easy is it to understand?

-

* Recommendations to simplify

-


* Did the implementers use anything special?

-


* Anything else funny?

-


### Testing

* Task completed by: Bruno, Bohdan
* Reviewer: Artem Grukhal

* Was the stated problem solved in an acceptable manner?

- Yes, everything works.

* What has been done well and why?

- Full coverage, everything is running well.


* What is not well implemented and why?

- GRPC test could be a bit better (code style and more test-cases)


* How easy is it to understand?

- Pretty easy. The tests are intuitive.

* Recommendations to simplify

- Make "dummy" variables in a separate file, and use where needed by importing (tests are basically testing the same thing but on different protocols). Also preferred pattern to have another database (like H2) to conduct tests, so they do not affect main DB.


* Did the implementers use anything special?

- I wouldn't say so.


* Anything else funny?

- How Rucy keeps praising GRPC. How there don't seem to be a lot of other comments in the code.

### Documentation: HTTP

* Task completed by: Bohdan
* Reviewer: 

* Was the stated problem solved in an acceptable manner?

-

* What has been done well and why?

-


* What is not well implemented and why?

-


* How easy is it to understand?

-

* Recommendations to simplify

-


* Did the implementers use anything special?

-


* Anything else funny?

-

### Documentation: WebSocket

* Task completed by: Bohdan
* Reviewer: 

* Was the stated problem solved in an acceptable manner?

-

* What has been done well and why?

-


* What is not well implemented and why?

-


* How easy is it to understand?

-

* Recommendations to simplify

-


* Did the implementers use anything special?

-


* Anything else funny?

-

### Documentation: gRPC

* Task completed by: Bruno
* Reviewer: 

* Was the stated problem solved in an acceptable manner?

-

* What has been done well and why?

-


* What is not well implemented and why?

-


* How easy is it to understand?

-

* Recommendations to simplify

-


* Did the implementers use anything special?

-


* Anything else funny?

-

### Deployment

* Task completed by: Bohdan
* Reviewer: 

* Was the stated problem solved in an acceptable manner?

-

* What has been done well and why?

-


* What is not well implemented and why?

-


* How easy is it to understand?

-

* Recommendations to simplify

-


* Did the implementers use anything special?

-


* Anything else funny?

-

