### Core: PostgreSQL

* Task completed by: Bruno, Bohdan
* Reviewer: Artem Grukhal

* Was the stated problem solved in an acceptable manner?

Yes, everything works as expected.

* What has been done well and why?

- Error handling returns meaningful messages.


* What is not well implemented and why?

Probably, there could be more ORM-like approach, instead of writing plain SQL.


* How easy is it to understand?

The code is self-explanatory.

* Recommendations to simplify

Use ORM.


* Did the implementers use anything special?

SQL scripts sometimes were rather complicated, but work well.


* Anything else funny?

Its difficult to say, that there is anything funny in databases.

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
* Reviewer: Jekaterina

* Was the stated problem solved in an acceptable manner?

Yup, and it works just fine.

* What has been done well and why?

Code is clear and human readable, comments are also of help.


* What is not well implemented and why?

I think it's just fine since it's working as needed.


* How easy is it to understand?

Easy enough.

* Recommendations to simplify

None.


* Did the implementers use anything special?

Nope.


* Anything else funny?

No.

### Core: gRPC

* Task completed by: Bruno, Jekaterina
* Reviewer: Bohdan

* Was the stated problem solved in an acceptable manner?

Yes, gRPC interface works as expected (verified by tests) and the code looks alright.

* What has been done well and why?

Personally, I think gRPC is not the simplest thing, so good job for adding it to our project.


* What is not well implemented and why?

I don't see any problems, everything seems to be pretty consistent.


* How easy is it to understand?

Some parts are not easy for me, but all the abstractions are very clear.

* Recommendations to simplify

None.


* Did the implementers use anything special?

No.


* Anything else funny?

No.


### Testing

* Task completed by: Bruno, Bohdan
* Reviewer: Artem Grukhal

* Was the stated problem solved in an acceptable manner?

Yes, everything works.

* What has been done well and why?

Full coverage, everything is running well.


* What is not well implemented and why?

GRPC test could be a bit better (code style and more test-cases)


* How easy is it to understand?

Pretty easy. The tests are intuitive.

* Recommendations to simplify

Make "dummy" variables in a separate file, and use where needed by importing (tests are basically testing the same thing but on different protocols). Also preferred pattern to have another database (like H2) to conduct tests, so they do not affect main DB.


* Did the implementers use anything special?

I wouldn't say so.


* Anything else funny?

How Rucy keeps praising GRPC. How there don't seem to be a lot of other comments in the code.

### Documentation: HTTP

* Task completed by: Bohdan
* Reviewer: Jekaterina

* Was the stated problem solved in an acceptable manner?

Yes.

* What has been done well and why?

Everything needed for setup explained in clear manner.


* What is not well implemented and why?

Can't say anything, everything is cool.


* How easy is it to understand?

For me it's just fine.

* Recommendations to simplify

None.


* Did the implementers use anything special?

No.


* Anything else funny?

Can't really say anything fun about documentation stuff.

### Documentation: WebSocket

* Task completed by: Bohdan
* Reviewer: Jekaterina

* Was the stated problem solved in an acceptable manner?

Yes, everything seems nice and clear.

* What has been done well and why?

Explanations with the examples. 


* What is not well implemented and why?

Hard to say since it seems that there's nothing else to write here.


* How easy is it to understand?

Super clear.

* Recommendations to simplify

None.


* Did the implementers use anything special?

No.


* Anything else funny?

Can't really say anything fun about documentation stuff x2.

### Documentation: gRPC

* Task completed by: Bruno
* Reviewer: Bohdan

* Was the stated problem solved in an acceptable manner?

Since I am not experienced with requirements for implementing gRPC clients I would say yes. Proto file is indeed self-explanatory.

* What has been done well and why?

Not much to talk about here.


* What is not well implemented and why?

Some human language descriptions would be nice.


* How easy is it to understand?

Easy for me.

* Recommendations to simplify

None.


* Did the implementers use anything special?

No.


* Anything else funny?

No.

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

