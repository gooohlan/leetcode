package WeeklyContest


type DataStream struct {
    // 优化一个版本,直接通过计数器计算,出现value一样的值就+1,否则清零
    k,value, cnt int
}


func Constructor2526(value int, k int) DataStream {
    return DataStream{
        // 优化一个版本,直接通过计数器计算,出现value一样的值就+1,否则清零
        k, value,0,
    }
}


func (d *DataStream) Consec(num int) bool {
   if num == d.value {
       d.cnt ++
   }else{
       d.cnt=0
   }
   return d.cnt >= d.k
}


/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */
