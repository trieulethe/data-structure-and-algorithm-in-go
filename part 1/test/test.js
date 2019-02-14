
const item = "i love you";
let hash = 0;

for (let charIndex = 0; charIndex < item.length; charIndex += 1) {
   const char = item.charCodeAt(charIndex);
   hash = (hash << 5) + hash + char;
   console.log("hash before:", hash)
   // hash &= hash; // Convert to 32bit integer
   console.log("hash after:", hash)
   hash = Math.abs(hash);
}

console.log("hash:", hash);