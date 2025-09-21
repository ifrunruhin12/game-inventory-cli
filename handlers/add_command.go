package handlers

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ifrunruhin12/inventory/models"
	"github.com/ifrunruhin12/inventory/utils"
)

func (ch *CommandHandler) handleAdd(playerData *models.ReportData) error {
	// Check if inventory is full first
	if utils.IsInventoryFull(playerData.Inventory) {
		err := ch.tmplMgr.Execute(os.Stdout, "inventory_full.tmpl", nil)
		if err != nil {
			utils.Logger.Error("Failed to execute inventory full template", "error", err)
			return err
		}
		utils.Logger.Warn("Player tried to add item when inventory full", "player", playerData.Player.Name)
		return nil
	}

	// Prompt for item name
	fmt.Print("Name of the item you wanna add: ")
	itemName, err := ch.reader.ReadString('\n')
	if err != nil {
		utils.Logger.Error("Failed to read item name input", "error", err)
		return err
	}
	itemName = strings.TrimSpace(strings.ToLower(itemName))

	// Validate item exists in available items
	availableItems := GetAvailableItems()
	isValidItem := false
	for _, item := range availableItems {
		if strings.ToLower(item) == itemName {
			isValidItem = true
			break
		}
	}

	if !isValidItem {
		invalidItemData := struct {
			ItemName string
		}{
			ItemName: itemName,
		}
		err := ch.tmplMgr.Execute(os.Stdout, "invalid_item.tmpl", invalidItemData)
		if err != nil {
			utils.Logger.Error("Failed to execute invalid item template", "error", err)
			return err
		}
		utils.Logger.Warn("Player entered invalid item", "player", playerData.Player.Name, "item", itemName)
		return nil
	}

	// Check if player already has this item
	existingItemIndex := -1
	for i, item := range playerData.Inventory {
		if strings.ToLower(item.Name) == itemName {
			existingItemIndex = i
			break
		}
	}

	availableSlots := utils.GetAvailableSlots(playerData.Inventory)

	// If player already has the item, they can add any amount (no slot restriction)
	// If it's a new item, they need at least 1 slot
	if existingItemIndex == -1 && availableSlots == 0 {
		err := ch.tmplMgr.Execute(os.Stdout, "inventory_full.tmpl", nil)
		if err != nil {
			utils.Logger.Error("Failed to execute inventory full template", "error", err)
			return err
		}
		utils.Logger.Warn("Player tried to add new item when inventory full", "player", playerData.Player.Name, "item", itemName)
		return nil
	}

	// Quantity selection loop
	for {
		quantityPromptData := struct {
			ItemName       string
			AvailableSlots int
			IsNewItem      bool
		}{
			ItemName:       itemName,
			AvailableSlots: availableSlots,
			IsNewItem:      existingItemIndex == -1,
		}

		err := ch.tmplMgr.Execute(os.Stdout, "quantity_prompt.tmpl", quantityPromptData)
		if err != nil {
			utils.Logger.Error("Failed to execute quantity prompt template", "error", err)
			return err
		}

		quantityStr, err := ch.reader.ReadString('\n')
		if err != nil {
			utils.Logger.Error("Failed to read quantity input", "error", err)
			return err
		}
		quantityStr = strings.TrimSpace(quantityStr)

		quantity, err := strconv.Atoi(quantityStr)
		if err != nil || quantity <= 0 {
			invalidQuantityData := struct {
				Input string
			}{
				Input: quantityStr,
			}
			err := ch.tmplMgr.Execute(os.Stdout, "invalid_quantity.tmpl", invalidQuantityData)
			if err != nil {
				utils.Logger.Error("Failed to execute invalid quantity template", "error", err)
				return err
			}
			continue
		}

		// For new items, check if we have slots
		if existingItemIndex == -1 && availableSlots == 0 {
			err := ch.tmplMgr.Execute(os.Stdout, "no_slots_for_new_item.tmpl", nil)
			if err != nil {
				utils.Logger.Error("Failed to execute no slots template", "error", err)
				return err
			}
			continue
		}

		// Add the item to inventory
		if existingItemIndex != -1 {
			// Item exists, just increase count
			playerData.Inventory[existingItemIndex].Count += quantity
		} else {
			// New item, add to inventory
			newItem := models.Item{
				Name:  itemName,
				Count: quantity,
			}
			playerData.Inventory = append(playerData.Inventory, newItem)
		}

		// Success message
		successData := struct {
			Quantity int
			ItemName string
		}{
			Quantity: quantity,
			ItemName: itemName,
		}

		err = ch.tmplMgr.Execute(os.Stdout, "add_success.tmpl", successData)
		if err != nil {
			utils.Logger.Error("Failed to execute add success template", "error", err)
			return err
		}

		utils.Logger.Info("Player added item successfully",
			"player", playerData.Player.Name,
			"item", itemName,
			"quantity", quantity,
			"wasExisting", existingItemIndex != -1,
			"inventorySize", len(playerData.Inventory))

		break
	}

	return nil
}
