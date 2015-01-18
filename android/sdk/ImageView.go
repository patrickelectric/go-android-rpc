// It is autogenerated bindings for Android sdk class.
//
// See documentation about methods on: https://developer.android.com//reference/android/widget/ImageView.html
package sdk

import "github.com/seletskiy/go-android-rpc/android"

type ImageView struct {
	View
}

func NewImageView(id string) interface{} {
	obj := ImageView{NewView(id).(View)}

	return obj
}


func init() {
	android.ViewTypeConstructors["ImageView"] = NewImageView
}

func (obj ImageView) ClearColorFilter() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"clearColorFilter",
	)
}

func (obj ImageView) DrawableHotspotChanged(x_ float64, y_ float64) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"drawableHotspotChanged",
		android.Float(x_), android.Float(y_),
	)
}

func (obj ImageView) GetAdjustViewBounds() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"getAdjustViewBounds",
	)
}

func (obj ImageView) GetBaseline() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"getBaseline",
	)
}

func (obj ImageView) GetBaselineAlignBottom() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"getBaselineAlignBottom",
	)
}

func (obj ImageView) GetCropToPadding() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"getCropToPadding",
	)
}

func (obj ImageView) GetImageAlpha() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"getImageAlpha",
	)
}

func (obj ImageView) GetMaxHeight() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"getMaxHeight",
	)
}

func (obj ImageView) GetMaxWidth() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"getMaxWidth",
	)
}

func (obj ImageView) HasOverlappingRendering() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"hasOverlappingRendering",
	)
}

func (obj ImageView) IsOpaque() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"isOpaque",
	)
}

func (obj ImageView) JumpDrawablesToCurrentState() {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"jumpDrawablesToCurrentState",
	)
}

func (obj ImageView) OnRtlPropertiesChanged(layoutDirection_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"onRtlPropertiesChanged",
		layoutDirection_,
	)
}

func (obj ImageView) SetAdjustViewBounds(adjustViewBounds_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"setAdjustViewBounds",
		adjustViewBounds_,
	)
}

func (obj ImageView) SetAlpha(alpha_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"setAlpha",
		alpha_,
	)
}

func (obj ImageView) SetBaseline(baseline_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"setBaseline",
		baseline_,
	)
}

func (obj ImageView) SetBaselineAlignBottom(aligned_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"setBaselineAlignBottom",
		aligned_,
	)
}

func (obj ImageView) SetColorFilter(color_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"setColorFilter",
		color_,
	)
}

func (obj ImageView) SetCropToPadding(cropToPadding_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"setCropToPadding",
		cropToPadding_,
	)
}

func (obj ImageView) SetImageAlpha(alpha_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"setImageAlpha",
		alpha_,
	)
}

func (obj ImageView) SetImageLevel(level_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"setImageLevel",
		level_,
	)
}

func (obj ImageView) SetImageResource(resId_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"setImageResource",
		resId_,
	)
}

func (obj ImageView) SetImageState(merge_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"setImageState",
		merge_,
	)
}

func (obj ImageView) SetMaxHeight(maxHeight_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"setMaxHeight",
		maxHeight_,
	)
}

func (obj ImageView) SetMaxWidth(maxWidth_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"setMaxWidth",
		maxWidth_,
	)
}

func (obj ImageView) SetSelected(selected_ bool) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"setSelected",
		selected_,
	)
}

func (obj ImageView) SetVisibility(visibility_ int) {
	android.CallViewMethod(
		obj.GetInternalId_(),
		"ImageView",
		"setVisibility",
		visibility_,
	)
}
