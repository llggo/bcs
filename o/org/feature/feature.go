package feature

type FeatureAction string

type Action map[FeatureAction]bool

type FeatureName string

// type featureName map[Name]bool

type Feature map[FeatureName]Action

const Create = FeatureAction("create")
const Update = FeatureAction("update")
const List = FeatureAction("list")
const Count = FeatureAction("count")
const Get = FeatureAction("get")
const Delete = FeatureAction("delete")

const Qrcode = FeatureName("qrcode")
const Customize = FeatureName("customize")
