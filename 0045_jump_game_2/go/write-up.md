# Problem: 45. Jump Game

## Approach

> solution_1

This initial approach is another greedy algorithm. It makes the best decision based on the current state. It does not look ahead and does not look backwards. It assumes that there is always a path, and from this tries to find the shortest possible path. This is impressive with a greedy algorithm, to me at least.

### Complexity

> Time: O(n) 
> Space: O(1) 

## Outcome

- Runtime `0ms` beats 100%
- Memory `7.82.MB` beats 64.28%
- Time `01:20:00`

## Approach

> solution_2

I saw that people were submitting DP answers. Dynamic programming answers.

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
