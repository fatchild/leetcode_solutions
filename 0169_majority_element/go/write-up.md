# Problem: 169. Majority Element

## Initial Approach

> solution_1

Using a hashmap we can attribute the number of occurrences as the value to the element as the key. A loop through the hashmap will reveal the key with the highest value.

### Complexity

> Time: O(2n) We also loop through a map with the same number of key value pairs as elements in nums. 
> Space: O(?) More than 1 as we are using hashmaps (maps).

## Outcome

- Runtime `0ms` beats 100%
- Memory `8.42MB` beats 71.43%
- Time `00:40:00`

## Improvements

> solution_2

The problems description asks if the problem can be solved in linear time and O(1) space complexity.

We can get it down to O(n) by storing the highest occurrences value as we go.

### Complexity

> Time: O(n)
> Space: O(n)

## Outcome

- Runtime `0ms` beats 100%
- Memory `9.14MB` beats 5.48%
- Time `00:09:00`

## Improvements After Looking at Other Answers

> solution_3

Thanks to [lacfo](https://leetcode.com/u/lacfo/) for the best answer and providing the O(1) space complexity.

There is always a majority element. Because of this, we can observe that an array must have adjacent values which are the majority. It may have adjacent values which are not majority, but the will be less than the overall number of occurrences of the majority value.

This is known as the [Boyer Moore voting algorithm](https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm).

## Outcome

- Runtime `0ms` beats 100%
- Memory `9.08MB` beats 7.84%
- Time `00:15:00`