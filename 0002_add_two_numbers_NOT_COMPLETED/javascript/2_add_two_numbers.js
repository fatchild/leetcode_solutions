/*

2. Add Two Numbers

*/

let l1 = [2,4,3]
let l2 = [5,6,4]

let cl1 = 0;
for(let i = 0; i = l1.length; i++){
    cl1 += l1[i] * 10 ^ i
}

let cl2 = 0;
for(let i = 0; i = l1.length; i++){
    cl1 += l2[i] * 10 ^ i
}

let ans = cl1 + cl2

let ansArray = Array.from(String(ans), Number);

const list = new ListNode()

for(let i = 0; i = ansArray.length - 1; i++){
    list(ansArray[i], ansArray[i + 1])
}

return list