function hourseRobber (arr) {
  const length = arr.length
  if (length === 0) return 0
  if (length === 1) return arr[0]
  if (length === 2) return Math.max(arr[0], arr[1])

  let memozation = []
  memozation[0] = arr[0]
  memozation[1] = Math.max(arr[1], arr[0])

  for (let i=2; i<length; i++) {
    memozation[i] = Math.max((arr[i]+memozation[i-2]), memozation[i-1])
  }
  
  return memozation[length-1]
}

console.log(hourseRobber([5,3,3,9]))