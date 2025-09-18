package themes

import (
	"github.com/AllenDang/cimgui-go/imgui"
	g "github.com/AllenDang/giu"
)

func Light() *g.StyleSetter {
	return g.Style().
		SetStyleFloat(g.StyleVarWindowRounding, 2).
		SetStyleFloat(g.StyleVarFrameRounding, 4).
		SetStyleFloat(g.StyleVarGrabRounding, 4).
		SetStyleFloat(g.StyleVarFrameBorderSize, 1).
		SetColorVec4(g.StyleColorText, imgui.Vec4{X: 0.10, Y: 0.10, Z: 0.10, W: 1.00}).
		SetColorVec4(g.StyleColorTextDisabled, imgui.Vec4{X: 0.60, Y: 0.60, Z: 0.60, W: 1.00}).
		SetColorVec4(g.StyleColorWindowBg, imgui.Vec4{X: 0.94, Y: 0.94, Z: 0.94, W: 1.00}).
		SetColorVec4(g.StyleColorChildBg, imgui.Vec4{X: 0.97, Y: 0.97, Z: 0.97, W: 1.00}).
		SetColorVec4(g.StyleColorPopupBg, imgui.Vec4{X: 1.00, Y: 1.00, Z: 1.00, W: 0.98}).
		SetColorVec4(g.StyleColorBorder, imgui.Vec4{X: 0.70, Y: 0.70, Z: 0.70, W: 1.00}).
		SetColorVec4(g.StyleColorBorderShadow, imgui.Vec4{X: 0.00, Y: 0.00, Z: 0.00, W: 0.00}).
		SetColorVec4(g.StyleColorFrameBg, imgui.Vec4{X: 1.00, Y: 1.00, Z: 1.00, W: 1.00}).
		SetColorVec4(g.StyleColorFrameBgHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.40}).
		SetColorVec4(g.StyleColorFrameBgActive, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.67}).
		SetColorVec4(g.StyleColorTitleBg, imgui.Vec4{X: 0.96, Y: 0.96, Z: 0.96, W: 1.00}).
		SetColorVec4(g.StyleColorTitleBgActive, imgui.Vec4{X: 0.82, Y: 0.82, Z: 0.82, W: 1.00}).
		SetColorVec4(g.StyleColorTitleBgCollapsed, imgui.Vec4{X: 1.00, Y: 1.00, Z: 1.00, W: 0.51}).
		SetColorVec4(g.StyleColorMenuBarBg, imgui.Vec4{X: 0.86, Y: 0.86, Z: 0.86, W: 1.00}).
		SetColorVec4(g.StyleColorScrollbarBg, imgui.Vec4{X: 0.98, Y: 0.98, Z: 0.98, W: 0.53}).
		SetColorVec4(g.StyleColorScrollbarGrab, imgui.Vec4{X: 0.69, Y: 0.69, Z: 0.69, W: 1.00}).
		SetColorVec4(g.StyleColorScrollbarGrabHovered, imgui.Vec4{X: 0.59, Y: 0.59, Z: 0.59, W: 1.00}).
		SetColorVec4(g.StyleColorScrollbarGrabActive, imgui.Vec4{X: 0.49, Y: 0.49, Z: 0.49, W: 1.00}).
		SetColorVec4(g.StyleColorCheckMark, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00}).
		SetColorVec4(g.StyleColorSliderGrab, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00}).
		SetColorVec4(g.StyleColorSliderGrabActive, imgui.Vec4{X: 0.06, Y: 0.53, Z: 0.98, W: 1.00}).
		SetColorVec4(g.StyleColorButton, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.40}).
		SetColorVec4(g.StyleColorButtonHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00}).
		SetColorVec4(g.StyleColorButtonActive, imgui.Vec4{X: 0.06, Y: 0.53, Z: 0.98, W: 1.00}).
		SetColorVec4(g.StyleColorHeader, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.31}).
		SetColorVec4(g.StyleColorHeaderHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80}).
		SetColorVec4(g.StyleColorHeaderActive, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00}).
		SetColorVec4(g.StyleColorSeparator, imgui.Vec4{X: 0.39, Y: 0.39, Z: 0.39, W: 0.62}).
		SetColorVec4(g.StyleColorSeparatorHovered, imgui.Vec4{X: 0.14, Y: 0.44, Z: 0.80, W: 0.78}).
		SetColorVec4(g.StyleColorSeparatorActive, imgui.Vec4{X: 0.14, Y: 0.44, Z: 0.80, W: 1.00}).
		SetColorVec4(g.StyleColorResizeGrip, imgui.Vec4{X: 0.35, Y: 0.35, Z: 0.35, W: 0.17}).
		SetColorVec4(g.StyleColorResizeGripHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.67}).
		SetColorVec4(g.StyleColorResizeGripActive, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.95}).
		SetColorVec4(g.StyleColorTab, imgui.Vec4{X: 0.76, Y: 0.80, Z: 0.84, W: 0.93}).
		SetColorVec4(g.StyleColorTabHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80}).
		SetColorVec4(g.StyleColorTabActive, imgui.Vec4{X: 0.60, Y: 0.73, Z: 0.88, W: 1.00}).
		SetColorVec4(g.StyleColorTabUnfocused, imgui.Vec4{X: 0.92, Y: 0.93, Z: 0.94, W: 1.00}).
		SetColorVec4(g.StyleColorTabUnfocusedActive, imgui.Vec4{X: 0.74, Y: 0.82, Z: 0.91, W: 1.00}).
		SetColorVec4(g.StyleColorPlotLines, imgui.Vec4{X: 0.39, Y: 0.39, Z: 0.39, W: 1.00}).
		SetColorVec4(g.StyleColorPlotLinesHovered, imgui.Vec4{X: 1.00, Y: 0.43, Z: 0.35, W: 1.00}).
		SetColorVec4(g.StyleColorPlotHistogram, imgui.Vec4{X: 0.90, Y: 0.70, Z: 0.00, W: 1.00}).
		SetColorVec4(g.StyleColorPlotHistogramHovered, imgui.Vec4{X: 1.00, Y: 0.45, Z: 0.00, W: 1.00}).
		SetColorVec4(g.StyleColorTextSelectedBg, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.35}).
		SetColorVec4(g.StyleColorDragDropTarget, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.95}).
		SetColorVec4(g.StyleColorNavWindowingHighlight, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80}).
		SetColorVec4(g.StyleColorNavWindowingHighlight, imgui.Vec4{X: 0.70, Y: 0.70, Z: 0.70, W: 0.70}).
		SetColorVec4(g.StyleColorTableHeaderBg, imgui.Vec4{X: 0.78, Y: 0.87, Z: 0.98, W: 1.00}).
		SetColorVec4(g.StyleColorTableBorderStrong, imgui.Vec4{X: 0.57, Y: 0.57, Z: 0.64, W: 1.00}).
		SetColorVec4(g.StyleColorTableBorderLight, imgui.Vec4{X: 0.68, Y: 0.68, Z: 0.74, W: 1.00})
}
