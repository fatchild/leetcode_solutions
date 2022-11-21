var longestCommonPrefix = function(strs) {
    if (strs.length == 0 || strs[0] == "") { return "" }
    if (strs.length == 1) { return strs[0] }

    let lcp = ""

    for (let i = 0; i < strs[0].length; i++){
        let hold = strs[0][i]

        for (let j = 0; j < strs.length; j++){
            if (strs[j][i] == undefined || strs[j][i] != hold[0]) { return lcp }
        }
        lcp += hold
    }
    return lcp
};

let testString1 = ["flower","flow","flight"]
let testStrings2 = ["dog","racecar","car"]

console.assert(longestCommonPrefix(testString1) === "fl", 1);
console.assert(longestCommonPrefix(testStrings2) === "", 2);
console.assert(longestCommonPrefix([""]) === "", 3);
console.assert(longestCommonPrefix(["a"]) === "a", 4);
console.assert(longestCommonPrefix(["flower","flower","flower","flower"]) === "flower", 5);