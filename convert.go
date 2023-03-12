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

			//			fieldKind := OldKitval.Field(i).Kind() //string

			//fmt.Println("fieldName ", fieldName, "fieldTag", fieldTag, "fieldkind", fieldKind)

			//			valFROMFILE := reflect.Indirect(reflect.ValueOf((element.(map[string]interface{})[fieldTag])))
			//	fmt.Println("field tag", fieldTag, "valFROMFILE", valFROMFILE)
			//fmt.Println("valfromfile.kind()", valFROMFILE.Kind())
			if element.(map[string]interface{})[fieldTag] != nil {
				fmt.Println()

				fmt.Println("type of element by tag", reflect.TypeOf(element.(map[string]interface{})[fieldTag]))
				//fmt.Println("element.fieldtag", element.(map[string]interface{})[fieldTag])
				if fieldTag == "items" {

					// /var kititems := make([]OldKitItem,len)
					//We have to assert/cast this to be an array. Let's see if the length still works!
					//	items := make([]*OldKitItem, len(element.(map[string]interface{})[fieldTag].([]interface{})))
					//	fmt.Println(items)
					////	fmt.Println(element.(map[string]interface{})[fieldTag].([]interface{}))
					//fmt.Println(len(element.(map[string]interface{})[fieldTag].([]interface{})))
					//					kititems := make([]OldKitItem, len(element.(map[string]interface{})[fieldTag].([]interface{})))
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
						//item.(map[string]interface{})["itemid"]
						//Let's just do it the hard way?
						kititem.Mods = mods
						//	fmt.Println(kititem)
						kit.Items = append(kit.Items, kititem)
					}
					/*for i, row := range cr.Rows {
					s := &Showcase{}
					err := json.Unmarshal(*row[0], s)
					if err != nil {
						return items, err
					}


									ivp := &InvitedPerson{}
					err := json.Unmarshal(*row[0], ivp)*/

					//////////////////////OldKitval := reflect.Indirect(reflect.ValueOf(&kit)) ///this works!!!

					//////////////////////numItemField := OldKitval.NumField()

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

	jsonOutput := make(map[string]interface{})
	jsonOutput["Kits"] = kits

	file, err := json.MarshalIndent(jsonOutput, "", " ")
	if err != nil {
		panic(err)
	}
	_ = ioutil.WriteFile("test.json", file, 0644)

	//////////////////////////////
	//	fmt.Println(resultList2.(OldKit))

	/*
		fmt.Println("resultList2.(map[string]interface{})['akrat'].(map[string]interface{})")
		fmt.Println(resultList2.(map[string]interface{})["akrat"].(map[string]interface{}))
		fmt.Println("resultList2.(map[string]interface{})['akrat'].(map[string]interface{})['authlevel']")
		fmt.Println(resultList2.(map[string]interface{})["akrat"].(map[string]interface{})["authlevel"])
		//	log.Printf("INFO: resultList, %s", resultList)
		/*for _, element := range resultList {

		//		var kitThing OldKit
		/*
			fmt.Print("key type: ")
			fmt.Println(reflect.TypeOf(key))
			fmt.Print("Element type : ")
			fmt.Println(reflect.TypeOf(element)) */ /*

		fmt.Println(element.(OldKit))
		//fmt.Println(element.(map[string]interface{})["name"])
		fmt.Println(element.(map[string]interface{})["items"])

		for key2, element2 := range element.(map[string]interface{})["items"].([]interface{}) {
			fmt.Printf("key2 %d type: ", key2)
			fmt.Println(reflect.TypeOf(key2))
			fmt.Println("")
			fmt.Printf("Element2 %s type : ", element)
			fmt.Println(reflect.TypeOf(element2))
		}
		/*
			for key2, element2 := range element {
				fmt.Print("key: ")
				fmt.Println(reflect.TypeOf(key2))
				fmt.Print("Elment: ")
				fmt.Println(reflect.TypeOf(element2))
			}
			/*
				anotherElement := fmt.Sprintf("%s", element.(map[string]interface{}))
				byteElement := []byte(anotherElement)
				thirdElement := element.(map[string]interface{}) ///fmt.Sprintf("%s", element)
				fmt.Println(thirdElement)
				err := json.Unmarshal(byteElement, &kitThing)

				if err != nil {
					log.Printf("ERROR: %s fail to unmarshal nested JSON, %s", key, err.Error())
				}

	}*/

	if err != nil {
		log.Printf("ERROR: fail to unmarshla json, %s", err.Error())
	}
	//	log.Printf("INFO: jsonMap, %s", jsonMap)

	// defer the closing of our jsonFile so that we can parse it later on
	return
}

func oldmain() {
	/*
		csv_file, err := os.Open("Oldkits.csv")
		if err != nil {
			fmt.Println(err)
		}
		defer csv_file.Close()
		r := csv.NewReader(csv_file)
		records, err := r.ReadAll()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		output := fmt.Sprintf("%+v", records)
		fmt.Println(output)

		var Oldkititem OldKitItem
		var OldKitItems []OldKitItem
		initialName := "Survivor 01"
		for itemNumber, rec := range records {
			if rec[0] != "" {
				Oldkititem.KitName = rec[0]
				initialName = rec[0]
			} else {
				Oldkititem.KitName = initialName
			}
			//		Oldkititem.OldKitName = rec[0]
			itemID, _ := strconv.Atoi(rec[3])
			Oldkititem.ItemID = itemID
			if len(rec) > 5 {
				Oldkititem.Container = WhichContainer(rec[4], rec[5], rec[6])
				Oldkititem.Amount, _ = strconv.Atoi(rec[7])
				if Oldkititem.Amount == 0 {
					Oldkititem.Amount = 1
				}
			} else {
				Oldkititem.Container = "skipped"
			}
			//defaulting to no mods for now
			Oldkititem.Mods = []int{}

			OldKitItems = append(OldKitItems, Oldkititem)
			fmt.Println(itemNumber)
		}

		prevName := ""
		i := 0
		//j := 0

		//var OldKits []*OldKit
		tmap := make(map[string]interface{})
		ourLimit := len(OldKitItems)
		for i < len(OldKitItems) {
			Oldkit := &OldKit{}
			for OldKitItems[i].KitName == prevName {
				Oldkit.Name = OldKitItems[i].KitName
				Oldkit.Cooldown = 0.0
				Oldkit.Items = append(Oldkit.Items, OldKitItems[i])
				i++
				if i == ourLimit {
					tmap[Oldkit.Name] = Oldkit
					//OldKits = append(OldKits, Oldkit)
					break
				}
			}
			if i < ourLimit {
				fmt.Print("previous name was: " + prevName + "   ")
				prevName = OldKitItems[i].KitName
				fmt.Print("name is now: " + prevName)
				fmt.Println(i)
				//			OldKits = append(OldKits, Oldkit)
				tmap[Oldkit.Name] = Oldkit
			}

		}

		fmt.Println("Done")

		json_data, err := json.MarshalIndent(tmap, "", "  ")
		//json_data, err := json.Marshal(OldKits)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		output = fmt.Sprintf("%+s", json_data)
		fmt.Println(output)
		json_file, err := os.Create("sample.json")
		if err != nil {
			fmt.Println(err)
		}
		defer json_file.Close()

		json_file.Write(json_data)
		json_file.Close() */
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

func listStructFieldsAndTags(kit OldKit) {

	val := reflect.Indirect(reflect.ValueOf(kit))
	numField := val.NumField()
	//s := reflect.ValueOf(kit).FieldByIndex()
	for i := 0; i < numField; i++ {
		//		fmt.Println("a field maybe:", val.Type().Field(i).Name)
		fieldName := val.Type().Field(i).Name
		fmt.Println(fieldName)
		fieldTag := getStructTag(val.Type().Field(i), "json")

		fmt.Println(fieldTag)
		/*	field, ok := reflect.TypeOf(kit).Elem().FieldByName(fieldName)
			if !ok {
				panic("Field not found")
			}
			getStructTag(field, "json")
			//val.Type().Field(i)*/
	}
}

func getStructTag(f reflect.StructField, tagName string) string {

	return string(f.Tag.Get(tagName))
}
