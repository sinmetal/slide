type Favorite struct {
	Id          string    `json:"id" datastore:"-"`
	Name        string    `json:"name"`
	Created     time.Time `json:"created"`
}

func (f *Favorite) save(c appengine.Context) (*Favorite, error) {
	f.Created = time.Now()
	k, err := datastore.Put(c, f.key(c), f)
	if err != nil {
		return nil, err
	}
	f.Id = k.StringID()
	return f, nil
}