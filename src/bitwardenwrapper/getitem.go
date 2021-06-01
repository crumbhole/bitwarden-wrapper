package bitwardenwrapper

import (
	"encoding/json"
	"fmt"
)

func GetItem(itemName string, folderId BwFolderId) (*BwItem, error) {
	itemjson, err := runBw("list", "items")
	if err != nil {
		return nil, err
	}
	items := make([]BwItem, 0)
	json.Unmarshal([]byte(itemjson), &items)
	for _, item := range items {
		if item.Name == itemName && item.FolderId == folderId {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("Cannot find item called %s", itemName)
}

func GetItemFromFolder(itemName string, folderName string) (*BwItem, error) {
	folder, err := GetFolder(folderName)
	if err != nil {
		return nil, err
	}
	item, err := GetItem(itemName, folder.Id)
	if err != nil {
		return nil, err
	}
	return item, nil
}
