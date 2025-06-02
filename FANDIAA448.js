/*var findDisappearesNumber = function(nums){
  nums.sort();
  let arr = [];
  console.log(nums);
  for(let i=0 ;i<nums.length; i++){
    
      arr.push(i)
    
  }
  console.log(arr);
  return arr.filter(item=>!nums.includes(item));
};
*/
const findDisappearesNumber= function(nums){
let i =0;
  while (i <nums.length){
    let position = nums[i]-1
    if(nums[i]!=nums[position])
    {
      let swap = nums[i];
      nums[i] = nums[position]
      nums[position]=swap;

    } else{
      i+=1
    }
  }
  let miss =[]
  for(let i = 0; i<nums.length; i++){
    if (nums[i]!=i+1){
      miss.push(i+1)
    }
  }
  return miss

}
console.log(findDisappearesNumber([4,3,2,7,8,2,3,1]))
