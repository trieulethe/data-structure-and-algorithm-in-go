// Given a sorted integer array that does not contain any duplicates, return a summary of the number ranges it contains.
// For nums = [-1, 0, 1, 2, 6, 7, 9], the output should be
// composeRanges(nums) = ["-1->2", "6->7", "9"]

function composeRanges(nums) {
  let result = []

  const length = nums.length
  if (length === 0) return []
  if (length === 1) return [nums[0].toString()]
  let headNumber = nums[0]
  let tailNumber = nums[0]
  result.push(tailNumber.toString())

  for (let i=1; i<length; i++) {
    if (tailNumber + 1 === nums[i]) {


      // if (result[result.length - 1][0] === '-') {
        // result[result.length - 1] = `${result[result.length - 1][0]}${result[result.length - 1][1]}->${nums[i]}`
      // } else {
        // result[result.length - 1] = `${result[result.length - 1][0]}->${nums[i].toString()}`
      // }
      
      tailNumber = nums[i]
    } else {
      result.push(nums[i].toString())
      tailNumber = nums[i]
    }
  }

  return result
} 

console.log(composeRanges([-1, 0, 1, 2, -6, 7, 9]))
