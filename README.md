# priority queue

### Simple priority queue implementation

Queue can be either a max or min priority queue depending on the given comparator function if v[i] < [j] the it's a max priority queue, otherwise if the function is
v[i] > v[j] then it becomes min queue. 
The only thing to be done in order to use the queue is to create a function that satisfies the comparator function type signature, because the use of interfaces to
achieve a general purpose data structures the compare function has to do type casting in order use the <,> operators ( cannot be done on interfaces but you already know that ).