# Problem: 121. Best Time to Buy and Sell Stock

## Approach

> solution_1

I can see 5 options for decision when traversing the input prices and deciding which prices we want to to buy and sell at.

1. valid new profit of adjacent days which is better than current profit
2. always find the lowest value as a potential buy value
3. check if the new sell value is a better sell option
4. check if the new low option is valid and more profitable - with respect to current sell value
5. check if the new low option is valid and more profitable - with respect to next sell value

This approach stores pointers to the lowest values, highest values, the best combination of values and the indexes for all of the aforementioned pointers.

### Complexity

> Time: O(n) We iterate the input array once. 
> Space: O(1) Space is constant but arguably there are too many variables.

## Outcome

- Runtime `0ms` beats 100%
- Memory `10.57MB` beats 34.28%
- Time `05:00:00`

## Improvements

> solution_2

After looking at other solutions I can see one main approach. It means my precious approach could be called 4 too many pointers.

In this version we are combining a lot of the logic of my previous answer and making some efficiency savings by taking out the tracking of things which do not need to be tracked.

### Complexity

> Time: O(1)
> Space: O(n)

Although the complexity is the same as my initial solution I think that this version will have a much higher chance at improving on the time complexity by making less decisions and only tracking minimal data.

## Outcome

- Runtime `0ms` beats %
- Memory `9.61MB` beats %
- Time `00:15:00`
