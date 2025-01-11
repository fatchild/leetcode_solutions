# Problem: 88. Merge Sorted Array

## Initial Approach

> solution_1

Combining two arrays as a single sorted ascending array. I used bubble sort initially. Just to get the ball rolling. To my surprise leetcode seems to think that my answer beat 100% of other submissions as it has a runtime of 0ms. However the space complexity is higher, the memory usage is 4.4 MB.

### Time Complexity

> Best: O(n) this will only be the case if the two lists append in the correct order.
> Worst: O(n^2) this will be the case if the list is completely out of order i.e. everything is backwards, but because the lists are ordered themselves it should never be this slow.

### Space Complexity

> O(1) bubble sort requires a consistent amount of space to store additional variables.

## Improvements

I was planning on using a better sorting algorithm but I don't think it's needed. We can loop through the array nums1 which is m + n and will give us a time complexity of O(m + n).

> solution_2

This works and it's running in O(m+n), but the leetcode analysis says it runs in 1ms and only beats 3.86% of other answers. I think this is down to my implementation. I have a lot of repeated code and this can be smoothed out.

### Time Complexity

>Best: O(n+m)
>Worst: O(n+m)

A full iteration of n+m will always be performed.

## Improvements: Round 2

So it was print statements holding me back. It's now a solution beating 100% of people in the time domain and 29.83% in the memory domain.