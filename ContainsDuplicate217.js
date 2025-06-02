/*const containsDuplicate = function(nums){
  for(let i =0; i<nums.length; i++){
    for (let j =1; j<nums.length; j++){
      console.log(nums[i]);
      console.log(nums[j]);
      if(nums[i]===nums[j]&&i!==j) return true;
      
    }
  }
  return false;
}*/

const containsDuplicate = function(nums){
let seen = new Set();
  for(const item of nums){
    if(seen.has(item)){return true;}
  seen.add(item);
  }
  return false;
}



console.log(containsDuplicate([1,2,3,4]))
