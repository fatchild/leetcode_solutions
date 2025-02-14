# Problem: 380. Insert Delete GetRandom O(n)

## Approach

> solution_1

I had to look this one up and found an [explanation on GeeksForGeeks](https://www.geeksforgeeks.org/design-a-data-structure-that-supports-insert-delete-getrandom-in-o1-with-duplicates/). 

You can trivially make insert and delete O(1) but then that makes getting a random element harder to do.

I started with just a hashmap which stored the elements and the number of each element. I thought it would be a case of counting the unique element up or down based on the inserts and deletes. But instead a better approach is to store an array of the unique values and store the indexes of those values in the hashmap.

The implementation can be simpler than the one used on GeeksForGeeks because we don't have to track all instances of a unique value. We can track any instance, if there is more than one. I can see another GetRandom(0) where it allowed duplicates.

### Complexity

> Time: O(1) 
> Space: O(n) 

## Outcome

- Runtime `55ms` beats 72.16%
- Memory `45.72MB` beats 90.84%
- Time `04:00:00`