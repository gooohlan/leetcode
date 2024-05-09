use crate::Solution;

impl Solution {
    pub fn abc(plants: Vec<i32>, capacity_a: i32, capacity_b: i32) -> i32 {
        let mut res = 0;
        let n = plants.len();
        let mut pos_a = 0; // A位置
        let mut pos_b = n - 1; // B位置
        let mut val_a = capacity_a; // A剩余水量
        let mut val_b = capacity_b; // B剩余水量

        while pos_a < pos_b {
            if val_a < plants[pos_a] { // A 水壶余量不够
                res += 1;
                val_a = capacity_a - plants[pos_a];
            } else {
                val_a -= plants[pos_a];
            }
            pos_a += 1;

            if val_b < plants[pos_b] {
                res += 1;
                val_b = capacity_b - plants[pos_b];
            } else {
                val_b -= plants[pos_b]
            }
            pos_b -= 1;
        }

        // Alice 和 Bob 到达同一株植物，那么当前水罐中水更多的人会给这株植物浇水
        if pos_a == pos_b && val_a.max(val_b) < plants[pos_a] {
            res += 1;
        }
        res
    }
}