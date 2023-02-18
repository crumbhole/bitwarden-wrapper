package bitwardenwrapper

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

// GetItem finds a bitwarden item by name within a given bitwarden folder and returns it or nil and error if it cannot
func GetItem(itemName string, folderID BwFolderID) (*BwItem, error) {
	itemjson, err := runBw("list", "items")
	if err != nil {
		return nil, err
	}
	items := make([]BwItem, 0)
	json.Unmarshal([]byte(itemjson), &items)
	for _, item := range items {
		if item.Name == itemName && item.FolderID == folderID {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("Cannot find item called %s", itemName)
}

// GetItemFromFolder finds a bitwarden item by name within a named folder and returns it or nil and error if it cannot
func GetItemFromFolder(itemName string, folderName string) (*BwItem, error) {
	folder, err := GetFolder(folderName)
	if err != nil {
		return nil, err
	}
	item, err := GetItem(itemName, folder.ID)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func GetItemByPath(path string) (*BwItem, error) {
	folder, itemName := filepath.Split(path)
	return GetItemFromFolder(folder, itemName)
}
