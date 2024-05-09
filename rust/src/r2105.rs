use crate::Solution;

impl Solution {
    pub fn minimum_refill(plants: Vec<i32>, capacity_a: i32, capacity_b: i32) -> i32 {
        let mut res = 0;
        let n = plants.len();
        let mut pos_a = 0; // A位置
        let mut pos_b = n - 1; // B位置
        let mut val_a = capacity_a; // A剩余水量
        let mut val_b = capacity_b; // B剩余水量

        while pos_a < pos_b {
            if val_a < plants[pos_a] { // A 水壶余量不够
                res += 1;
                val_a = capacity_a;
            }
            val_a -= plants[pos_a];
            pos_a += 1;

            if val_b < plants[pos_b] {
                res += 1;
                val_b = capacity_b;
            }
            val_b -= plants[pos_b];
            pos_b -= 1;
        }


        if pos_a == pos_b {
            if val_a >= val_b && val_a < plants[pos_a] {
                res += 1;
            }
            if val_b > val_a && val_b < plants[pos_b] {
                res += 1;
            }
        }
        res
    }
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test2105() {
        let plants = vec![2, 2, 3, 3];
        let capacity_a = 5;
        let capacity_b = 5;

        assert_eq!(Solution::minimum_refill(plants, capacity_a, capacity_b), 1);
    }
}


