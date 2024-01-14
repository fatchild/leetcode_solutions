/*

2. Palindrome Number

*/

let example1 = 121
let example2 = -121
let example3 = 10


// From the internet
function isPalendrome(x) {
    if (x < 0) { return false }

    let rev = 0
    for(let i = x; i >= 1; i = Math.floor(i/10)){
        // console.log(Math.floor(i/10))
        rev = rev*10 + i%10
    }
    return rev === x
}

console.log(isPalendrome(example1))
console.log(isPalendrome(example2))
console.log(isPalendrome(example3))


// Second go

// A negative number can never be a palindrome
// Reverse the number and then compare the two
// Source: https://medium.com/@ManBearPigCode/how-to-reverse-a-number-mathematically-97c556626ec6

function isPalendrome2(x) {
    if ( x < 0 ) { return false }

    let number = x; // Store a pristine version of the number to check the equality at the end
    let reverse = 0 // Initialize a variable to store the reversed number

    // Perform the following mathematical operation until the number is completely reversed
    while(x > 0){

        // Isolate the last digit in the number
        let lastDigit = x % 10

        // Append the last digit 
        reverse = (reverse * 10) + lastDigit // Move the last digit to the left and add on the newest last digit

        // Remove the last digit from the number
        x = Math.floor(x / 10)
    }

    return reverse === number
}

console.log(isPalendrome2(example1))
console.log(isPalendrome2(example2))
console.log(isPalendrome2(example3))


// Third go - for loops are more efficient than while loops

// A negative number can never be a palindrome
// Reverse the number and then compare the two
// Source: https://medium.com/@ManBearPigCode/how-to-reverse-a-number-mathematically-97c556626ec6

function isPalendrome3(x) {
    if ( x < 0 ) { return false }

    let number = x; // Store a pristine version of the number to check the equality at the end
    let reverse = 0 // Initialize a variable to store the reversed number

    // Perform the following mathematical operation until the number is completely reversed
    for (let i = x; x > 0; x = Math.floor(x / 10)){

        // Isolate the last digit in the number
        let lastDigit = x % 10

        // Append the last digit 
        reverse = (reverse * 10) + lastDigit // Move the last digit to the left and add on the newest last digit
    }

    return reverse === number
}

console.log(isPalendrome3(example1))
console.log(isPalendrome3(example2))
console.log(isPalendrome3(example3))


// Fourth go - post submission

// A negative number can never be a palindrome
// Reverse the number and then compare the two
// Source: https://medium.com/@ManBearPigCode/how-to-reverse-a-number-mathematically-97c556626ec6
// 0 us always a palindrome a number ending in 0 cannot be a palindrome

function isPalendrome4(x) {
    if ( x < 0 || ( x % 10 == 0 && x != 0 )) { return false }

    let number = x; // Store a pristine version of the number to check the equality at the end
    let reverse = 0 // Initialize a variable to store the reversed number

    // Perform the following mathematical operation until the number is completely reversed
    for (let i = x; x > 0; x = Math.floor(x / 10)){

        // Isolate the last digit in the number
        let lastDigit = i % 10

        // Append the last digit 
        reverse = (reverse * 10) + lastDigit // Move the last digit to the left and add on the newest last digit
    }

    return reverse === number
}

console.time()
console.log(isPalendrome4(example1))
console.log(isPalendrome4(example2))
console.log(isPalendrome4(example3))
console.timeEnd()