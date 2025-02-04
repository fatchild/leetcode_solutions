# Problem: 122. Best Time to Buy and Sell Stocks II

## Approach

> solution_1

I think we can start with the two pointer approach from the first iteration of this task. I've noticed that all we need to track is the positive sales between two pointers. The accumulative profit of buying and selling on any day which is positive is the same as waiting to sell at the highest point. The code implementation is simply that we need to test for a positive outcome of buying today and selling tomorrow, then add that profit to an accumulator. Done. Why was this considered a `medium` when the easy one took me so long and had far more complicated code and logic associated with it? Maybe I got lucky and found the trick.

### Complexity

> Time: O(n) Will need to loop through the entire array `n`.
> Space: O(1) The space complexity is constant as no more space is used with a growth in input.

## Outcome

- Runtime `0ms` beats 100%
- Memory `5.14MB` beats 45.93%
- Time `00:10:00`