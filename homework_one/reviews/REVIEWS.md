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




* Reviewer: Jekaterina
* Was the stated problem solved in an acceptable manner?

Yes, it works as it meant to be.

* What has been done well and why?

Everything works, it's cool.

* What is not well implemented and why?

Can't really say anything new, the problem with the tests was mentioned above.

* How easy is it to understand?

Everything is nice and clear.

* Recommendations to simplify

Pay more attention when writing tests.

* Did the implementers use anything special?

No.

* Anything else funny?

Still no :< 

### Task 3 

* Task completed by: 1, 2, 3, 5 by Bruno & Jekaterina and 4 by Bohdan
* Reviewer: Artem

* Was the stated problem solved in an acceptable manner?

Yes, the tests are present as required, and whats even better they pass. The applicaition is dockerized and docker-compose is able to start tests before running.

* What has been done well and why?

I would say, this task is pretty straightforward, so since everything works - everything was done well.

* What is not well implemented and why?

Again, since everything works as expected, and there is really not that much room for imagination, everything seems well. Probably the only thing that could be better is documentation on GH page.

* How easy is it to understand?

Everything is structured nicely and is similar to what we did on labs, so it is not difficult to understand.

* Recommendations to simplify

The only thing I would add, is the explanation of how the CLI issue was resolved.

* Did the implementers use anything special?

No, nothing special.

* Anything else funny?

The amount of anger Bruno developed during this task.

### Task 4 

* Task completed by: Bruno & Jekaterina
* Reviewer: Artem

* Was the stated problem solved in an acceptable manner?

Yes, the image is present in docker registry and everything was optimized to authors' best capabilities.

* What has been done well and why?

The amount of space that was saved by the base image (Alpine).

* What is not well implemented and why?

This is even more straightforward task and it is basically based on following a tutorial. Since everything works correctly, there are no flaws I could find.

* How easy is it to understand?

It is basically following the tutorial of creating a registry. The only difficulty could have been creating a cloud machine, but that shouldn't be difficult.

* Recommendations to simplify

I don't have any recommendations.

* Did the implementers use anything special?

Using the Google Cloud was quite unusual, given that we have access to Hetzner and some have access to OpenStack VMs.

* Anything else funny?

How the GH page has 140 times size reduction in size, not taking into account the weight of the project and Golang compiler.

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


