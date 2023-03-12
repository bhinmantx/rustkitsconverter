package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
)

func convertOldtoNew(kits map[string]OldKit, output_filename string) (newKits map[string]interface{}) {
	newKits = make(map[string]interface{}, len(kits))
	////let's just do it the hard way!
	////The old kits were based on 3 location designated by a field.
	////We need to instead convert to 3 arrays for each location
	for _, old_kit := range kits {
		main_items := make([]KitItem, 0)
		wear_items := make([]KitItem, 0)
		belt_items := make([]KitItem, 0)

		var new_kit Kit

		new_kit.Name = old_kit.Name
		new_kit.Description = old_kit.Description
		new_kit.RequiredPermission = "" //KNOWN ISSUE! No good way currently to convert permissions
		new_kit.RequiredAuth = 0        //Same as above
		new_kit.MaximumUses = int64(old_kit.Max)
		new_kit.Cooldown = int64(old_kit.Cooldown) //Type conversion float to int for seconds.
		new_kit.Cost = 0                           //You'll need to adjust costs yourself
		new_kit.IsHidden = old_kit.Hide
		new_kit.CopyPasteFile = "" //this is new for the CopyPaste mod so it's always blank
		new_kit.KitImage = old_kit.Image

		//Now the nested stuff!
		for _, old_kit_item := range old_kit.Items {
			//SO! What do we have here?
			new_kit_item := KitItem{}
			//old kit system used numeric ID's and so we need to convert it to the short-name
			short_name, err := idToItem(old_kit_item.ItemID)
			if err != nil {
				//cheating since we're not doing "real" logging
				fmt.Println(fmt.Sprintf("Error finding ID for %s %s", old_kit_item.ItemID, err.Error()))
				continue
			}
			new_kit_item.Shortname = short_name

			new_kit_item.Skin = old_kit_item.SkinID
			new_kit_item.Amount = old_kit_item.Amount

			new_kit_item.Condition = 100.0
			new_kit_item.MaxCondition = 100.0

			///Does it need ammo? what's the default ammo? what's the default amount?

			needs_ammo, ammo_type, ammo_amount, err := needsAmmoWhatType(new_kit_item.Shortname)
			if err != nil {
				panic(err)
			}
			if needs_ammo {
				new_kit_item.Ammotype = ammo_type
				new_kit_item.Ammo = ammo_amount
			}

			//It is now necessary to check for kititems within kit items (mods for guns)
			//And we have to move things to new nested positions
			//new item types are used for "full" item types, i.e. tools and weapons
			if old_kit_item.Container == ContainerBelt {
				belt_items = append(belt_items, new_kit_item)
			} else if old_kit_item.Container == ContainerMainBP {
				main_items = append(main_items, new_kit_item)
			} else if old_kit_item.Container == ContainerWear {
				wear_items = append(wear_items, new_kit_item)
			}
		}
		new_kit.BeltItems = belt_items
		new_kit.WearItems = wear_items
		new_kit.MainItems = main_items
		newKits[new_kit.Name] = new_kit

	}

	jsonOutput := make(map[string]interface{})
	jsonOutput["_kits"] = newKits

	file, err := json.MarshalIndent(jsonOutput, "", " ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(output_filename, file, 0644)
	//We should *NOT* attempt to recover from failed conversion as a bad kit file
	// can have unexpected affects on Oxide mods
	if err != nil {
		panic(err)
	}

	return
}

func extractOldKits(oldkits_filename string) (kits map[string]OldKit) {

	filedata, err := ioutil.ReadFile(oldkits_filename)
	fmt.Println(fmt.Sprintf("%s", filedata))
	if err != nil {
		panic(err)
	}

	jsonMap := make(map[string](interface{}))

	err = json.Unmarshal(filedata, &jsonMap)

	if err != nil {
		panic(err)
	}

	kits = make(map[string]OldKit, 0)
	resultList := jsonMap["Kits"].(map[string]interface{})

	for key, element := range resultList {

		//Need to account for change in field names and positions
		var kit OldKit
		//Since we want to deal with arbitrary types
		OldKitval := reflect.Indirect(reflect.ValueOf(&kit)) ///

		numField := OldKitval.NumField()

		for i := 0; i < numField; i++ {

			fieldName := OldKitval.Type().Field(i).Name //"Description" or "Permission"

			fieldTag := getStructTag(OldKitval.Type().Field(i), "json") //json:"description"

			if element.(map[string]interface{})[fieldTag] != nil {
				///A bit messy but need to deal with arbitrary data types AND positions in the arrays
				if fieldTag == "items" {
					for _, item := range element.(map[string]interface{})[fieldTag].([]interface{}) {
						kititem := OldKitItem{}
						floatItemID := item.(map[string]interface{})["itemid"]
						floatString := fmt.Sprintf("%.0f", floatItemID)
						itemInt, err := strconv.Atoi(floatString)
						if err != nil {
							panic(err)
						}
						container := item.(map[string]interface{})["container"].(string)

						modlist := item.(map[string]interface{})["mods"]
						mods := make([]int, 0)
						for _, modF := range modlist.([]interface{}) {
							modString := fmt.Sprintf("%.0f", modF.(float64))
							mod, err := strconv.Atoi(modString)
							if err != nil {
								panic(err)
							}
							mods = append(mods, mod)
						}
						amountString := fmt.Sprintf("%.0f", item.(map[string]interface{})["amount"].(float64))
						amount, err := strconv.Atoi(amountString)
						if err != nil {
							panic(err)
						}
						skinidString := fmt.Sprintf("%.0f", item.(map[string]interface{})["skinid"].(float64))
						skinID, err := strconv.Atoi(skinidString)
						if err != nil {
							panic(err)
						}
						weapon := item.(map[string]interface{})["weapon"].(bool)
						blueprintTargetString := fmt.Sprintf("%.0f", item.(map[string]interface{})["blueprintTarget"].(float64))
						blueprintTarget, err := strconv.Atoi(blueprintTargetString)
						if err != nil {
							panic(err)
						}

						kititem.ItemID = itemInt
						kititem.Amount = amount
						kititem.SkinID = skinID
						kititem.Container = container
						kititem.Amount = amount
						kititem.BlueprintTarget = blueprintTarget
						kititem.Weapon = weapon

						//Going to need to deal with nested items!
						kititem.Mods = mods

						kit.Items = append(kit.Items, kititem)
					}
				} else {
					//was getting some comparison errors. Might as well compare string to string
					switch reflect.TypeOf(element.(map[string]interface{})[fieldTag]).String() {
					case "float64":
						OldKitval.FieldByName(fieldName).SetFloat(element.(map[string]interface{})[fieldTag].(float64))
					case "int64":
						OldKitval.FieldByName(fieldName).SetInt(element.(map[string]interface{})[fieldTag].(int64))
					case "string":
						OldKitval.FieldByName(fieldName).SetString(element.(map[string]interface{})[fieldTag].(string))
					case "bool":
						OldKitval.FieldByName(fieldName).SetBool(element.(map[string]interface{})[fieldTag].(bool))
					default:
						fmt.Println("default")
					}
				}
			}
		}

		kits[key] = kit
	}

	if err != nil {
		log.Printf("ERROR: fail to unmarshal json, %s", err.Error())
	}

	return
}

func WhichContainer(wearing string, belt string, backpack string) string {
	if wearing != "" {
		return "wear"
	}
	if belt != "" {
		return "belt"
	}
	return "main"

}

//For debugging
//Userful since the old kit fields can be of arbitrary types
func listStructFieldsAndTags(kit OldKit) {

	val := reflect.Indirect(reflect.ValueOf(kit))
	numField := val.NumField()

	for i := 0; i < numField; i++ {

		fieldName := val.Type().Field(i).Name
		fmt.Println(fieldName)
		fieldTag := getStructTag(val.Type().Field(i), "json")

		fmt.Println(fieldTag)

	}
}

func getStructTag(f reflect.StructField, tagName string) string {
	return string(f.Tag.Get(tagName))
}
