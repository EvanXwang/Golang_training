package main
import "fmt"

/*
你有50個金幣，需要分配給以下幾個人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth
分配規則如下：
1. 名字中每包含1個 'e' 或 'E' 分1枚金幣
2. 名字中每包含1個 'i' 或 'I' 分1枚金幣
3. 名字中每包含1個 'o' 或 'O' 分1枚金幣
4. 名字中每包含1個 'u' 或 'U' 分1枚金幣

寫一個功能，計算每個用戶可分到多少金幣，以及最後剩多少金幣
程式結構如下，請實現 'dispatchCoin' func
 */

var (
	coins int
	users = []string{
		"Matthew","Sarah","Augustus","Heidi","Emilie","Peter","Giana","Adriano","Aaron","Elizabeth",
	}
	distribution = make(map[string]int,len(users))
)



func dispatchCoin() (left int) {
	// 1.依次拿到每個人的名字
	for _, name:= range users{
		// 2.拿到一個人名 依照規則分金幣
		for _, c:= range name{
			switch c{
			case 'e','E':
				//滿足該條件分1枚金幣
				distribution[name] ++
				coins--
			case 'i','I':
				//滿足該條件分2枚金幣
				distribution[name] +=2
				coins-=2
			case 'o','O':
				//滿足該條件分3枚金幣
				distribution[name] +=3
				coins-=3
			case 'u','U':
				//滿足該條件分4枚金幣
				distribution[name] +=4
				coins-=4
			}
		}
	}
	left = coins

	// 3.每個人分的金幣數應該保存到distribution
	// 4.還要記錄下剩餘的金幣數
	// 5.跑完3、4步 就能得到最終每個人分的金幣數和剩餘金幣數
	return
}

func main(){

	fmt.Println("請輸入金幣：")
	fmt.Scan(&coins)

	left := dispatchCoin()
	fmt.Println("剩下：",left,"枚")
	fmt.Println("每個人可分得金幣數量如下：")
	for k,v := range distribution{
		fmt.Printf("%s:%d\n",k,v)
	}
}