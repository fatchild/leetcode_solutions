


var isValid = function(s) {
    
};



console.assert(isValid("()") == true, "1")
console.assert(isValid("()[]{}") == true, "2")
console.assert(isValid("(]") == false, "3")
console.assert(isValid("(([]))") == true, "4")
console.assert(isValid("()][") == false, "5")
console.assert(isValid("(][)") == false, "6")