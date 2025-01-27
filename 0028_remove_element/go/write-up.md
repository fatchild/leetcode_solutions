# Problem: 28. Remove elements

## Initial Approach

> solution_1

I think the simplicity of this answer comes down to the fact that there are basically no constraints. In the real world I'm not sure it would be acceptable to mutate an array to contain only the numbers you want at the start and not care about what comes after.

This program reads the array. If the element read is equal to the remove value then a counter is incremented. If the element is not equal to the remove value then we assign the element to the first index of the array, which has not yet been reassigned.

### Time Complexity

> O(n) loops once around the array no matter what happens

## Space Complexity

> O(1) as the program mutates the array itself

## Outcome

- Runtime `0ms` beats 100%
- Memory `4.09MB` beats 69.08%
- Time `~2h`
