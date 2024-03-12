package memory

import "inmemory/local/models"

type Base struct {
	items         []models.Users
	searchName    *SearchName
	searchAccount *SearchAccount
	searchValue   *SearchValue
}

type SearchName struct {
	items map[string]models.SearchName
}
type SearchAccount struct {
	items map[int]models.SearchAccount
}
type SearchValue struct {
	items map[float64]models.SearchValue
}

func (b *Base) Create(filter *models.Filter) error {
	b.items = append(b.items, models.Users{
		Account: filter.Account,
		Name:    filter.Name,
		Value:   filter.Value,
	})
	b.searchAccount.items[filter.Account] = models.SearchAccount{
		Name:  filter.Name,
		Value: filter.Value,
	}
	b.searchName.items[filter.Name] = models.SearchName{
		Account: filter.Account,
		Value:   filter.Value,
	}
	b.searchValue.items[filter.Value] = models.SearchValue{
		Name:    filter.Name,
		Account: filter.Account,
	}
	return nil
}

func (b *Base) Delete(account int) []models.Users {
	var res []models.Users
	for i, tmp := range b.items {
		if tmp.Account == account {
			res = append(res, tmp)
			b.items = append(b.items[:i], b.items[i+1:]...)
			return res
		}
	}
	return nil
}

func (b *Base) Update(filter *models.Filter, account int) []models.Users {
	var res []models.Users
	for i, tmp := range b.items {
		if tmp.Account == account {
			res = append(res, tmp)
			b.items = append(append(b.items[:i], models.Users(*filter)), b.items[i+1:]...)
			return res
		}
	}
	return nil
}

func (b *Base) List(filter *models.Filter) []models.Users {
	var res []models.Users
	if filter.Account != 0 {
		tmp := b.searchAccount.items[filter.Account]
		res = append(res, models.Users{
			Value:   tmp.Value,
			Account: filter.Account,
			Name:    tmp.Name,
		})
		return res
	}
	if filter.Name != "" {
		tmp := b.searchName.items[filter.Name]
		res = append(res, models.Users{
			Value:   tmp.Value,
			Account: tmp.Account,
			Name:    filter.Name,
		})
		return res
	}
	if filter.Value != 0 {
		tmp := b.searchValue.items[filter.Value]
		res = append(res, models.Users{
			Value:   filter.Value,
			Account: tmp.Account,
			Name:    tmp.Name,
		})
		return res
	}
	return b.items
}

func NewBase() *Base {
	tmp := []models.Users{
		{
			Account: 0,
			Name:    "Test",
			Value:   0.1,
		},
	}

	tmp2 := SearchName{
		items: map[string]models.SearchName{},
	}
	tmp2.items[tmp[0].Name] = models.SearchName{
		Account: tmp[0].Account,
		Value:   tmp[0].Value,
	}

	tmp3 := SearchValue{
		items: map[float64]models.SearchValue{},
	}
	tmp3.items[tmp[0].Value] = models.SearchValue{
		Account: tmp[0].Account,
		Name:    tmp[0].Name,
	}

	tmp4 := SearchAccount{
		items: map[int]models.SearchAccount{},
	}
	tmp4.items[tmp[0].Account] = models.SearchAccount{
		Name:  tmp[0].Name,
		Value: tmp[0].Value,
	}
	return &Base{
		items:         tmp,
		searchName:    &tmp2,
		searchValue:   &tmp3,
		searchAccount: &tmp4,
	}
}
