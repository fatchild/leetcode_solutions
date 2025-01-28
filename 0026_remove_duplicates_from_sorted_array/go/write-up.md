# Problem: 26. Remove Duplicates from Sorted Array

## Initial Approach

> solution_1

Similar to 28 remove element, iterate through the array from index `1` storing index `0` first outside the loop. compare each element of the array with the previous, if they are not the same then assign that element to the first element of the array which has not yet been assigned, this is controlled with the counter `l`.

### Time Complexity

> O(n) will always iterate through all elements of the input array `nums`

## Space Complexity

> O(1) Mutates the array in place

## Outcome

- Runtime `0` beats 100%
- Memory `6.34` beats 27.31%
- Time `00:04:35`

## Improvements

`l` and `k` are both always the same as they are incremented at the same time. They conceptually represent different things but the integers are always the same, so we can just use `k`.

The complexity of the program stays the same. There is memory usage improvement.

## Outcome after Improvements

- Runtime `0ms` beats 100%
- Memory `6.20MB` beats 95.73%
- Time `00:00:30`
