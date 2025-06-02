var NumArray = function(nums){
  this.prifixSum = [0]
  for(let i = 0; i<nums.length; i++){
    this.prifixSum[i+1]= this.prifixSum[i]+nums[i]
  }
};

NumArray.prototype.sumRange = function(left, right){
  return this.prifixSum[right+1] - this.prifixSum[left];
}
