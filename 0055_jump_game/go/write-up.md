# Problem: 55. Jump Game

## Approach

> solution_1

My first approach was marred by misunderstanding the problem. I thought we had to use the full value of each element.

### Complexity

> Time: O(n) 
> Space: O(1) 

## Outcome

Incorrect...

## Approach

> solution_2

I looked into greedy algorithms as that was one of the tags. These types of algorithms make the best choice to find a global optimum without storing any extra information about previous choices. They are not used for finding the optimal path, but the question only asks if it can be done. So if the greedy algorithm can show that it's possible without finding the best path we can solve the problem quickly even if the answer it finds isn't optimal.

I also saw an entry for solving this problem, I didn't read the whole article but I saw it's tagline about starting from the end. This basically solved it.

Back to front greedy method I shall call it until I know what the actual algorithm is called.

Take the last element of the array and call that `position` check if the element before it can "get to" `position`. If it can then we have a valid path (but maybe not optimal). If it can't keep `position` at the last element and check to see if the element before it can reach `position` if nothing can reach a `position` element then there is no path. If however, the `position` element can be reached, then we make `position` that next element and check to see if there is a path to the current `position` element.

### Complexity

> Time: O(n) 
> Space: O(1) 

## Outcome

- Runtime `0ms` beats 100%
- Memory `8.98.MB` beats 70.67%
- Time `00:20:00`

## Approach

> solution_3

There is a more elegant way to express what I was trying to do above. In both approaches we are checking if a path can be made back to front. But in this implementation you just track the jumps needed, rather than the value of the index.

### Complexity

> Time: O(n) 
> Space: O(1) Constant but a smaller constant.

## Outcome

- Runtime `0ms` beats 100%
- Memory `8.83MB` beats 84.04%
- Time `00:10:00`
