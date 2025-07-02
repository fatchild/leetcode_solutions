# Problem: 238. Product of Array Except Self

## Approach

> solution_1

We can create a prefix array. This is made in O(n) time complexity. Then iterate through the prefix array and take arr[i] of the original array away from preArr[len(preArr)-1] and assign that to preArr[i], building up the new array. This will be in O(2n) linear time. But will require two loops, I imagine we can make this faster. This will use the division operator which is not allowed.

Hint 1 suggests a suffix.

### Complexity

> Time: O() 
> Space: O() 

## Outcome

- Runtime `ms` beats %
- Memory `.MB` beats %
- Time `00:00:00`