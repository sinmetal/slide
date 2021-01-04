q := datastore.NewQuery("Item").Order("-UpdatedAt").Limit(limit)
if len(cursorParam) > 0 {
	cursor, err := datastore.DecodeCursor(cursorParam)
	if err == nil {
		q = q.Start(cursor)
	}
}