/*var singleNumber = function(nums){
  for(let i =0; i<nums.length; i++){
    let check = boolCheck(nums[i],nums,i)
      if(check===true){
        return nums[i]
      }
    }
  }


var boolCheck = function(i,nums,index){
  for (let j=0;j<nums.length;j++){
    if((i==nums[j])&&(index!=j)){
      console.log(i, nums[j])
      return false;
    }
  }
  return true;

}
*/

var singleNumber = function(nums){
  let xor =0;
  for(let i = 0; i<nums.length; i++){
    xor = xor^nums[i];
  }
  return xor;
}

singleNumber([2,2,1])
