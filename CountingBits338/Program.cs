/*public class Solution {
  public int[] CountBits(int n){
    int[] nums = new int[n+1];
    for(int i = 0; i<n+1;i++){

      string binary = Convert.ToString(i,2);
      int count = CountUnit(binary);
      nums[i]=count;
    }
    return nums;
  }
  public int CountUnit(string num){
    int count = 0;
    char[] charArray = num.ToCharArray();
      for(int i = 0; i < num.Length; i++){
        if (charArray[i] == '1'){
          count++;
        }

    }
      return count;
  }
}*/

public class Solution{
  public int[] CountBits(int n){
   int[] ans = new int[n+1];
   int exp = 1;
   int nextDigit =2;
   for(int i = 0; i <= n; i++){
    if(i > 1) ans[i] = ans[i-exp] + 1;
    else ans [i] = i;
    if(i+1 == nextDigit){
      exp *=2;
      nextDigit *=2;
    }
   }
   return ans;

  }

}
