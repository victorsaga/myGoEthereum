package CacheHelper

// 因為Cache.Set沒辦法存陣列，所以加了一層Wrap
// type GetThirdpartyTradingPairsWrap struct {
// 	Data []RepositoryModel.Trading_pair_mapping
// }

// func GetThirdpartyTradingPairs() (result []HedgingRepositoryModel.Trading_pair_mapping) {
// 	key := "GetThirdpartyTradingPairs"
// 	if data, found := MemoryCacheHelper.GetCache(key); found {
// 		result = data.(GetThirdpartyTradingPairsWrap).Data
// 		return
// 	}

// 	data := Repository.GetEnableThirdpartyTradingPairs()
// 	MemoryCacheHelper.SetCache(key, GetThirdpartyTradingPairsWrap{Data: data}, time.Minute)
// 	return data
// }
