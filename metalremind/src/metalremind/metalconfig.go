package metalremind

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type MetalConfig struct {
	Id           string    `datastore:"-" goon:"id" json:"id"`        // metal-config-id 固定
	ClientId     string    `json:"clientId" datastore:",noindex"`     // GCP Client Id
	ClientSecret string    `json:"clientSecret" datastore:",noindex"` // GCP Client Secret
	SlackPostURL string    `json:"slackPostUrl" datastore:",noindex"` // Slackにぶっこむ用URL
	CreatedAt    time.Time `json:"createdAt"`                         // 作成日時
	UpdatedAt    time.Time `json:"updatedAt"`                         // 更新日時
}

const (
	metalConfigId = "metal-config-id"
)

func init() {
	http.HandleFunc("/admin/config", handlerMetalConfig)
}

func handlerMetalConfig(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	var mc MetalConfig
	err := json.NewDecoder(r.Body).Decode(&mc)
	if err != nil {
		log.Errorf(c, "request body json decode error : %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, err = datastore.Put(c, datastore.NewKey(c, "MetalConfig", metalConfigId, 0, nil), &mc)
	if err != nil {
		log.Warningf(c, "%v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func GetMetalConfig(c context.Context) (MetalConfig, error) {
	key := datastore.NewKey(c, "MetalConfig", metalConfigId, 0, nil)

	var mc MetalConfig
	err := datastore.Get(c, key, &mc)
	if err == datastore.ErrNoSuchEntity {
		return mc, err
	}
	if err != nil {
		return mc, err
	}
	return mc, nil
}

func (mc *MetalConfig) Save(c chan<- datastore.Property) error {
	now := time.Now()
	mc.UpdatedAt = now

	if mc.CreatedAt.IsZero() {
		mc.CreatedAt = now
	}

	if _, err := datastore.SaveStruct(mc); err != nil {
		return err
	}
	return nil
}
