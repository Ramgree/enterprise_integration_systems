## Here are all Reviews

### Task 1

* Task completed by: Bruno & Jekaterina
* Reviewer: Bohdan

* Was the stated problem solved in an acceptable manner?

Yes, the backend works as expected, required functionality is implemented, some potential errors are handled.


* What has been done well and why?

The code is well-written.
For every request, it is checked whether the received key exists.


* What is not well implemented and why?

Type declarations are all over the place, I think it would be reasonable to add a "model" package. 
The TODO status is a string, and backend actually allows to set it to anything via a request, instead of just marking it as completed. Update TODO endpoint has an unnecessary path parameter {id}.


* How easy is it to understand?

I think the code is quite clear.


* Recommendations to simplify

Changing the string "Status" to boolean "isCompleted" would simplify some logic on both front end and back end.



* Did the implementers use anything special?

Some graph algorithms are probably the most special thing.


* Anything else funny?

At some point of development, I found it kinda funny when the create TODO endpoint was expecting the whole object (including the ID).



### Task 2 

* Task completed by: Bohdan & Artem
* Reviewer: Bruno
* Was the stated problem solved in an acceptable manner?

Yes, the front end works as expected, and it has all features of CRUD and the integration with the dependency graph.

* What has been done well and why?

The backend service interfacing was done well. It seems that they tried to retain parity with the backend tests.

* What is not well implemented and why?

The integration tests are not good, and they very heavily conflict with the functional tests, which caused me headache in order to fix them.

* How easy is it to understand?

The code was clear and self documenting.

* Recommendations to simplify

Do not install the cli. Compile everything to a binary, as it's meant to be.

* Did the implementers use anything special?

No.

* Anything else funny?

I thought it was amusing how they didn't seem to be aware that a go project can be compiled to a binary, making their lifes considerably more complicated than it should have been.

### Task 3 

* Task completed by: 1, 2, 3, 5 by Bruno & Jekaterina and 4 by Bohdan
* Reviewer:

### Task 4 

* Task completed by: Bruno & Jekaterina
* Reviewer:

### Task 5 

* Task completed by: Bruno & Jekaterina
* Reviwer: Bohdan
* Was the stated problem solved in an acceptable manner?

Yes. Services are served as required.

* What has been done well and why?

Traefik proxy was setup well, everything is easy to run and verify that it works.


* What is not well implemented and why?

At the time of the review the documentation was a bit incomplete, in my opinion.


* How easy is it to understand?

Quite easy, thanks to traefik.


* Recommendations to simplify

Add examples of all commands needed to run and test the solution (done).

* Did the implementers use anything special?

No.

* Anything else funny?

No.


