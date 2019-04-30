Bradley-Terry Model
===================

The [Bradley-Terry
Model](https://en.wikipedia.org/wiki/Bradley%E2%80%93Terry_model) is useful when
comparing the results of a bunch of pairings to see what item is most relevant
or best choice. 

For example, given 3 players and four games like so:

* Document 2 was chosen over Document 1
* Document 2 was chosen over Document 3
* Document 3 was chosen over Document 2
* Document 3 was chosen over Document 2

The most relevance scores for each document are:

* Document 1: **0**
* Document 2: **0.38**
* Document 3: **0.64**

So the most relevant document is #3.
