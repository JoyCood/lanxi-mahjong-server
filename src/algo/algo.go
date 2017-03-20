/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2016-12-23 10:00
 * Filename      : algo.go
 * Description   : 玩牌算法
 * *******************************************************/
package algo

//import "fmt"

//// external function

// 打牌检测,胡牌, 放炮胡检测
func DiscardHu(card byte, cs []byte) uint32 {
	cards := make([]byte, len(cs)+1)
	copy(cards, cs)
	cards[len(cards)-1] = card
	//胡
	var status uint32 = existHu(cards)
	if status > 0 {
		//if status&HU_PING > 0 {
		//	status ^= HU_PING //去掉平胡
		//}
		status |= PAOHU
	}
	return status
}

//去掉平胡(有大牌型时去掉平胡番值,在牌局时才去掉)
func CancelHuPing(status uint32) uint32 {
	if status&HU_PING == 0 {
		return status
	}
	switch {
	case status&HU_SINGLE > 0:
		status ^= HU_PING //去掉平胡
	case status&HU_SINGLE_ZI > 0:
		status ^= HU_PING //去掉平胡
	case status&HU_SEVEN_PAIR_BIG > 0:
		status ^= HU_PING //去掉平胡
	case status&HU_SEVEN_PAIR > 0:
		status ^= HU_PING //去掉平胡
	case status&HU_SEVEN_PAIR_KONG > 0:
		status ^= HU_PING //去掉平胡
	case status&HU_ONE_SUIT > 0:
		status ^= HU_PING //去掉平胡
	case status&HU_ALL_ZI > 0:
		status ^= HU_PING //去掉平胡
	case status&TIAN_HU > 0:
		status ^= HU_PING //去掉平胡
	case status&DI_HU > 0:
		status ^= HU_PING //去掉平胡
	}
	return status
}

// 打牌检测,明杠/碰牌
func DiscardPong(card byte, cs []byte) uint32 {
	var status uint32
	//碰杠
	if existMingKong(card, cs) {
		status |= MING_KONG
		status |= KONG
	}
	if existPeng(card, cs) {
		status |= PENG
	}
	return status
}

// 打牌检测,吃
func DiscardChow(s1, s2 uint32, card byte, cs []byte) uint32 {
	var status uint32
	//吃
	if NextSeat(s1) == s2 {
		if existChow(card, cs) {
			status |= CHOW
		}
	}
	return status
}

// 摸牌检测,胡牌／暗杠／补杠
func DrawDetect(card byte, cs []byte, ch, ps, ks []uint32) uint32 {
	le := len(cs)
	cards := make([]byte, le)
	copy(cards, cs)
	var status uint32
	//自摸胡检测
	status = existHu(cards)
	if status > 0 {
		//if status&HU_PING > 0 {
		//	status ^= HU_PING //去掉平胡
		//}
		status |= ZIMO
	}
	if status > 0 {
		if le == 14 || menQing(ch, ps, ks) {
			status |= HU_MENQING
			if status&HU_PING > 0 {
				status ^= HU_PING //去掉平胡
			}
		}
	}
	if len(existAnKong(cards)) > 0 {
		status |= AN_KONG
		status |= KONG
	} else if existBuKong(card, ps) {
		status |= BU_KONG
		status |= KONG
	}
	return status
}

// 门清,(检测是否全暗杠或全没有返回true),TODO:优化,可以放在结算时判断
func menQing(cs, ps, ks []uint32) bool {
	if len(cs) > 0 {
		return false
	}
	if len(ps) > 0 {
		return false
	}
	if len(ks) == 0 {
		return true
	}
	for _, v2 := range ks { //杠
		_, _, v := DecodeKong(v2) //解码
		if v != AN_KONG { //不是暗杠
			return false
		}
	}
	return true
}

// 有暗杠时不算点炮(放冲,单钓)
func DanDiao(v uint32, ks []uint32) uint32 {
	if v&HU_DANDIAO == 0 {
		return v
	}
	if len(ks) == 0 {
		return v
	}
	for _, v2 := range ks { //杠
		_, _, v1 := DecodeKong(v2) //解码
		if v1 == AN_KONG { //有暗杠
			return v^HU_DANDIAO //去掉暗杠
		}
	}
	return v
}

// 胡牌牌型检测
func HuTypeDetect(hu, chow, kong bool, cs []byte) uint32 {
	Sort(cs, 0, len(cs)-1)
	return existHuType(hu, chow, kong, cs)
}

// 检测手牌是否有杠
func KongDetect(cs []byte) bool {
	var n uint32 = kongs(cs)
	return n > 0
}

// 验证吃, c1,c2,c3有序
func VerifyChow(c1, c2, c3 byte) bool {
	if c1 != c2 && c2 != c3 {
		if c1+1 == c2 && c2+1 == c3 {
			return true
		}
		if c1 >= 0x41 && c3 <= 0x44 {
			return true
		}
	}
	return false
}

// 碰杠吃数据
func EncodePeng(seat uint32, card byte) uint32 {
	seat = seat << 8
	seat |= uint32(card)
	return seat
}

func DecodePeng(value uint32) (seat uint32, card byte) {
	seat = value >> 8
	card = byte(value & 0xFF)
	return
}

func EncodeKong(seat uint32, card byte, value uint32) uint32 {
	value = value << 16
	value |= (seat << 8)
	value |= uint32(card)
	return value
}

func DecodeKong(value uint32) (seat uint32, card byte, v uint32) {
	v = value >> 16
	seat = (value >> 8) & 0xFF
	card = byte(value & 0xFF)
	return
}

func EncodeChow(c1, c2, c3 byte) (value uint32) {
	value =  uint32(c1) << 16
	value |= uint32(c2) << 8
	value |= uint32(c3)
	return
}

func DecodeChow(value uint32) (c1, c2, c3 byte) {
	c1 = byte(value >> 16)
	c2 = byte(value >> 8 & 0xFF)
	c3 = byte(value & 0xFF)
	return
}

func DecodeChow2(value uint32) (c1, c2 byte) {
	c1 = byte(value >> 8)
	c2 = byte(value & 0xFF)
	return
}

// 正常流程走牌令牌移到下一家
func NextSeat(seat uint32) uint32 {
	if seat == 4 {
		return 1
	}
	return seat + 1
}

// 是否存在n个牌
func Exist(c byte, cs []byte, n int) bool {
	for _, v := range cs {
		if n == 0 {
			return true
		}
		if c == v {
			n--
		}
	}
	return n == 0
}

// 移除一个牌
func Remove(c byte, cs []byte) []byte {
	for i, v := range cs {
		if c == v {
			cs = append(cs[:i], cs[i+1:]...)
			break
		}
	}
	return cs
}

// 移除n个牌,返回是否存在n个牌
func RemoveE(c byte, cs []byte, n int) ([]byte, bool) {
	var m int = n
	for n > 0 {
		for i, v := range cs {
			if c == v {
				cs = append(cs[:i], cs[i+1:]...)
				m--
				break
			}
		}
		n--
	}
	return cs, n == m
}

// 移除n个牌
func RemoveN(c byte, cs []byte, n int) []byte {
	for n > 0 {
		for i, v := range cs {
			if c == v {
				cs = append(cs[:i], cs[i+1:]...)
				break
			}
		}
		n--
	}
	return cs
}

// 对牌值从小到大排序，采用快速排序算法
func Sort(arr []byte, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			Sort(arr, start, j)
		}
		if end > i {
			Sort(arr, i, end)
		}
	}
}

//组合牌型番值(取最大值),TODO:优化
func huType2(v uint32) int32 {
	var f int32 = 0
	if v&HU_ONE_SUIT > 0 {
		f = max(f, huTypeOneSuit(v^HU_ONE_SUIT))
	}
	if v&HU_ALL_ZI > 0 {
		f = max(f, huTypeAllZi(v^HU_ALL_ZI))
	}
	if v&TIAN_HU > 0 {
		f = max(f, huTypeTian(v^TIAN_HU))
	}
	if v&DI_HU > 0 {
		f = max(f, huTypeDi(v^DI_HU))
	}
	return f
}

func huTypeOneSuit(v uint32) int32 {
	var f int32 = 0
	if v&HU_ONE_SUIT_PAIR_BIG > 0 {
		f = max(f, FAN[HU_ONE_SUIT_PAIR_BIG])
	}
	if v&HU_ONE_SUIT_PAIR > 0 {
		f = max(f, FAN[HU_ONE_SUIT_PAIR])
	}
	if v&HU_ONE_SUIT_PAIR_KONG > 0 {
		f = max(f, FAN[HU_ONE_SUIT_PAIR_KONG])
	}
	return f
}

func huTypeAllZi(v uint32) int32 {
	var f int32 = 0
	if v&HU_ALL_ZI_PAIR_BIG > 0 {
		f = max(f, FAN[HU_ALL_ZI_PAIR_BIG])
	}
	if v&HU_ALL_ZI_PAIR > 0 {
		f = max(f, FAN[HU_ALL_ZI_PAIR])
	}
	if v&HU_ALL_ZI_PAIR_KONG > 0 {
		f = max(f, FAN[HU_ALL_ZI_PAIR_KONG])
	}
	return f
}

func huTypeTian(v uint32) int32 {
	var f int32 = 0
	if v&TIAN_HU_ONE_SUIT > 0 {
		f = max(f, FAN[TIAN_HU_ONE_SUIT])
	}
	if v&TIAN_HU_ALL_ZI > 0 {
		f = max(f, FAN[TIAN_HU_ALL_ZI])
	}
	if v&TIAN_HU_SEVEN_PAIR > 0 {
		f = max(f, FAN[TIAN_HU_SEVEN_PAIR])
	}
	if v&TIAN_HU_SEVEN_PAIR_KONG > 0 {
		f = max(f, FAN[TIAN_HU_SEVEN_PAIR_KONG])
	}
	return f
}

func huTypeDi(v uint32) int32 {
	var f int32 = 0
	if v&DI_HU_ONE_SUIT > 0 {
		f = max(f, FAN[DI_HU_ONE_SUIT])
	}
	if v&DI_HU_ALL_ZI > 0 {
		f = max(f, FAN[DI_HU_ALL_ZI])
	}
	if v&DI_HU_SEVEN_PAIR > 0 {
		f = max(f, FAN[DI_HU_SEVEN_PAIR])
	}
	if v&DI_HU_SEVEN_PAIR_KONG > 0 {
		f = max(f, FAN[DI_HU_SEVEN_PAIR_KONG])
	}
	return f
}

func max(n, m int32) int32 {
	if n > m {
		return n
	}
	return m
}

//算番(牌型) TODO:优化
func HuType(v uint32, cs []byte) int32 {
	var f int32 = 0
	f = huType2(v)
	if f > 0 {
		return f //组合牌型时只取组合牌型
	}
	f = 1 //相乘时为1
	//var f int32 = 1 //相加时应该为0
	if v&HU_PING > 0 {
		f *= FAN[HU_PING]
	}
	if v&HU_SINGLE > 0 {
		f *= FAN[HU_SINGLE]
	}
	if v&HU_SINGLE_ZI > 0 {
		f *= FAN[HU_SINGLE_ZI]
	}
	if v&HU_SEVEN_PAIR_BIG > 0 {
		f *= FAN[HU_SEVEN_PAIR_BIG]
	}
	if v&HU_SEVEN_PAIR > 0 {
		f *= FAN[HU_SEVEN_PAIR]
	}
	if v&HU_SEVEN_PAIR_KONG > 0 {
		var n uint32 = kongs(cs)
		f *= max(16, FAN[HU_SEVEN_PAIR_KONG] * (1 << n)) //n=手牌中暗杠数:4*2^n,最大16番
	}
	if v&HU_ONE_SUIT > 0 {
		f *= FAN[HU_ONE_SUIT]
	}
	if v&HU_ALL_ZI > 0 { //没牌型4番,有8番
		if v == (HU|HU_ALL_ZI|ZIMO) ||
		v == (HU|HU_ALL_ZI|PAOHU) {
			f *= FAN[HU_ALL_ZI] / 2
		} else {
			f *= FAN[HU_ALL_ZI]
		}
	}
	//天地胡时取最大值
	if v&TIAN_HU > 0 {
		f = max(f, FAN[TIAN_HU])
	}
	if v&DI_HU > 0 {
		f = max(f, FAN[DI_HU])
	}
	return f
}

//算番(胡牌方式) TODO:优化
func HuWay(v uint32) int32 {
	var f int32 = 1
	if v&QIANG_GANG > 0 {
		f *= FAN[QIANG_GANG]
	}
	if v&HU_KONG_FLOWER > 0 {
		f *= FAN[HU_KONG_FLOWER]
	}
	if v&HU_MENQING > 0 {
		//必然是门清的显示上已经去掉,不用重复去除
		//if menQingFan(v) {
		//	f *= FAN[HU_MENQING]
		//}
		f *= FAN[HU_MENQING]
	}
	if v&HU_DANDIAO > 0 {
		f *= FAN[HU_DANDIAO]
	}
	//if v&TIAN_HU > 0 {
	//	f *= FAN[TIAN_HU]
	//}
	//if v&DI_HU > 0 {
	//	f *= FAN[DI_HU]
	//}
	return f
}

//算番(杠)
func HuKong(v uint32) int32 {
	return FAN[v]
}

//(小七对,十三烂, 天地胡)不算门清番
func menQingFan(v uint32) bool {
	if v&HU_SEVEN_PAIR > 0 {
		return false
	}
	if v&HU_SINGLE > 0 {
		return false
	}
	if v&HU_SINGLE_ZI > 0 {
		return false
	}
	if v&TIAN_HU > 0 {
		return false
	}
	if v&DI_HU > 0 {
		return false
	}
	return true
}

//(小七对,十三烂, 天地胡)不算门清番
func MenQingFan2(v uint32) uint32 {
	if v&HU_MENQING == 0 {
		return v
	}
	if v&HU_SEVEN_PAIR > 0 {
		return v^HU_MENQING
	}
	if v&HU_SINGLE > 0 {
		return v^HU_MENQING
	}
	if v&HU_SINGLE_ZI > 0 {
		return v^HU_MENQING
	}
	if v&TIAN_HU > 0 {
		return v^HU_MENQING
	}
	if v&DI_HU > 0 {
		return v^HU_MENQING
	}
	return v
}

//// internal function 

// 手牌有多少个杠牌
func kongs(cards []byte) uint32 {
	var i uint32 = 0
	var m = make(map[byte]int)
	for _, v := range cards {
		if n, ok := m[v]; ok {
			m[v] = n + 1
		} else {
			m[v] = 1
		}
	}
	for _, v := range m {
		if v == 4 {
			i += 1
		}
	}
	return i
}

//是否存在暗杠
func existAnKong(cards []byte) (kong []byte) {
	le := len(cards)
	for j := 0; j < le-3; j++ {
		count := 0
		for i := j + 1; i < le; i++ {
			if cards[j] == cards[i] {
				count = count + 1
				if count == 3 {
					kong = append(kong, cards[i])
					break
				}
			}
		}
	}
	return
}

//是否存在碰
func existPeng(card byte, cards []byte) bool {
	le := len(cards)
	count := 0
	for i := 0; i < le; i++ {
		if card == cards[i] {
			count = count + 1
			if count == 2 {
				return true
			}
		}
	}
	return false
}

//是否存在补杠
func existBuKong(card byte, pongs []uint32) bool {
	le := len(pongs)
	for i := 0; i < le; i++ {
		_, c := DecodePeng(pongs[i])
		if card == c {
			return true
		}
	}
	return false
}

//是否存在明杠
func existMingKong(card byte, cards []byte) bool {
	le := len(cards)
	count := 0
	for i := 0; i < le; i++ {
		if card == cards[i] {
			count = count + 1
			if count == 3 {
				return true
			}
		}
	}
	return false
}

//检测,是否存在吃 TODO:返回吃牌组合 [[1,2,3],[2,3,4],[3,4,5]]
func existChow(card byte, cs []byte) bool {
	le := len(cs)
	if le == 1 {
		return false
	}
	cards := make([]byte, le)
	copy(cards, cs)
	var t byte = card >> 4
	if uint32(t) >= FENG {
		return existChow1(t, card, cards)
	}
	return existChow2(card, cards)
}

//风牌,字牌两张不同即可
func existChow1(t, card byte, cs []byte) bool {
	count := 0
	var c byte
	for _, v := range cs {
		if v != card && v != c && v >> 4 == t {
			count += 1
			c = v
		}
	}
	if count >= 2 {
		return true
	}
	return false
}

//数牌三张
func existChow2(card byte, cs []byte) bool {
	var cards []byte = make([]byte, 5)
	cards[2] = card
	for _, v := range cs {
		var s int = int(card) - int(v) + 2
		if s >= 0 && s <= 4 {
			cards[s] = v
		}
	}
	count := 0
	for _, v := range cards {
		if count >= 3 {
			break
		}
		if v > 0 {
			count += 1
		} else {
			count = 0
		}
	}
	if count >= 3 {
		return true
	}
	return false
}

//有序slices,判断是否有4个相同的牌
func existFour2(cards []byte) bool {
	var c byte
	var i int
	for _, v := range cards {
		if c == v {
			i += 1
		} else {
			c = v
			i = 1
		}
		if i == 4 {
			return true
		}
	}
	return false
}

//有序slices,包含多少个杠
func haveKong(cards []byte) int {
	var c byte
	var i int
	var n int
	for _, v := range cards {
		if c == v {
			i += 1
		} else {
			c = v
			i = 1
		}
		if i == 4 {
			n += 1
		}
	}
	return n
}

// 胡牌后检测
// 清一色(同一花色,必须有牌型,没有字牌)/字一色(全部是字牌)
// 大七对, 1对子(将)+(刻子,杠) + 不能吃 + 杠(非手牌中杠)
// 有序slices cs 包含吃，碰，杠, chow=false没有吃,=true有吃
// kong=false没有杠,=true有杠, hu=false没有牌型,=true有牌型
func existHuType(hu, chow, kong bool, cs []byte) uint32 {
	var all_zi bool = true
	var one_suit bool = true
	var seven_pair_big bool = true
	var b bool
	var c byte
	var huType uint32
	var m = make(map[byte]int)
	for _, v := range cs {
		if n, ok := m[v]; ok {
			m[v] = n + 1
		} else {
			m[v] = 1
		}
		if !one_suit && !all_zi {
			continue
		}
		if c == 0 {
			c = v
		} else if c >> 4 != v >> 4 {
			one_suit = false
		} else if uint32(v >> 4) >= FENG {
			one_suit = false
		}
		if uint32(v >> 4) < FENG {
			all_zi = false
		}
	}
	for _, v := range m {
		if v == 2 && !b {
			b = true
			continue
		}
		if v < 3 || chow || kong {
			seven_pair_big = false
			break
		}
	}
	if seven_pair_big {
		huType |= HU
		huType |= HU_SEVEN_PAIR_BIG
	}
	// (huType > 0 || hu) //有牌型
	if one_suit && (huType > 0 || hu) {
		huType |= HU
		huType |= HU_ONE_SUIT
	}
	if all_zi {
		huType |= HU
		huType |= HU_ALL_ZI
	}
	return huType
}

// 13烂(数牌相隔2个或以上,字牌不重复)/七星13烂(13烂基础上包含7个不同字牌)
// 有序slices cs && len(cs) == 14
func existThirteen(cs []byte) uint32 {
	var thirteen_single bool = true
	var thirteen_single_zi bool = true
	var c byte
	var n int
	var huType uint32
	for _, v := range cs {
		if uint32(v >> 4) < FENG {
			if c == 0 || c >> 4 != v >> 4 {
				c = v
				continue
			}
			if v - c < 0x03 {
				thirteen_single = false
				break
			}
			c = v
		} else {
			n++
			if v == c {
				thirteen_single_zi = false
				break
			}
			c = v
		}
	}
	if thirteen_single_zi && thirteen_single_zi && n == 7 {
		huType |= HU_SINGLE_ZI
	} else if thirteen_single && thirteen_single_zi {
		huType |= HU_SINGLE
	}
	return huType
}

// 判断是否小七对(7个对子),豪华小七对(7个对子,其中有杠)
// 有序slices cs && len(cs) == 14
func exist7pair(cs []byte) uint32 {
	var seven_pair bool = true
	var c byte
	var i int
	var huType uint32
	var le int = len(cs)
	for n := 0; n < le-1; n += 2 {
		if cs[n] != cs[n+1] {
			seven_pair = false
			break
		}
		if i != 4 && cs[n] == c {
			i += 2
		} else if i != 4 {
			c = cs[n]
			i = 2
		}
	}
	if seven_pair && i == 4 {
		huType |= HU_SEVEN_PAIR_KONG
	} else if seven_pair {
		huType |= HU_SEVEN_PAIR
	}
	return huType
}

// 判断是否胡牌,0表示不胡牌,非0用32位表示不同的胡牌牌型
func existHu(cs []byte) uint32 {
	le := len(cs)
	//单钓胡牌
	if le == 2 && cs[0] == cs[1] {
		return HU | HU_DANDIAO
	}
	//排序slices
	Sort(cs, 0, le-1)
	// 14张牌型胡牌
	if le == 14 {
		// 七小对牌型胡牌
		t1 := exist7pair(cs)
		if t1 > 0 {
			return HU | t1
		}
		// 十三烂牌型胡牌
		t2 := existThirteen(cs)
		if t2 > 0 {
			return HU | t2
		}
	}
	//分开数牌和字牌(东南西北)
	cs_shu := make([]byte, 0)
	cs_zim := make(map[byte]int)
	for i := 0; i < le; i++ {
		if uint32(cs[i] >> 4) == FENG {
			cs_zim[cs[i]] += 1
		} else {
			cs_shu = append(cs_shu, cs[i])
		}
	}
	//TODO 优化
	cs_shu_l := len(cs_shu)
	if (cs_shu_l - 2) % 3 == 0 { //数牌做将
		if existHu3n2(cs_shu, cs_shu_l) { //是否3n+2牌型
			cs_zis := map2s(cs_zim)
			if existHu3n_zi(cs_zis, cs_zis) { //是否3n牌型
				return HU | HU_PING
			}
			return 0
		}
		return 0
	} else { //字牌做将
		if existHu3n(cs_shu, cs_shu_l) { //是否3n牌型
			cs_zis := map2s(cs_zim)
			if existHu3n2_zi(cs_zis) { //是否3n+2牌型
				return HU | HU_PING
			}
			return 0
		}
		return 0
	}
	return 0
}

func map2s(cs map[byte]int) []int {
	ms := make([]int, 0)
	for _, v := range cs {
		ms = append(ms, v)
	}
	return ms
}

// 3n +2 牌型胡牌 有序slices cs
func existHu3n2(cs []byte, le int) bool {
	list := make([]byte, le)
	for n := 0; n < le-1; n++ {
		if cs[n] == cs[n+1] { //
			copy(list, cs)
			list[n] = 0
			list[n+1] = 0
			for i := 0; i < le-2; i++ {
				if list[i] > 0 {
					for j := i + 1; j < le-1; j++ {
						if list[j] > 0 && list[i] > 0 {
							for k := j + 1; k < le; k++ {
								if list[k] > 0 && list[i] > 0 && list[j] > 0 {
									//刻子
									if list[i] == list[j] && list[j] == list[k] {
										list[i], list[j], list[k] = 0, 0, 0
										break
									}
									//顺子
									if list[i]+1 == list[j] && list[j]+1 == list[k] {
										list[i], list[j], list[k] = 0, 0, 0
										break
									}
								}
							}
						}
					}
				}
			}
			if existHu_(list, le) {
				return true
			}
		}
	}
	return false
}

// 有序slices cs, 3n 牌型胡牌
func existHu3n(cs []byte, le int) bool {
	for i := 0; i < le-2; i++ {
		if cs[i] > 0 {
			for j := i + 1; j < le-1; j++ {
				if cs[j] > 0 && cs[i] > 0 {
					for k := j + 1; k < le; k++ {
						if cs[k] > 0 && cs[i] > 0 && cs[j] > 0 {
							//刻子
							if cs[i] == cs[j] && cs[j] == cs[k] {
								cs[i], cs[j], cs[k] = 0, 0, 0
								break
							}
							//顺子
							if cs[i]+1 == cs[j] && cs[j]+1 == cs[k] {
								cs[i], cs[j], cs[k] = 0, 0, 0
								break
							}
						}
					}
				}
			}
		}
	}
	//是否可以胡
	if existHu_(cs, le) {
		return true
	}
	return false
}

//slice cs 有序, 是否可以胡
func existHu_(cs []byte, le int) bool {
	for i := 0; i < le; i++ {
		if cs[i] > 0 {
			return false
		}
	}
	return true
}

// 3n +2 牌型胡牌 [1,2,1,2]
func existHu3n2_zi(ms []int) bool {
	le := len(ms)
	for i, v := range ms {
		if v == 2 { //正常将
			list := make([]int, le-1)
			copy(list, append(ms[:i], ms[i+1:]...))
			if existHu3n_zi(list, list) { //是否3n牌型
				return true
			}
		} else if v > 2 { //刻子或杠折做将
			list := make([]int, le)
			copy(list, ms)
			list[i] -= 2
			if existHu3n_zi(list, list) { //是否3n牌型
				return true
			}
		}
	}
	return false
}

// 3n 牌型胡牌 [1,2,1,2]
func existHu3n_zi(ms, old []int) bool {
	le := len(ms)
	//判断是否组合成3n牌型
	if le == 0 {
		return true
	}
	if le < 3 { //少于3个时必须是刻子
		for i := 0; i < le; i++ {
			if ms[i] != 3 {
				return false //非刻子
			}
		}
		return true //是刻子
	}
	//可能的组合
	var ms_s [][]int
	if le == 3 { //存在3个时组合
		ms_s = [][]int{{1,1,1}}
	} else { //le == 4,存在4个时组合
		ms_s = [][]int{{1,1,1,0},{1,1,0,1},{1,0,1,1},{0,1,1,1}}
	}
	// [[1 1 1 0] [1 1 0 1] [1 0 1 1] [0 1 1 1]]
	// ms_s := zuheResult(le, 3)
	//遍历可能的组合
	for _, vs := range ms_s {
		//失败开始下种可能
		ms := make([]int, len(old))
		copy(ms, old)
		//fmt.Println("1 ms -> ", ms, " old -> ", old)
		//去除顺子值
		for k, v := range vs {
			if v == 0 {
				continue
			}
			if v == 1 {
				ms[k] -= 1
			}
		}
		//去除空值
		var n int = len(ms)
		for n > 0 {
			for i, v := range ms {
				if v == 0 {
					ms = append(ms[:i], ms[i+1:]...)
					break
				}
			}
			n--
		}
		//fmt.Println("2 ms -> ", ms, " old -> ", old)
		//
		ms_old := make([]int, len(ms))
		copy(ms_old, ms)
		//递归
		if existHu3n_zi(ms, ms_old) {
			return true
		}
	}
	return false
}

//组合算法(从nums中取出m个数)
func zuheResult(n int, m int) [][]int {
	if m < 1 || m > n {
		//fmt.Println("Illegal argument. Param m must between 1 and len(nums).")
		return [][]int{}
	}
	//保存最终结果的数组，总数直接通过数学公式计算
	result := make([][]int, 0, mathZuhe(n, m))
	//保存每一个组合的索引的数组，1表示选中，0表示未选中
	indexs := make([]int, n)
	for i := 0; i < n; i++ {
		if i < m {
			indexs[i] = 1
		} else {
			indexs[i] = 0
		}
	}
	//fmt.Println("indexs -> ", indexs)
	//第一个结果
	result = addTo(result, indexs)
	for {
		find := false
		//每次循环将第一次出现的 1 0 改为 0 1，同时将左侧的1移动到最左侧
		for i := 0; i < n-1; i++ {
			if indexs[i] == 1 && indexs[i+1] == 0 {
				find = true
				indexs[i], indexs[i+1] = 0, 1
				if i > 1 {
					moveOneToLeft(indexs[:i])
				}
				result = addTo(result, indexs)
				break
			}
		}
		//本次循环没有找到 1 0 ，说明已经取到了最后一种情况
		if !find {
			break
		}
	}
	return result
}

//将ele复制后添加到arr中，返回新的数组
func addTo(arr [][]int, ele []int) [][]int {
	newEle := make([]int, len(ele))
	copy(newEle, ele)
	arr = append(arr, newEle)
	return arr
}

func moveOneToLeft(leftNums []int) {
	//计算有几个1
	sum := 0
	for i := 0; i < len(leftNums); i++ {
		if leftNums[i] == 1 {
			sum++
		}
	}
	//将前sum个改为1，之后的改为0
	for i := 0; i < len(leftNums); i++ {
		if i < sum {
			leftNums[i] = 1
		} else {
			leftNums[i] = 0
		}
	}
}

//数学方法计算组合数(从n中取m个数)
func mathZuhe(n int, m int) int {
	return factorial(n) / (factorial(n-m) * factorial(m))
}

//阶乘
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
