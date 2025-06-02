/*const missingNumber = function(nums){

 nums.sort()
 let i = 0;
  while(i<=nums.length){
    i++
    if(nums[i]-nums[i-1]>1){
      return i
    }
  
  }
  return i
}*/

cosnt missingNumber = function(nums){
n = length.nums;
  const sum = num.reduce((acc, number)=> acc+ number,0);
  return (n*(n+1))/2 - sum

}
console.log(missingNumber([3,0,1]));


