package metalremind

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/file"
	"google.golang.org/appengine/image"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/image", handlerImage)
}

func handlerImage(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	dbn, err := file.DefaultBucketName(c)
	if err != nil {
		log.Errorf(c, "default bucket name get error : %v", err)
	}

	path := r.FormValue("path")
	if len(path) < 1 {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("path is required"))
		return
	}

	bk, err := blobstore.BlobKeyForFile(c, "/gs/"+dbn+"/"+path)
	if err != nil {
		log.Errorf(c, "blob key create error : %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	url, err := image.ServingURL(c, bk, nil)
	if err != nil {
		log.Errorf(c, "image servingURL error : %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(url.String()))
	return
}
