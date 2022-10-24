package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Text(t *testing.T) {
	k := "作曲 : 姜云升\n作词 : 姜云升\n编曲 : 朴冉\n混音 : 朴冉\n录音工程师 : 王润天\n制作人 : 朴冉\n你今夜为了什么带上耳机\n直到路灯全部熄灭还没睡去\n对失去感觉不舍还是可惜\n期待着明天 又藏着几分畏惧\n你说你已经慢慢学会狠心\n尤其是把感情看的很轻\n你说你只有这样才能快乐\n那你回答我你此刻为何红了眼睛\n过去那些伤口应该怎么样去治理\n现在似乎只有钞票能够带来刺激\n疲于人际关系 被卷入那些不相干的事里\n就关闭卧室房门以后你才是你自己\n孤独是种状态 即便多么勤奋\n出现在生活中你眼神某一刻的停顿\n是谁在白天笑得开心 为人多么平顺\n又是谁在网易云看了一整夜的评论\n总说着别为我担心 我们装得顶天立地绝对不会退缩\n我们总是说着别为我担心 可是今天为何泪水却打湿了被窝\n那些是非对错你难以分清\n变得成熟也不再真心\n被子把头捂住 喉咙像被堵住\n成年人的哭原来没有声音\n不得不去告别 这结果早就认得\n你知道她不爱你 只是不敢去承认吧\n没必要去抱怨 也不用去讨论了\n她都说了抱歉 那还有什么好问的\n想要把电话线给 都掐断了\n又是谁的心脏 碎成稀巴烂了\n第二天的清晨 不留一点痕迹\n除了垃圾桶里的 酒瓶和易拉罐子\n你是为了什么总听一首歌\n为了什么歇斯底里 为了什么不想说话 为了什么总听一首歌\n为了什么让不喜欢喝酒的你 把酒当水喝\n这微信把你我变成游子 浮躁且容易厌倦\n沟通只需动动你的手指 连分手都不用见面\n谁在私心戒备 把借口挑明\n谁痴心绝对 被烈酒扫平\n眷眷之心变味 把怨毒表明\n明明撕心裂肺 你却又面无表情\n希望可以等到那个她的来到\n希望每次伸手都能得到怀抱\n希望找到一种方式可以让你解脱\n希望把过去所有不幸全部埋掉\n你今夜为了什么带上耳机\n直到路灯全部熄灭还没睡去\n对失去感觉不舍还是可惜\n期待着明天 又藏着几分畏惧\n你说你已经慢慢学会狠心\n尤其是把感情看得很轻\n你说你只有这样才能快乐\n那你回答我你此刻为何红了眼睛\n你今夜为了什么带上耳机\n即便过去糟糕到不敢去回看\n的确感觉不舍还有可惜\n但我相信未来有你陪伴\n我也尝试学会所谓狠心\n却不想让我自己 变得冷冰\n愿每一个人都能得到快乐\n愿你们今夜可以不再红着眼睛"
	tt := NewText(k)
	tt.Participle()
	tt.CountWord()
	fmt.Println(len(tt.WordCountList[2][1].(string)))
	fmt.Println(len(tt.WordCountList[3][1].(string)))
	fmt.Println([]byte(tt.WordCountList[2][1].(string)))
	fmt.Println([]byte(tt.WordCountList[3][1].(string)))
	fmt.Println(strings.Compare(tt.WordCountList[2][1].(string), tt.WordCountList[3][1].(string)))
	fmt.Println(strings.EqualFold(tt.WordCountList[2][1].(string), tt.WordCountList[3][1].(string)))
	fmt.Println(len(tt.WordCountList[7][1].(string)))
	fmt.Println(len(tt.WordCountList[8][1].(string)))
	fmt.Println(len(tt.WordMap))
	fmt.Println(len(tt.WordCountList))
	fmt.Println(tt.WordCountList)
	s1, s2 := 0, 0
	for _, v := range tt.WordCountList {
		s1 += v[0].(int)
	}
	for _, v := range tt.WordMap {
		s2 += v
	}
	fmt.Println(s1, s2)
}
