# Problem: 169. Rotate Array

## Approach

> solution_1

My initial approach, which failed, was the following: iterate through the array. Swap elements i and j (which was essentially i+k). Once j reached the end of the array it flips back to it's starting position and you continue. This failed because there are too many edge cases.

I could see that it's possible to nest an array and swap over and over again until you have completed the number of iterations you desire. You could also do it in O(n) time by using O(n) memory by creating a new array. I didn't bother writing these because it's not the challenge.

### Complexity

> Time: O() - If it has worked, this approach would have been both time and memory efficient. 
> Space: O() - Alas.

## Outcome

- Runtime `ms` beats %
- Memory `.MB` beats %
- Time `04:00:00`

## Approach

> solution_2 [juggling algorithm](https://www.geeksforgeeks.org/juggling-algorithm-for-array-rotation/)

I can see after some research that there are a few methods for this.

1. Juggling algorithm: Splitting the array into parts and swapping the array based on those parts
2. Reversing algorithm: partially reversing some of the array in parts. O(n) time, O(1) memory

For the juggling algorithm we are rotating left, which is equal to rotating right where k = k - the length of the array.

### Complexity

> Time: O(n) One time around the array, albeit in groups
> Space: O(1) The array is updated in place so memory is conserved

## Outcome

- Runtime `0ms` beats 100%
- Memory `9.48MB` beats 66.08%
- Time `01:00:00`

## Approach

> solution_3 [Reversing Algorithm](https://www.geeksforgeeks.org/array-rotation/#expected-approach-2-using-reversal-algorithm-on-time-and-o1-space:~:text=The%20idea%20is%20based%20on%20the%20observation)

By splitting the array into n and array length - n we reverse the subsets then reverse the superset i.e. the whole array and then hey presto. The answer is left rotated.

It tripped me up for a while trying to work out why this method is O(n) because I assumed that reverse was iterating through each slice of the array equal to the size of the slice. But it's doing half as many iterations as there are array elements. so subset reversals use O(n/2) and the reverse on the whole array uses O(n/2), together they are O(n). I believe this is the correct understanding. Please let me know if not.

To left rotate we flip the order of the reversals.

### Complexity

> Time: O(n) 
> Space: O(1) 

## Outcome

- Runtime `0ms` beats 100%
- Memory `49.76MB` beats %
- Time `01:00:00`