package constants

const (
	ImagesDirname             = "images"
	TemplateDirname           = "TEMPLATE"
	MainTfFilename            = "main.tf"
	GeneratedTfFilename       = "generated.tf"
	GeneratedTfHeader         = "# DO NOT EDIT - this file is autogenerated by tfgen"
	TfFileExtension           = ".tf"
	TfTemplateFileExtention   = ".tf.tpl"
	TfTypeOutput              = "output"
	TfTypeModule              = "module"
	TfTypeResource            = "resource"
	TfLibTagger               = "../../tflib/tagger"
	TfEachKey                 = "each.key"
	TfEachValue               = "each.value"
	OutputNameSummary         = "summary"
	OutputNameImageRef        = "image_ref"
	OutputNameTagList         = "tag_list"
	OutputNameLatestTagMap    = "latest_tag_map"
	AttributeSource           = "source"
	AttributePublisher        = "publisher"
	AttributeForEach          = "for_each"
	AttributeValue            = "value"
	AttributeProviders        = "providers"
	AttributeTargetRepository = "target_repository"
	AttributeLicenseKey       = "license_key"
	AttributeName             = "name"
	AttributeDigestRef        = "digest_ref"
	AttributeDependsOn        = "depends_on"
	AttributeTag              = "tag"
	AttributeTags             = "tags"
	AttributeMainPackage      = "main_package"
	ResourceOciTag            = "oci_tag"
	NewRelicImagePrefix       = "newrelic-"
	ScannerTestImageName      = "scanner-test"
	CustomImagePrefix         = "request-"
	DetectTfLib               = "/tflib/"
)
