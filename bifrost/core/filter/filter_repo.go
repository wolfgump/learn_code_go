package filter

type FilterRepo struct {
	filterContainer map[string]IFilter
}

var filterRepo *FilterRepo

func FilterRepoInstance() *FilterRepo {
	if filterRepo == nil {
		filterRepo = &FilterRepo{}
		filterRepo.filterContainer = make(map[string]IFilter)
	}
	return filterRepo
}

func (fr *FilterRepo) Register(filterName string, filter IFilter) {
	fr.filterContainer[filterName] = filter
}
func RegisterStaticFilter() {
	register := FilterRepoInstance()
	register.Register("PrefixStrip", &PrefixStripFilter{Filter{}})
	register.Register("Base", &Filter{})
	register.Register("LoadBalance", &LoadBalanceFilter{Filter{}})
	register.Register("RequestBackend",&RequestBackendFilter{Filter{}})

}
func AllFilters() []IFilter {
	filterMap := FilterRepoInstance().filterContainer
	filters := make([]IFilter, len(filterMap))
	cnt := 0
	for _, value := range filterMap {
		filters[cnt] = value
		cnt++
	}
	return filters
}
func GetFilterByName(filterName string) IFilter {
	filterMap := FilterRepoInstance().filterContainer
	return filterMap[filterName]
}
