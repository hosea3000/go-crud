package crud

type Service[T any] interface {
	Fetch() []*T
	FetchOne(string) (*T, error)
	Create(interface{}) error
	Update(*T) error
	Delete(string) error
}

var _ Service[any] = (*BaseService[any])(nil)

type BaseService[T any] struct {
}

func (s *BaseService[T]) Fetch() []*T {
	books := make([]*T, 0)
	DB.Find(&books)
	return books
}

func (s *BaseService[T]) FetchOne(id string) (*T, error) {
	t := new(T)
	result := DB.First(t, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return t, nil
}

func (s *BaseService[T]) Create(m interface{}) error {
	result := DB.Create(m)
	//user.ID             // 返回插入数据的主键
	//result.Error        // 返回 error
	//result.RowsAffected // 返回插入记录的条数
	return result.Error
}

func (s *BaseService[T]) Update(t *T) error {
	result := DB.Save(&t)
	return result.Error
}

func (s *BaseService[T]) Delete(id string) error {
	result := DB.Delete(new(T), id)

	return result.Error
}
