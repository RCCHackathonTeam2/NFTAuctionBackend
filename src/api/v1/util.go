package v1

const (
	CursorDelimiter = "_"
)

type chainIDMap map[int]string

type categorieIDMap map[int]string

var chainIDToChain = chainIDMap{
	1:        "eth",
	2:        "Solana",
	5:        "Ethereum Goerli",
	56:       "Binance Smart Chain",
	10:       "optimism",
	137:      "Polygon",
	11155111: "sepolia",
}

var categorieIDToCategorie = categorieIDMap{
	1: "艺术",
	2: "音乐",
	3: "摄影",
	4: "游戏资产",
	5: "收藏品",
	6: "域名",
	7: "体育",
	8: "虚拟世界",
}
