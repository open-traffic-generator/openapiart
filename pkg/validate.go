package openapiart

import "fmt"

func (obj *prefixConfig) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"prefixConfig"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

	// RequiredObject is required
	if obj.obj.RequiredObject == nil {
		obj.errors = append(obj.errors, "RequiredObject is required field on interface PrefixConfig")
	}

	if obj.obj.RequiredObject != nil {

		obj.RequiredObject().validateObj(setDefault)
	}

	if obj.obj.OptionalObject != nil {

		obj.OptionalObject().validateObj(setDefault)
	}

	if obj.obj.FullDuplex_100Mb != nil {
		if *obj.obj.FullDuplex_100Mb < 0 || *obj.obj.FullDuplex_100Mb > 4261412864 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PrefixConfig.FullDuplex_100Mb <= 4261412864 but Got %d", *obj.obj.FullDuplex_100Mb))
		}

	}

	// A is required
	if obj.obj.A == "" {
		obj.errors = append(obj.errors, "A is required field on interface PrefixConfig")
	}

	// A is underReview
	if obj.obj.A != "" {
		obj.underReview(`A: property under review`)
	}

	// B is deprecated
	if obj.obj.B != 0 {
		obj.deprecated(`B: Property b is being deprecated from the sdk version x.x.x
and property x shall be used instead`)
	}

	// DValues is deprecated
	if obj.obj.DValues != nil {
		obj.deprecated(`DValues: Property d_values is being deprecated from the sdk version x.x.x
and property d shall be used instead`)
	}

	if obj.obj.E != nil {
		obj.deprecated(`E: Property 'e' is being deprecated from the sdk version x.x.x
and property 'ep' shall be used instead`)
		obj.E().validateObj(setDefault)
	}

	if obj.obj.F != nil {

		obj.F().validateObj(setDefault)
	}

	if len(obj.obj.G) != 0 {

		if setDefault {
			obj.G().clearHolderSlice()
			for _, item := range obj.obj.G {
				newObj := newGObject(obj.validator)
				newObj.self().obj = item
				obj.G().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.G().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.J) != 0 {

		if setDefault {
			obj.J().clearHolderSlice()
			for _, item := range obj.obj.J {
				newObj := newJObject(obj.validator)
				newObj.self().obj = item
				obj.J().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.J().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.K != nil {

		obj.K().validateObj(setDefault)
	}

	if obj.obj.L != nil {

		obj.L().validateObj(setDefault)
	}

	if obj.obj.Level != nil {

		obj.Level().validateObj(setDefault)
	}

	if obj.obj.Mandatory != nil {

		obj.Mandatory().validateObj(setDefault)
	}

	if obj.obj.Ipv4Pattern != nil {

		obj.Ipv4Pattern().validateObj(setDefault)
	}

	if obj.obj.Ipv6Pattern != nil {

		obj.Ipv6Pattern().validateObj(setDefault)
	}

	if obj.obj.MacPattern != nil {

		obj.MacPattern().validateObj(setDefault)
	}

	if obj.obj.IntegerPattern != nil {

		obj.IntegerPattern().validateObj(setDefault)
	}

	if obj.obj.ChecksumPattern != nil {

		obj.ChecksumPattern().validateObj(setDefault)
	}

	if obj.obj.Case != nil {

		obj.Case().validateObj(setDefault)
	}

	if obj.obj.MObject != nil {

		obj.MObject().validateObj(setDefault)
	}

	if obj.obj.Integer64List != nil {

		for _, item := range obj.obj.Integer64List {
			if item < 0 || item > 4261412864 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PrefixConfig.Integer64List <= 4261412864 but Got %d", item))
			}

		}

	}

	if obj.obj.HeaderChecksum != nil {

		obj.HeaderChecksum().validateObj(setDefault)
	}

	// StrLen is underReview
	if obj.obj.StrLen != nil {
		obj.underReview(`StrLen: property under review`)
	}

	if obj.obj.StrLen != nil {
		if len(*obj.obj.StrLen) < 3 || len(*obj.obj.StrLen) > 6 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf(
					"3 <= length of PrefixConfig.StrLen <= 6 but Got %d",
					len(*obj.obj.StrLen)))
		}

	}

	// HexSlice is underReview
	if obj.obj.HexSlice != nil {
		obj.underReview(`HexSlice: property under review`)
	}

	if obj.obj.HexSlice != nil {

		err := obj.validateHexSlice(obj.HexSlice())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PrefixConfig.HexSlice"))
		}

	}

	if obj.obj.AutoFieldTest != nil {

		obj.AutoFieldTest().validateObj(setDefault)
	}

	if len(obj.obj.WList) != 0 {

		if setDefault {
			obj.WList().clearHolderSlice()
			for _, item := range obj.obj.WList {
				newObj := newWObject(obj.validator)
				newObj.self().obj = item
				obj.WList().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.WList().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.XList) != 0 {

		if setDefault {
			obj.XList().clearHolderSlice()
			for _, item := range obj.obj.XList {
				newObj := newZObject(obj.validator)
				newObj.self().obj = item
				obj.XList().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.XList().Items() {
			item.validateObj(setDefault)
		}

	}

	if obj.obj.ZObject != nil {

		obj.ZObject().validateObj(setDefault)
	}

	if obj.obj.YObject != nil {

		obj.YObject().validateObj(setDefault)
	}

	if len(obj.obj.SubList) != 0 {

		if setDefault {
			obj.SubList().clearHolderSlice()
			for _, item := range obj.obj.SubList {
				newObj := newSubObject(obj.validator)
				newObj.self().obj = item
				obj.SubList().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.SubList().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *prefixConfig) checkUnique() {
	if !obj.isUnique("prefixConfig", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *prefixConfig) checkConstraint() {

	vobjectCons := []string{
		"wObject.WName",
	}

	for _, v := range obj.VObject() {
		if !obj.validateConstraint(vobjectCons, v) {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("%s is not a valid wObject.WName type", v),
			)
		}
	}

}

func (obj *updateConfig) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.G) != 0 {

		if setDefault {
			obj.G().clearHolderSlice()
			for _, item := range obj.obj.G {
				newObj := newGObject(obj.validator)
				newObj.self().obj = item
				obj.G().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.G().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *updateConfig) checkUnique() {

}

func (obj *updateConfig) checkConstraint() {

}

func (obj *metricsRequest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *metricsRequest) checkUnique() {

}

func (obj *metricsRequest) checkConstraint() {

}

func (obj *apiTestInputBody) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *apiTestInputBody) checkUnique() {

}

func (obj *apiTestInputBody) checkConstraint() {

}

func (obj *setConfigResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_404 != nil {

		obj.StatusCode404().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *setConfigResponse) checkUnique() {

}

func (obj *setConfigResponse) checkConstraint() {

}

func (obj *updateConfigurationResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *updateConfigurationResponse) checkUnique() {

}

func (obj *updateConfigurationResponse) checkConstraint() {

}

func (obj *getConfigResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *getConfigResponse) checkUnique() {

}

func (obj *getConfigResponse) checkConstraint() {

}

func (obj *getMetricsResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *getMetricsResponse) checkUnique() {

}

func (obj *getMetricsResponse) checkConstraint() {

}

func (obj *getWarningsResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *getWarningsResponse) checkUnique() {

}

func (obj *getWarningsResponse) checkConstraint() {

}

func (obj *clearWarningsResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *clearWarningsResponse) checkUnique() {

}

func (obj *clearWarningsResponse) checkConstraint() {

}

func (obj *getRootResponseResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

}

func (obj *getRootResponseResponse) checkUnique() {

}

func (obj *getRootResponseResponse) checkConstraint() {

}

func (obj *dummyResponseTestResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

}

func (obj *dummyResponseTestResponse) checkUnique() {

}

func (obj *dummyResponseTestResponse) checkConstraint() {

}

func (obj *postRootResponseResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

}

func (obj *postRootResponseResponse) checkUnique() {

}

func (obj *postRootResponseResponse) checkConstraint() {

}

func (obj *getAllItemsResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

}

func (obj *getAllItemsResponse) checkUnique() {

}

func (obj *getAllItemsResponse) checkConstraint() {

}

func (obj *getSingleItemResponse) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

}

func (obj *getSingleItemResponse) checkUnique() {

}

func (obj *getSingleItemResponse) checkConstraint() {

}

func (obj *getSingleItemLevel2Response) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.StatusCode_400 != nil {

		obj.StatusCode400().validateObj(setDefault)
	}

	if obj.obj.StatusCode_500 != nil {

		obj.StatusCode500().validateObj(setDefault)
	}

	if obj.obj.StatusCode_200 != nil {

		obj.StatusCode200().validateObj(setDefault)
	}

}

func (obj *getSingleItemLevel2Response) checkUnique() {

}

func (obj *getSingleItemLevel2Response) checkConstraint() {

}

func (obj *eObject) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

}

func (obj *eObject) checkUnique() {
	if !obj.isUnique("eObject", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *eObject) checkConstraint() {

}

func (obj *fObject) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *fObject) checkUnique() {

}

func (obj *fObject) checkConstraint() {

}

func (obj *gObject) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *gObject) checkUnique() {

}

func (obj *gObject) checkConstraint() {

}

func (obj *jObject) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.JA != nil {

		obj.JA().validateObj(setDefault)
	}

	if obj.obj.JB != nil {

		obj.JB().validateObj(setDefault)
	}

}

func (obj *jObject) checkUnique() {

}

func (obj *jObject) checkConstraint() {

}

func (obj *kObject) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.EObject != nil {

		obj.EObject().validateObj(setDefault)
	}

	if obj.obj.FObject != nil {

		obj.FObject().validateObj(setDefault)
	}

}

func (obj *kObject) checkUnique() {

}

func (obj *kObject) checkConstraint() {

}

func (obj *lObject) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Integer != nil {
		if *obj.obj.Integer < 10 || *obj.obj.Integer > 90 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("10 <= LObject.Integer <= 90 but Got %d", *obj.obj.Integer))
		}

	}

	if obj.obj.Mac != nil {

		err := obj.validateMac(obj.Mac())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on LObject.Mac"))
		}

	}

	if obj.obj.Ipv4 != nil {

		err := obj.validateIpv4(obj.Ipv4())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on LObject.Ipv4"))
		}

	}

	if obj.obj.Ipv6 != nil {

		err := obj.validateIpv6(obj.Ipv6())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on LObject.Ipv6"))
		}

	}

	if obj.obj.Hex != nil {

		err := obj.validateHex(obj.Hex())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on LObject.Hex"))
		}

	}

}

func (obj *lObject) checkUnique() {

}

func (obj *lObject) checkConstraint() {

}

func (obj *levelOne) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.L1P1 != nil {

		obj.L1P1().validateObj(setDefault)
	}

	if obj.obj.L1P2 != nil {

		obj.L1P2().validateObj(setDefault)
	}

}

func (obj *levelOne) checkUnique() {

}

func (obj *levelOne) checkConstraint() {

}

func (obj *mandate) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// RequiredParam is required
	if obj.obj.RequiredParam == "" {
		obj.errors = append(obj.errors, "RequiredParam is required field on interface Mandate")
	}
}

func (obj *mandate) checkUnique() {

}

func (obj *mandate) checkConstraint() {

}

func (obj *ipv4Pattern) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Ipv4 != nil {

		obj.Ipv4().validateObj(setDefault)
	}

}

func (obj *ipv4Pattern) checkUnique() {

}

func (obj *ipv4Pattern) checkConstraint() {

}

func (obj *ipv6Pattern) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Ipv6 != nil {

		obj.Ipv6().validateObj(setDefault)
	}

}

func (obj *ipv6Pattern) checkUnique() {

}

func (obj *ipv6Pattern) checkConstraint() {

}

func (obj *macPattern) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Mac != nil {

		obj.Mac().validateObj(setDefault)
	}

}

func (obj *macPattern) checkUnique() {

}

func (obj *macPattern) checkConstraint() {

}

func (obj *integerPattern) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Integer != nil {

		obj.Integer().validateObj(setDefault)
	}

}

func (obj *integerPattern) checkUnique() {

}

func (obj *integerPattern) checkConstraint() {

}

func (obj *checksumPattern) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(setDefault)
	}

}

func (obj *checksumPattern) checkUnique() {

}

func (obj *checksumPattern) checkConstraint() {

}

func (obj *layer1Ieee802X) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *layer1Ieee802X) checkUnique() {

}

func (obj *layer1Ieee802X) checkConstraint() {

}

func (obj *mObject) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// StringParam is required
	if obj.obj.StringParam == "" {
		obj.errors = append(obj.errors, "StringParam is required field on interface MObject")
	}

	if obj.obj.Integer != 0 {
		if obj.obj.Integer < 10 || obj.obj.Integer > 90 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("10 <= MObject.Integer <= 90 but Got %d", obj.obj.Integer))
		}

	}

	// Mac is required
	if obj.obj.Mac == "" {
		obj.errors = append(obj.errors, "Mac is required field on interface MObject")
	}
	if obj.obj.Mac != "" {

		err := obj.validateMac(obj.Mac())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on MObject.Mac"))
		}

	}

	// Ipv4 is required
	if obj.obj.Ipv4 == "" {
		obj.errors = append(obj.errors, "Ipv4 is required field on interface MObject")
	}
	if obj.obj.Ipv4 != "" {

		err := obj.validateIpv4(obj.Ipv4())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on MObject.Ipv4"))
		}

	}

	// Ipv6 is required
	if obj.obj.Ipv6 == "" {
		obj.errors = append(obj.errors, "Ipv6 is required field on interface MObject")
	}
	if obj.obj.Ipv6 != "" {

		err := obj.validateIpv6(obj.Ipv6())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on MObject.Ipv6"))
		}

	}

	// Hex is required
	if obj.obj.Hex == "" {
		obj.errors = append(obj.errors, "Hex is required field on interface MObject")
	}
	if obj.obj.Hex != "" {

		err := obj.validateHex(obj.Hex())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on MObject.Hex"))
		}

	}

}

func (obj *mObject) checkUnique() {

}

func (obj *mObject) checkConstraint() {

}

func (obj *patternPrefixConfigHeaderChecksum) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {
		if *obj.obj.Custom < 0 || *obj.obj.Custom > 65535 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternPrefixConfigHeaderChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternPrefixConfigHeaderChecksum) checkUnique() {

}

func (obj *patternPrefixConfigHeaderChecksum) checkConstraint() {

}

func (obj *patternPrefixConfigAutoFieldTest) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternPrefixConfigAutoFieldTest.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternPrefixConfigAutoFieldTest.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {
		if *obj.obj.Auto < 0 || *obj.obj.Auto > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternPrefixConfigAutoFieldTest.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternPrefixConfigAutoFieldTest) checkUnique() {

}

func (obj *patternPrefixConfigAutoFieldTest) checkConstraint() {

}

func (obj *wObject) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	// WName is required
	if obj.obj.WName == "" {
		obj.errors = append(obj.errors, "WName is required field on interface WObject")
	}
}

func (obj *wObject) checkUnique() {
	if !obj.isUnique("wObject", obj.WName(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("WName with %s already exists", obj.WName()))
	}
}

func (obj *wObject) checkConstraint() {

}

func (obj *zObject) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface ZObject")
	}
}

func (obj *zObject) checkUnique() {
	if !obj.isUnique("zObject", obj.Name(), obj) && obj.resolve {
		obj.errors = append(obj.errors, fmt.Sprintf("Name with %s already exists", obj.Name()))
	}
}

func (obj *zObject) checkConstraint() {

}

func (obj *yObject) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"prefixConfig"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

}

func (obj *yObject) checkUnique() {

}

func (obj *yObject) checkConstraint() {

	ynameCons := []string{
		"zObject.Name", "wObject.WName",
	}

	if !obj.validateConstraint(ynameCons, obj.YName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid zObject.Name|wObject.WName type", obj.YName()),
		)
	}

}

func (obj *subObject) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}
	obj.deps = append(obj.deps, obj)

	var found bool
	for _, key := range []string{"prefixConfig"} {
		found = false
		for _, k := range obj.rootKeys {
			if k == key {
				found = true
			}
		}
		if found {
			continue
		}
		obj.rootKeys = append(obj.rootKeys, key)
	}

}

func (obj *subObject) checkUnique() {

}

func (obj *subObject) checkConstraint() {

	snameCons := []string{
		"wObject.WName",
	}

	if !obj.validateConstraint(snameCons, obj.SName()) && obj.resolve {
		obj.errors = append(
			obj.errors,
			fmt.Sprintf("%s is not a valid wObject.WName type", obj.SName()),
		)
	}

}

func (obj *errorDetails) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *errorDetails) checkUnique() {

}

func (obj *errorDetails) checkConstraint() {

}

func (obj *_error) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *_error) checkUnique() {

}

func (obj *_error) checkConstraint() {

}

func (obj *metrics) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Ports) != 0 {

		if setDefault {
			obj.Ports().clearHolderSlice()
			for _, item := range obj.obj.Ports {
				newObj := newPortMetric(obj.validator)
				newObj.self().obj = item
				obj.Ports().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Ports().Items() {
			item.validateObj(setDefault)
		}

	}

	if len(obj.obj.Flows) != 0 {

		if setDefault {
			obj.Flows().clearHolderSlice()
			for _, item := range obj.obj.Flows {
				newObj := newFlowMetric(obj.validator)
				newObj.self().obj = item
				obj.Flows().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Flows().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *metrics) checkUnique() {

}

func (obj *metrics) checkConstraint() {

}

func (obj *warningDetails) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *warningDetails) checkUnique() {

}

func (obj *warningDetails) checkConstraint() {

}

func (obj *commonResponseError) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *commonResponseError) checkUnique() {

}

func (obj *commonResponseError) checkConstraint() {

}

func (obj *commonResponseSuccess) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *commonResponseSuccess) checkUnique() {

}

func (obj *commonResponseSuccess) checkConstraint() {

}

func (obj *serviceBItemList) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if len(obj.obj.Items) != 0 {

		if setDefault {
			obj.Items().clearHolderSlice()
			for _, item := range obj.obj.Items {
				newObj := newServiceBItem(obj.validator)
				newObj.self().obj = item
				obj.Items().appendHolderSlice(newObj)
			}
		}
		for _, item := range obj.Items().Items() {
			item.validateObj(setDefault)
		}

	}

}

func (obj *serviceBItemList) checkUnique() {

}

func (obj *serviceBItemList) checkConstraint() {

}

func (obj *serviceBItem) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *serviceBItem) checkUnique() {

}

func (obj *serviceBItem) checkConstraint() {

}

func (obj *levelTwo) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.L2P1 != nil {

		obj.L2P1().validateObj(setDefault)
	}

}

func (obj *levelTwo) checkUnique() {

}

func (obj *levelTwo) checkConstraint() {

}

func (obj *levelFour) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.L4P1 != nil {

		obj.L4P1().validateObj(setDefault)
	}

}

func (obj *levelFour) checkUnique() {

}

func (obj *levelFour) checkConstraint() {

}

func (obj *patternIpv4PatternIpv4) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternIpv4.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternIpv4.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternIpv4PatternIpv4) checkUnique() {

}

func (obj *patternIpv4PatternIpv4) checkConstraint() {

}

func (obj *patternIpv6PatternIpv6) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv6PatternIpv6.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv6PatternIpv6.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternIpv6PatternIpv6) checkUnique() {

}

func (obj *patternIpv6PatternIpv6) checkConstraint() {

}

func (obj *patternMacPatternMac) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternMac.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternMac.Values"))
		}

	}

	if obj.obj.Auto != nil {

		err := obj.validateMac(obj.Auto())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternMac.Auto"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternMacPatternMac) checkUnique() {

}

func (obj *patternMacPatternMac) checkConstraint() {

}

func (obj *patternIntegerPatternInteger) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Value != nil {
		if *obj.obj.Value < 0 || *obj.obj.Value > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternIntegerPatternInteger.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < 0 || item > 255 {
				obj.errors = append(
					obj.errors,
					fmt.Sprintf("0 <= PatternIntegerPatternInteger.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(setDefault)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(setDefault)
	}

}

func (obj *patternIntegerPatternInteger) checkUnique() {

}

func (obj *patternIntegerPatternInteger) checkConstraint() {

}

func (obj *patternChecksumPatternChecksum) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {
		if *obj.obj.Custom < 0 || *obj.obj.Custom > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternChecksumPatternChecksum.Custom <= 255 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternChecksumPatternChecksum) checkUnique() {

}

func (obj *patternChecksumPatternChecksum) checkConstraint() {

}

func (obj *patternPrefixConfigAutoFieldTestCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternPrefixConfigAutoFieldTestCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternPrefixConfigAutoFieldTestCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternPrefixConfigAutoFieldTestCounter) checkUnique() {

}

func (obj *patternPrefixConfigAutoFieldTestCounter) checkConstraint() {

}

func (obj *portMetric) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface PortMetric")
	}
}

func (obj *portMetric) checkUnique() {

}

func (obj *portMetric) checkConstraint() {

}

func (obj *flowMetric) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == "" {
		obj.errors = append(obj.errors, "Name is required field on interface FlowMetric")
	}
}

func (obj *flowMetric) checkUnique() {

}

func (obj *flowMetric) checkConstraint() {

}

func (obj *levelThree) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

}

func (obj *levelThree) checkUnique() {

}

func (obj *levelThree) checkConstraint() {

}

func (obj *patternIpv4PatternIpv4Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternIpv4Counter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternIpv4Counter.Step"))
		}

	}

}

func (obj *patternIpv4PatternIpv4Counter) checkUnique() {

}

func (obj *patternIpv4PatternIpv4Counter) checkConstraint() {

}

func (obj *patternIpv6PatternIpv6Counter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv6PatternIpv6Counter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv6PatternIpv6Counter.Step"))
		}

	}

}

func (obj *patternIpv6PatternIpv6Counter) checkUnique() {

}

func (obj *patternIpv6PatternIpv6Counter) checkConstraint() {

}

func (obj *patternMacPatternMacCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternMacCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			obj.errors = append(obj.errors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternMacCounter.Step"))
		}

	}

}

func (obj *patternMacPatternMacCounter) checkUnique() {

}

func (obj *patternMacPatternMacCounter) checkConstraint() {

}

func (obj *patternIntegerPatternIntegerCounter) validateObj(setDefault bool) {
	if setDefault {
		obj.setDefault()
	}

	if obj.obj.Start != nil {
		if *obj.obj.Start < 0 || *obj.obj.Start > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternIntegerPatternIntegerCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {
		if *obj.obj.Step < 0 || *obj.obj.Step > 255 {
			obj.errors = append(
				obj.errors,
				fmt.Sprintf("0 <= PatternIntegerPatternIntegerCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *patternIntegerPatternIntegerCounter) checkUnique() {

}

func (obj *patternIntegerPatternIntegerCounter) checkConstraint() {

}
