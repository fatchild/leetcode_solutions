# Problem: 26. Remove Duplicates from Sorted Array

## Initial Approach

> solution_1

The differences between this and the first version of the question Q.26 are that the input array is in non-decreasing order and unique elements appear at most twice.

I've implemented a new counter which increments based on the number of times an element has been assigned. If this is less than 2 then we can add duplicated numbers.

### Time Complexity

> O(n) will always iterate through all elements of the input array `nums`

## Space Complexity

> O(1) Mutates the array in place

## Outcome

- Runtime `5ms` beats 28.17%
- Memory `4.61` beats 78.66%
- Time `00:40:00`

## Improvements

Try swapping the if statement clauses as in the test set, duplications may be more prevalent. This does not work.

If you just run it again the runtime improves and the memory usage goes up. I guess this is because the test set is generated on the fly and isn't always the same. So sometimes your code will perform better than others, depending on the test set's complexity and how your code deals with the characteristics which are more prevalent sometimes.

## Outcome after Simply Rerunning

- Runtime `0ms` beats 100%
- Memory `4.78MB` beats 26.97%
- Time `00:10:00`

After running it a few more times, the results are consistently jumping between the ones above and the first results outcome.

## Improvements

> solution_2

As the length of the `nums` array is evaluated twice I stored the result in a variable. In the case where there is only one element in the incoming array, I simply return, but now the function does this before it does anything else.

## Outcome after Actual Code Improvements

- Best
- Runtime `0ms` beats 100%
- Memory `4.67MB` beats 78.66%

- Worst
- Runtime `4ms` beats 58.79%
- Memory `4.72MB` beats 27.09%

- Time `00:10:00`

- Overall Time `00:45:00`

## Better Solution After Looking at Answers

> solution_3

You don't need counters as you can just look back through the array by two indexes.

This is less complex. But you don't get any speed or memory gains. Still a better answer, but *only* from a code readability perspective. I've tweaked it slightly to improve memory usage but to the detriment of the readability. Worth it? I don't know. Splitting hairs probably. Scratch that, it doesn't improve the memory usage. I just made it worse...

- Runtime `0ms` beats 100%
- Memory `4.67MB` beats 78.66%
