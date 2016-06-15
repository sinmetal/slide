var entries []Entry
var comments []Comment
var friends []Friend
var footprint []Footprint
_, err := datastore.NewQuery("Entry").GetAll(ctx, &entries)
_, err = datastore.NewQuery("Comment").GetAll(ctx, &comments)
_, err = datastore.NewQuery("Friend").GetAll(ctx, &friends)
_, err = datastore.NewQuery("Footprint").GetAll(ctx, &footprint)