package main

import (
    "encoding/json"
    "errors"
    "fmt"
)

type User struct {
    Uid       int    `json:"uid"`
    Headurl   string `json:"headurl"`
    ChipsAll  int    `json:"chips_all"`
    Country   string `json:"country"`
    ChipsList []struct {
        ActivityId int `json:"activity_id"`
        LeftChips  int `json:"left_chips"`
    } `json:"chips_list"`
}

func main() {
    str := `{
   "uid":123,
   "headurl": "http://baidu.com",
   "chips_all": 19999,
   "country": "VN",
   "chips_list":[
       {
       "activity_id":2,
       "left_chips":200
   },
       {
       "activity_id":4,
       "left_chips":400
   },
       {
       "activity_id":1,
       "left_chips":400
   }
   ]
}`
    var user User
    json.Unmarshal([]byte(str), &user)
    fmt.Println(user)
    
    Bet(0, 800, 1600, &user)
    fmt.Println(user)
    
}

// 定义一个函数来处理用户投注
func Bet(uid int, bet int, award int, user *User) error {
    // 假设你已经从数据库中获取到用户信息
    // 遍历用户的每种筹码，按顺序扣除
    profits := award - bet
    for _, chips := range user.ChipsList {
        if bet <= 0 {
            break
        }
        if chips.LeftChips > 0 {
            if chips.LeftChips >= bet {
                chips.LeftChips -= bet
                bet = 0
            } else {
                bet -= chips.LeftChips
                chips.LeftChips = 0
            }
        }
    }
    // 计算每种筹码的分配比例
    totalChips := 0
    for _, chips := range user.ChipsList {
        totalChips += chips.LeftChips
    }
    if totalChips == 0 {
        return errors.New("User has no chips left")
    }
    for i, chips := range user.ChipsList {
        if chips.LeftChips > 0 {
            ratio := float64(chips.LeftChips) / float64(totalChips)
            awardForChips := int(float64(profits) * ratio)
            // 将奖励分配到对应的活动账户中
            
            user.ChipsList[i].LeftChips += awardForChips
        }
    }
    return nil
}
