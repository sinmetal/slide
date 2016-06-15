var entries []Entry
var comments []Comment
var friends []Friend
var footprint []Footprint
errc := make(chan error)
go func() {
	_, err := datastore.NewQuery("Entry").GetAll(ctx, &entries)
	errc <- err
}()
go func() {
	_, err := datastore.NewQuery("Comment").GetAll(ctx, &comments)
	errc <- err
}()
go func() {
	_, err := datastore.NewQuery("Friend").GetAll(ctx, &friends)
	errc <- err
}()
go func() {
	_, err := datastore.NewQuery("Footprint").GetAll(ctx, &footprint)
	errc <- err
}()
err1, err2, err3, err4 := <-errc, <-errc, <-errc, <-errc