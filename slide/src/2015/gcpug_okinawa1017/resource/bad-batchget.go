for key := range keys {
	err := datastore.Get(ctx, key, dst)
}