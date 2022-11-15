let example1 = "III"
let example2 = "LVIII"
let example3 = "MCMXCIV"

let example4 = "IV"
let example5 = "IX"
let example6 = "XL"
let example7 = "XC"
let example8 = "CD"
let example9 = "CM"

var romanToInt_firstGo = function(s) {
    // init a total
    let t = 0;

    const additiveMap = new Map();
    additiveMap.set("I", 1)
    additiveMap.set("V", 5)
    additiveMap.set("X", 10)
    additiveMap.set("L", 50)
    additiveMap.set("C", 100)
    additiveMap.set("D", 500)
    additiveMap.set("M", 1000)

    // 1. Remove the subtractive cases using regex
    let subtractivePatterns = ["IV", "IX", "XL", "XC", "CD", "CM"]
    let subtractionValues = [4, 9, 40, 90, 400, 900]

    for (let i = 0; i < subtractivePatterns.length; i++){
        let currentPattern = subtractivePatterns[i]
        if (s.match(currentPattern)){
            s = s.replace(currentPattern, "")
            t += subtractionValues[i]
        }
    }

    // 2. Add up the remaining using a hashmap
    for (let i = 0; i < s.length; i++){
        let letter = s.charAt(i)
        if (additiveMap.has(letter)){ t += additiveMap.get(letter) }
    }

    // Return the total
    return t
};

var romanToInt_secondGo = function(s) {
    // init a total
    let t = 0;

    const additiveMap = new Map();
    additiveMap.set("I", 1)
    additiveMap.set("V", 5)
    additiveMap.set("X", 10)
    additiveMap.set("L", 50)
    additiveMap.set("C", 100)
    additiveMap.set("D", 500)
    additiveMap.set("M", 1000)

    // 1. Move through the string as an array and check for a signed value before adding it to the total
    for (let i = 0; i < s.length; i++){
        let cVal = additiveMap.get(s[i])
        let nVal = additiveMap.get(s[i+1])
        if (cVal < nVal){
            t -= cVal
        }
        else {
            t += cVal
        }
    }

    // Return the total
    return t
};

var romanToInt = function(s) {
    // init a total
    let t = 0;

    const additiveMap = {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000
    }

    // 1. Move through the string as an array and check for a signed value before adding it to the total
    for (let i = 0; i < s.length; i++){
        let cVal = additiveMap[s[i]]
        let nVal = additiveMap[s[i+1]]
        if (cVal < nVal){
            t -= cVal
        }
        else {
            t += cVal
        }
    }

    // Return the total
    return t
};

var romanToInt_final = function(s) {
    let t = 0;

    const additiveMap = {"I": 1,"V": 5,"X": 10,"L": 50,"C": 100,"D": 500,"M": 1000
    }

    for (let i = 0; i < s.length; i++){
        let [c, n] = [additiveMap[s[i]], additiveMap[s[i+1]]]
        if (c < n) t -= c
        else t += c
    }

    return t
};

/*
Instances where subtraction is used

IV (4)
IX (9)
XL (40)
XC (90)
CD (400)
CM (900)
*/

console.assert(romanToInt(example1) == 3);
console.assert(romanToInt(example2) == 58);
console.assert(romanToInt(example3) == 1994);
console.assert(romanToInt(example4) == 4);
console.assert(romanToInt(example5) == 9);
console.assert(romanToInt(example6) == 40);
console.assert(romanToInt(example7) == 90);
console.assert(romanToInt(example8) == 400);
console.assert(romanToInt(example9) == 900);
