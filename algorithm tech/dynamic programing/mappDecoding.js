/**
 A top secret message containing uppercase letters from 'A' to 'Z' has been encoded as numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26

You are an FBI agent and you need to determine the total number of ways that the message can be decoded.

Since the answer could be very large, take it modulo 109 + 7.
 */

function mapDecoding(message) {
  let mapSecretKey = {}, count = 1

  for (let i = 1; i <= 26; i++) {
    mapSecretKey[i] = String.fromCharCode(i + 64)
  }

  if (message.length === 1) {
    if (message[0] == 0) {
      return 0
    }
  }

  if (message.length == 2) {
    if (message[0] == 0 || message[1] == 0) {
      if (!mapSecretKey[`${message[0]}${message[1]}`]) {
        return 0
      }
      return 1
    } else {
      if (mapSecretKey[`${message[0]}${message[1]}`]) {
        return 2
      }
      return 1
    }
  }

  for (let i = 1; i < message.length - 1; i++) {

    if (message[i] == 0) {
      if (!mapSecretKey[`${message[i - 1]}${message[i]}`] && !mapSecretKey[`${message[i]}${message[i + 1]}`]) {
        return 0
      }
    }

    if (mapSecretKey[`${message[i - 1]}${message[i]}`] && mapSecretKey[`${message[i]}${message[i + 1]}`]) {
      count += 2
    } else if (mapSecretKey[`${message[i]}${message[i + 1]}`]) {
      count++
    }
  }

  return count
}

console.log(mapDecoding('11115112112'))