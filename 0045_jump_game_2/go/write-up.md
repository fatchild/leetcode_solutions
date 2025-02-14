# Problem: 45. Jump Game

## Approach

> solution_1

This initial approach is another greedy algorithm. It makes the best decision based on the current state. It does not look ahead and does not look backwards. It assumes that there is always a path, and from this tries to find the shortest possible path. This is impressive with a greedy algorithm, to me at least.

Although this is storing some information about the state of the previous decisions in order to aid an optimal outcome of the problem as a whole it does display the elements of a greedy algorithm. If it backtracked then it wouldn't be greedy. It does make optimal decisions about the current state. It does however change it's view of the optimal solution based on some saved state, which is more akin to dynamic programming. Overall I think this is close enough to a greedy approach for it to be considered as such. Until my knowledge or changes or someone proves me otherwise.

### Complexity

> Time: O(n) 
> Space: O(1) 

## Outcome

- Runtime `0ms` beats 100%
- Memory `7.82MB` beats 64.28%
- Time `01:20:00`

## Approach

> solution_2

I saw that people were submitting DP answers. Dynamic programming answers. Dynamic Programming is an approach were you break down problems into sub problems in order to be more efficient than a normal recursive brute force approach.

This approach will use the memoization method. By doing so you optimize the time complexity from $O(n^2)$ to $O(n)$.

Firstly I will implement an approach top-down approach using recursion.

I started this then realized i'd just implemented a greedy recursive algorithm. It's not actually greedy because it doesn't move through in one pass, but memoization does not help here.

### Complexity

> Time: O(n^2) recursively calling f(n-1)
> Space: O(n)

## Outcome

- Runtime `13ms` beats 34.45%
- Memory `8.37MB` beats 70.67%
- Time `01:00:00`
