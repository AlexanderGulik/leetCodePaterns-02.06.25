var maxProfit = function(prices){
  let minPrice = Infinity;
  let maxPrice = 0;
  for(let i = 0; i<prices.length; i++){
    if(prices[i]<minPrice){
    minPrice = prices[i]
    } else if(prices[i]-minPrice>maxPrice){
      maxPrice = prices[i] - minPrice;
    }

  }
  return maxPrice;
  
}
