# pqueue 

### Package pqueue provides a simple priority queue implementation

[![GoDoc](https://godoc.org/github.com/menefotto/pqueue?status.svg)](https://godoc.org/github.com/menefotto/pqueue)
[![Build Status](https://travis-ci.org/menefotto/pqueue.svg?branch=master)](https://travis-ci.org/menefotto/pqueue)
[![Coverage Status](https://coveralls.io/repos/github/menefotto/pqueue/badge.svg?branch=master)](https://coveralls.io/github/menefotto/pqueue?branch=master)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


Queue can be either a max or min priority queue depending on the given Comparator function if v[i] < [j] the it's a max priority queue, otherwise if the function is
v[i] > v[j] then it becomes min queue. 

## How do I use it?

In order to use it as max priority queue instantiate a pqueue with the New method
like so:
```
    func less(i, j int, values []interface{}) bool {
        l, _ := values[i].(int)
        r, _ := values[j].(int)
        return l < r
    }

    pq := pqueue.New(less)

```

Otherwise if one wants to create min priority queue instantiate qa pqueue with the New method and a great Comparator function like so:
```
    func great(i ,j int, values []interface{}) bool {
        l, _ := values[i].(int)
        r, _ := values[j].(int)
        return l > r
    }

    pq := pqueue.New(great)
```

Type casting inside the Comparator function is essential since comparison cannot be
performed on interfaces.
To extract the largest or smallest value ( depending on the queue ), call the Top method or to extract and remove it call PopTop, like so:

```
    v := pq.PopTop() // removes the values
    v := pq.Top() // gives a copy of the value not removing it from the queue
```

### Important
- Nothing to add 

### Philosophy
This software is developed following the "mantra" keep it simple, stupid or better 
known as KISS.

### Thank You Notes
Implementation follows closely the implementation given in "Algorithms by Robert Sedgewick".


