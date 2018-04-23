# priority queue

### Package pqueue provides a simple implementation

[![GoDoc](https://godoc.org/github.com/menefotto/priority_queue?status.svg)](https://godoc.org/github.com/menefotto/priority_queue)
[![Build Status](https://travis-ci.org/menefotto/priority_queue.svg?branch=master)](https://travis-ci.org/menefotto/priority_queue)
[![Coverage Status](https://coveralls.io/repos/github/menefotto/priority_queue/badge.svg?branch=master)](https://coveralls.io/github/menefotto/priority_queue?branch=master)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


Queue can be either a max or min priority queue depending on the given Comparator function if v[i] < [j] the it's a max priority queue, otherwise if the function is
v[i] > v[j] then it becomes min queue. 
The only thing to be done in order to use the queue is to create a function that satisfies the Comparator function type signature, because the use of interfaces to
achieve a general purpose data structures the compare function has to do type casting in order use the <,> operators ( cannot be done on interfaces but you already know that ).