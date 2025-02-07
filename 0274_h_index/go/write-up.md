# Problem: 274. H-Index

## Approach

> solution_1

The [wikipedia article about H-Index](https://en.wikipedia.org/wiki/H-index) gives and approach to solve the problem.

1. Sort in non-increasing order.
2. Iterate through the list of citations and find the last entry which has a greater number of citations i.e. value of the element, than the element position.

Let's implement this even though I know there will be a faster solution.

The outcome is slow, but keeps O(1) constant space complexity.

### Complexity

> Time: O(n^2) 
> Space: O(1) Bubble sort is performed in place

## Outcome

- Runtime `18ms` beats 5.06%
- Memory `4.62MB` beats 5.89%
- Time `00:05:00`

## Approach

> solution_2

We can improve on this I am sure. With a better sorting algorithm. There may also be a greedy approach which will complete the task in O(n) time.

One of the three topics listed is counting sort and one of the others is sort so I don't think a greedy approach is possible. Let's improve the sort by performing a [counting sort](https://www.geeksforgeeks.org/counting-sort/).

This is a good result, but I think reverse could be removed and we could probably find the h-index whilst in the counting sort. If so we can create the output array from largest to smallest and check for the index whilst doing so.

Looking back this is a bad thought, because counting sort does not sort in order, so we cannot know if there is a higher h-value without finishing the sort of the array.

### Complexity

> Time: O(n+m+more and yet more?) n is input array and m is output array, but we also need to consider max function, and finding the h-index. This is a linear algorithm but it's got a lot of looping through the various arrays.
> Space: O(n+m) we need to create a new array output array which is of size m, which could be any integer.

## Outcome

- Runtime `0ms` beats 100%
- Memory `4.28MB` beats 42.33%
- Time `01:00:00`

## Approach

> solution_3

Counting sort but optimized.

> The h-value will never be greater than the total number of papers they have submitted.

So they can never have a h-value of 6 if they have only produced 5 papers.

Because of this we can implement only half of a complete counting sort. We don't need to find the value in the input array as we only need the length of the array plus one as the count array. The full counting sort counts up occurrences, works out sum at each element of the count array. Then finally uses that array to rebuild a new sorted array.

We don't need to do all of this, we can create a count array which collects all values in the last element which have a value of greater than the length of the input array plus one. Once assembled we then initialize a counter of `papers` which keeps track of how many papers the author has produced. We then compare the `paper` counter with the number of citations. If the number of papers is greater than or equal to the number of papers with that many citations, then we return that number of citations.

### Complexity

> Time: O(n+m) but with less overhead.
> Space: O(n) not constant time as we need to create a new array, count array for the sorting and comparisons.

## Outcome

- Runtime `0ms` beats 100%
- Memory `4.29MB` beats 42.33%
- Time `02:00:00`
