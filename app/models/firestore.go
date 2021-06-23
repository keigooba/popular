package models

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func init() {
	s, _ := ioutil.ReadFile("FirestoreAllKey.json")
	var jsondata map[string]interface{}
	json.Unmarshal([]byte(s), &jsondata)
	jsondata["private_key_id"] = os.Getenv("PrivateKeyId")
	//秘密鍵改行コード置き換え
	private_key := strings.Replace(os.Getenv("PrivateKey"), "\\n", "\n", -1)
	jsondata["private_key"] = private_key
	s, _ = json.Marshal(jsondata)
	file, err := os.Create("key.json")
	if err != nil {
		log.Println(err)
	}
	file.WriteString(string(s))
}

func ContactInsert(r *http.Request) error {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("key.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Println(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Println(err)
	}
	defer client.Close()

	// データ総数取得
	iter := client.Collection("ContactList").Documents(ctx)
	count := 1 //autoincrement分 +1しておく
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("繰り返しに失敗: %v", err)
		}
		count++
	}
	// 文字列変換
	str_count := strconv.Itoa(count)

	// その他場合、入力値を取得
	var other_browser, other_os string
	browser_id := r.FormValue("browser_id")
	if browser_id == "4" {
		other_browser = r.FormValue("other_browser")
	}
	os_id := r.FormValue("os_id")
	if os_id == "4" {
		other_os = r.FormValue("other_os")
	}

	contact := map[string]interface{}{
		"browser_id":    browser_id,
		"content":       r.FormValue("content"),
		"os_id":         os_id,
		"other_browser": other_browser,
		"other_os":      other_os,
		"twitter_name":  r.FormValue("twitter_name"),
	}

	// データ登録
	_, err = client.Collection("ContactList").Doc(str_count).Set(ctx, contact)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func ContactListGet() (contact_list []map[string]interface{}) {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("key.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Println(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Println(err)
	}
	defer client.Close()

	// すべてのデータを取得
	iter := client.Collection("ContactList").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("繰り返しに失敗: %v", err)
		}
		contact_list = append(contact_list, doc.Data())
	}

	return contact_list
}
