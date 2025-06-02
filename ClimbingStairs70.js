var climbStairs = function(n){
  let nums = 1;
  let numsItem = 1;
  let result;
  for(let i=0; i<n; i++){
    result = nums+numsItem;
    nums = numsItem;
    numsItem = result;
  }
  return reslut;
}
