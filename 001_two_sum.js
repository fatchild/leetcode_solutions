/*

1. Two Sum

*/

// Solution 1
// var twoSum = function(nums, target) {
//     for (let i = 0; i < nums.length; i++){
//         for (let x = i + 1; x < nums.length; x++){
//             if (nums[i] + nums[x] == target && i != x) { return[i, x] }
//         }
//     }
// };

// Solution 2 - Slightly more elegant 
// var twoSum = function(nums, target) {
//     for (let i = 0; i < nums.length; i++){
//         let diff = target - nums[i]
//         for (let x = i + 1; x < nums.length; x++){
//             if (nums[x] == diff) { return[i, x] }
//         }
//     }
// };


let nums = [2,7,11,15]
let target = 13

// const hashmap = new Map();

// for (let i = 0; i < nums.length; i++){
//     let diff = hashmap.get(target - nums[i])
//     if ( diff != undefined ) { console.log([i, diff]) }
//     hashmap.set(nums[i], i);
// }

// const hashmap = new Map();

// for (let i = 0; i < nums.length; i++){
//     if ( hashmap.has(target - nums[i]) ){ return [hashmap.get(target - nums[i]), i] }
//     hashmap.set(nums[i], i)
// }

let seenHash = {}

for (let i = 0; i < nums.length; i++){
    if ( ( target - nums[i] ) in seenHash ){ 
        console.log([ seenHash[target - nums[i]], i]) 
    }
    seenHash[nums[i]] = i
}